package thirdHttp

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy/core/log"
	"testing"
)

func TestSimpleGet(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_simple_get")
	client := NewSimpleClient("http://127.0.0.1:10110")

	res := struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Code     string `json:"code"`
		ShortMsg string `json:"short_msg"`
	}{}
	if err := client.Get(c, "/index", &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Logf("%+v", res)
}

func TestSimpleGetParams(t *testing.T) {
	go startOnce.Do(startHttpAndLog)

	c := &gin.Context{}
	c.Set(log.RequestIdKey(), "r_simple_get_params")
	client := NewSimpleClient("http://127.0.0.1:10110")

	res := struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Code     string `json:"code"`
		ShortMsg string `json:"short_msg"`
		Id       string `json:"id"`
	}{}
	params := map[string]string{
		"id": "12",
	}
	if err := client.GetParams(c, "/index", params, &res); err != nil {
		t.Errorf("get err: %+v", err)
	}

	t.Logf("%+v", res)
}
