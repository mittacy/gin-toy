package thirdHttp

import (
	"github.com/mittacy/gin-toy/core/log"
	"time"
)

type SimpleClient struct {
	logName string
	log     *log.Logger
	host    string
	timeout time.Duration
}

func NewSimpleClient(host string, options ...SimpleClientOption) *SimpleClient {
	c := &SimpleClient{
		logName: "thirdHttp",
		host:    host,
		timeout: time.Second * 5,
	}

	for _, option := range options {
		option(c)
	}

	c.log = log.New(c.logName, log.WithLevel(log.InfoLevel))

	return c
}

type SimpleClientOption func(client *SimpleClient)

// SimpleWithTimeout 自定义超时时间，默认为 5s
func SimpleWithTimeout(timeout time.Duration) SimpleClientOption {
	return func(c *SimpleClient) {
		c.timeout = timeout
	}
}

// SimpleWithLogName 自定义日志名，默认为 thirdHttp
func SimpleWithLogName(name string) SimpleClientOption {
	return func(c *SimpleClient) {
		c.logName = name
	}
}
