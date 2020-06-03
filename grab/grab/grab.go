package grab

import (
	"fmt"
	"time"

	"github.com/R0L/golang-test/grab/util/cipher/aes"
	"github.com/R0L/golang-test/grab/util/http"
)

type Graber interface {
	getXToken() (string, error)
	GetSceneInfo(sceneId string)
	panicBuying(sceneId string)
}

type grab struct {
	key    string
	url    string
	userId string
}

func NewGrab(opts ...Option) Graber {
	g := &grab{}
	for _, o := range opts {
		o.apply(g)
	}
	return g
}

// token
func (g *grab) getXToken() (string, error) {
	xTokenStr := fmt.Sprintf("%d_2_%s", time.Now().UnixNano()/1e6, g.userId)
	return aes.NewAES(g.key).Encrypt(xTokenStr)
}

// 可抢
func (g *grab) GetSceneInfo(sceneId string) {

	var err error

	url := fmt.Sprintf("%s/est/app/scene/getSceneInfo", g.url)

	result := make(map[string]interface{})

	xToken, err := g.getXToken()
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
			fmt.Println(result)
		}
	}
}

// 抢
func (g *grab) panicBuying(sceneId string) {

	url := fmt.Sprintf("%s/est/app/scene/panicBuying", g.url)

	xToken, err := g.getXToken()
	if nil != err {
		fmt.Println(err)
		return
	}

	result := make(map[string]interface{})

	err = http.NewHttp().Post(url, map[string]string{
		"sceneId": sceneId,
	}, map[string]string{
		"x-token":         xToken,
		"X-FORWARDED-FOR": "192.168.31.105",
		"CLIENT-IP":       "192.168.31.105",
		"CURLOPT_REFERER": "192.168.31.105",
	}, &result)
	if nil != err {
		// fmt.Println(err)
		return
	}

	fmt.Println(".")

	fmt.Println(result)

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
