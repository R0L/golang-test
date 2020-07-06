package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type RequestParam struct {
	Method      Method
	ContentType ContentType
	Header      map[string]string
	Timeout     uint32
	Param       interface{}
}

type Httper interface {
	Request(url string, rParam *RequestParam, result interface{}) (err error)
	Post(url string, param interface{}, header map[string]string, result interface{}) (err error)
	Get(url string, param interface{}, header map[string]string, result interface{}) (err error)
}

type customHttp struct {
}

func NewHttp() Httper {
	return &customHttp{}
}

type ContentType string
type Method string

const (
	// 默认超时时长
	DefaultTimeOut = 3

	ContentTypeMultipart ContentType = "multipart/form-data"
	ContentTypeForm      ContentType = "application/x-www-form-urlencoded"
	ContentTypeJson      ContentType = "application/json"

	POST Method = "POST"
	GET  Method = "GET"
)

/**
 * @param url 请求URL
 * @param rParam = &RequestParam{
 *       Method string  请求方法，支持POST,GET
 *	     ContentType ContentType 请求内容类型，支持
 *       Cookie string 请求cookie
 *       Param interface{} 请求参数，支持map、struct、*map、*struct
 *	}
 * @param Timeout uint32 请求超时时长，默认3s
 * @param result 返回结果
 */
func (c *customHttp) Request(url string, rParam *RequestParam, result interface{}) (err error) {
	if rParam == nil {
		return errors.New("rParam not nil")
	}
	param := rParam.Param

	// 默认时长
	if rParam.Timeout == 0 {
		rParam.Timeout = DefaultTimeOut
	}

	var queryParam = make(map[string]string)
	if param != nil {
		vp := reflect.ValueOf(param)
		tp := reflect.TypeOf(param)
		if tp.Kind() == reflect.Ptr {
			vp = vp.Elem()
			tp = tp.Elem()
		}
		if tp.Kind() == reflect.Map {
			iter := vp.MapRange()
			for iter.Next() {
				k := iter.Key()
				v := iter.Value()
				queryParam[k.String()] = v.String()
			}
		} else if tp.Kind() == reflect.Struct {
			for i := 0; i < vp.NumField(); i++ {
				k := tp.Field(i).Name
				v := vp.Field(i).String()
				queryParam[k] = v
			}
		} else {
			panic("param.param type err")
		}
	}
	var payload io.Reader
	if rParam.ContentType == ContentTypeMultipart {
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		for k, v := range queryParam {
			err = writer.WriteField(k, v)
			if nil != err {
				return err
			}
		}
		err := writer.Close()
		if err != nil {
			return err
		}
		rParam.ContentType = ContentType(writer.FormDataContentType())
	} else if rParam.ContentType == ContentTypeForm {
		var paramSlice []string
		for k, v := range queryParam {
			paramSlice = append(paramSlice, fmt.Sprintf("%s=%s", k, v))
		}
		payload = strings.NewReader(strings.Join(paramSlice, "&"))

	} else {
		byteQuery, _ := json.Marshal(queryParam)
		payload = strings.NewReader(string(byteQuery[:]))
	}
	client := &http.Client{
		Timeout: time.Duration(rParam.Timeout) * time.Second,
	}
	req, err := http.NewRequest(string(rParam.Method), url, payload)
	if err != nil {
		return err
	}

	for headerkey, headerValue := range rParam.Header {
		req.Header.Add(headerkey, headerValue)
	}
	/**
	multipart/form-data == writer.FormDataContentType()
	application/x-www-form-urlencoded
	application/json
	*/
	req.Header.Set("Content-Type", string(rParam.ContentType))

	res, err := client.Do(req)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer func() {
		err := res.Body.Close()
		if nil != err {
			fmt.Println(err)
			return
		}
	}()
	body, err := ioutil.ReadAll(res.Body)
	if nil != err {
		return err
	}

	err = json.Unmarshal(body, result)
	if nil != err {
		return err
	}
	return
}

func (c *customHttp) Post(url string, param interface{}, header map[string]string, result interface{}) (err error) {
	return c.Request(url, &RequestParam{
		Method:      POST,
		ContentType: ContentTypeJson,
		Header:      header,
		Param:       param,
	}, result)
}

func (c *customHttp) Get(url string, param interface{}, header map[string]string, result interface{}) (err error) {
	return c.Request(url, &RequestParam{
		Method:      GET,
		ContentType: ContentTypeForm,
		Header:      header,
		Param:       param,
	}, result)
}
