package grab

import (
	"fmt"
	"time"

	"github.com/R0L/golang-test/grab/util/cipher/aes"
	"github.com/R0L/golang-test/grab/util/http"
)

type Graber interface {
	GetXToken() (string, error)
	GetSceneInfo(sceneId string)
	panicBuying(sceneId string)
}

type grab struct {
	key    string
	url    string
	userId string
	ip     string
}

func NewGrab(opts ...Option) Graber {
	g := &grab{}
	for _, o := range opts {
		o.apply(g)
	}
	return g
}

// token
func (g *grab) GetXToken() (string, error) {
	xTokenStr := fmt.Sprintf("%d_2_%s", time.Now().UnixNano()/1e6, g.userId)
	return aes.NewAES(g.key).Encrypt(xTokenStr)
}

// 可抢
func (g *grab) GetSceneInfo(sceneId string) {

	var err error

	url := fmt.Sprintf("%s/est/app/scene/getSceneInfo", g.url)

	result := make(map[string]interface{})

	xToken, err := g.GetXToken()
	if nil != err {
		fmt.Println(err)
		return
	}

	err = http.NewHttp().Post(url, map[string]string{
		"sceneId": sceneId,
	}, map[string]string{
		"x-token": xToken,
	}, &result)
	if nil != err {
		fmt.Println(err)
		return
	}
	if result["code"] == float64(200) {
		if result["data"].(map[string]interface{})["countDown"].(float64) == float64(0) {
			g.panicBuying(sceneId)
		} else {
			fmt.Println(g.userId, result)
		}
	}
}

// 抢
func (g *grab) panicBuying(sceneId string) {

	url := fmt.Sprintf("%s/est/app/scene/panicBuying", g.url)

	xToken, err := g.GetXToken()
	if nil != err {
		fmt.Println(err)
		return
	}

	result := make(map[string]interface{})

	err = http.NewHttp().Post(url, map[string]string{
		"sceneId": sceneId,
	}, map[string]string{
		"x-token":         xToken,
		"X-FORWARDED-FOR": g.ip,
		"CLIENT-IP":       g.ip,
		"CURLOPT_REFERER": g.ip,
	}, &result)
	if nil != err {
		// fmt.Println(err)
		return
	}

	if result["code"] == float64(200) {
		fmt.Println(g.userId, result)
	} else {
		fmt.Print(".")
	}
	fmt.Println(g.userId, result)

}

type Option interface {
	apply(*grab)
}

type keyOption string

func (c keyOption) apply(g *grab) {
	g.key = string(c)
}

func WithKey(key string) Option {
	return keyOption(key)
}

type urlOption string

func (c urlOption) apply(g *grab) {
	g.url = string(c)
}

func WithUrl(url string) Option {
	return urlOption(url)
}

type userIdOption string

func (c userIdOption) apply(g *grab) {
	g.userId = string(c)
}

func WithUserId(userId string) Option {
	return userIdOption(userId)
}

type ipOption string

func (c ipOption) apply(g *grab) {
	g.ip = string(c)
}

func WithIp(ip string) Option {
	return ipOption(ip)
}
