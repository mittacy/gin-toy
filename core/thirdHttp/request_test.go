package thirdHttp

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/mittacy/gin-toy/core/log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestFullUrl(t *testing.T) {
	cases := []struct {
		Name   string
		Host   string
		Uri    string
		Expect string
	}{
		{"uri有/", "https://www.baidu.com", "/index.html", "https://www.baidu.com/index.html"},
		{"host有/", "https://www.baidu.com/", "index.html", "https://www.baidu.com/index.html"},
		{"都没有/", "https://www.baidu.com", "index.html", "https://www.baidu.com/index.html"},
		{"都有/", "https://www.baidu.com/", "/index.html", "https://www.baidu.com/index.html"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := fullUrl(c.Host, c.Uri); ans != c.Expect {
				t.Fatalf("input: %v, %v digits, expected %v, but %v got",
					c.Host, c.Uri, c.Expect, ans)
			}
		})
	}
}

type Student struct {
	Name string `mapstructure:"name"`
	Age  int    `mapstructure:"age"`
}

var startOnce sync.Once

func startHttpAndLog() {
	log.Init(log.WithPath("./"),
		log.WithTimeFormat("2006-01-02 15:04:05"),
		log.WithLevel(log.DebugLevel),
		log.WithEncoderJSON(true),
		log.WithLogInConsole(true),
		log.WithDistinguish(true))

	r := gin.New()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":      "mittacy",
			"age":       12,
			"code":      "success",
			"short_msg": "meddle with sth",
			"id":        c.Query("id"),
		})
	})
	r.GET("/students", func(c *gin.Context) {
		students := []Student{
			{"mittacy", 11},
			{"lise", 12},
			{"mick", 14},
			{"neo", 10},
			{"jack", 15},
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": students})
	})
	r.GET("/students_total", func(c *gin.Context) {
		students := []Student{
			{"mittacy", 11},
			{"lise", 12},
			{"mick", 14},
			{"neo", 10},
			{"jack", 15},
		}

		res := map[string]interface{}{
			"list":  students,
			"total": 5,
			"key":   c.Query("key"),
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": res})
	})
	r.POST("/change", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": "success"})
	})

	r.Run(":10110")
}

func TestGetObject(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_getObject")
	client := NewClient("http://127.0.0.1:10110")

	res := struct {
		List  []Student `mapstructure:"list"`
		Total int       `mapstructure:"total"`
	}{}
	if _, err := client.Get(c, "/students_total", &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Log(res)
}

func TestGetArr(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_testArr")
	client := NewClient("http://127.0.0.1:10110")

	var res []Student
	if _, err := client.Get(c, "/students", &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Log(res)
}

func TestGetParams(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_getParams")
	client := NewClient("http://127.0.0.1:10110")

	res := struct {
		List  []Student `mapstructure:"list"`
		Total int       `mapstructure:"total"`
		Key   string    `mapstructure:"key"`
	}{}
	params := map[string]string{
		"key": "hhhh",
	}
	if _, err := client.GetParams(c, "/students_total", params, &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Log(res)
}

func TestPost(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_post")
	client := NewClient("http://127.0.0.1:10110")

	var res string
	if _, err := client.Post(c, "/change", nil, &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Log(res)
}

func TestConfig(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_config")
	client := NewClient("http://127.0.0.1:10110", WithLogName("cutsomLog"), WithTimeout(time.Second))

	res := struct {
		List  []Student `mapstructure:"list"`
		Total int       `mapstructure:"total"`
	}{}
	if _, err := client.Get(c, "/students_total", &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Log(res)
}

func TestCustomReply(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_custom_reply")
	var res customReply
	client := NewClient("http://127.0.0.1:10110", WithReply(&res))

	if _, err := client.Get(c, "/index", &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Log(res)
}

func TestDefaultReply(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	SetDefaultReply(&customReply{})

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_custom_reply")
	var res customReply
	client := NewClient("http://127.0.0.1:10110")

	if _, err := client.Get(c, "/index", &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Log(res)
}

type customReply struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Code     string `json:"code"`
	ShortMsg string `json:"short_msg"`
}

func (r *customReply) GetCode() int {
	return 0
}
func (r *customReply) GetMsg() string {
	return r.ShortMsg
}
func (r *customReply) GetUnknownCode() int {
	return 500
}
func (r *customReply) IsSuccess() bool {
	return true
}
func (r *customReply) UnmarshalData(result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc([]string{"2006-01-02 15:04:05", time.RFC3339, time.RFC3339Nano}...)),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(r); err != nil {
		return err
	}
	return err
}
