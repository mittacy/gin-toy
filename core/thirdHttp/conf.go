package thirdHttp

import (
	"github.com/mittacy/gin-toy/core/log"
	"time"
)

type Client struct {
	logName string
	log     *log.Logger
	host    string
	timeout time.Duration
	reply   func() IReply
}

func NewClient(host string, options ...ClientOption) *Client {
	c := &Client{
		logName: "thirdHttp",
		host:    host,
		timeout: time.Second * 5,
		reply:   defaultReply,
	}

	for _, option := range options {
		option(c)
	}

	c.log = log.New(c.logName, log.WithLevel(log.InfoLevel))

	return c
}

type ClientOption func(client *Client)

// WithTimeout 自定义超时时间，默认为 5s
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithLogName 自定义日志名，默认为 thirdHttp
func WithLogName(name string) ClientOption {
	return func(c *Client) {
		c.logName = name
	}
}

// WithReply 自定义响应结构体, 默认为 {Code int, Msg string, Data interface{}}
func WithReply(reply IReply) ClientOption {
	return func(c *Client) {
		c.reply = func() IReply {
			return reply
		}
	}
}

var defaultReply = NewReply

// SetDefaultReply 设置默认的全局响应体，影响全局
func SetDefaultReply(reply IReply) {
	defaultReply = func() IReply {
		return reply
	}
}
