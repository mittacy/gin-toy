package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy/core/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"math/rand"
	"time"
)

// Config is config setting for Ginzap
type Config struct {
	TimeFormat string
	UTC        bool
	SkipPaths  []string
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestTraceAndLog returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func RequestTraceAndLog(logger *log.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return ginZapWithConfig(logger, &Config{TimeFormat: timeFormat, UTC: utc})
}

// ginZapWithConfig returns a gin.HandlerFunc using configs
func ginZapWithConfig(logger *log.Logger, conf *Config) gin.HandlerFunc {
	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		// 写入请求时间
		start := time.Now()

		// 写入请求id
		requestId := newRequestId()
		c.Set(log.RequestIdKey(), requestId)

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 响应体
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 提取请求体
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			logger.ErrorwWithTrace(c, "Invalid request body", "err", err)
			c.Abort()
			return
		}

		// 新建缓冲区并替换原有Request.body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		// 后置
		if _, ok := skipPaths[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)
			if conf.UTC {
				end = end.UTC()
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					logger.ErrorWithTrace(c, e)
				}
			} else {
				fields := []zapcore.Field{
					zap.Int("status", c.Writer.Status()),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("body", string(bodyBytes)),
					zap.String("resp", blw.body.String()),
					zap.Duration("latency", latency),
				}
				if conf.TimeFormat != "" {
					fields = append(fields, zap.String("time", end.Format(conf.TimeFormat)))
				}
				logger.InfoWithTrace(c, path, fields...)
			}
		}
	}
}

// newRequestId 生成新的请求id
func newRequestId() string {
	now := time.Now()
	s := fmt.Sprintf("%s%08x%05x", "r", now.Unix(), now.UnixNano()%0x100000)
	return s + "_" + randomStr(18)
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// randomStr 生成随机字符串
// @param n 生成的字符串长度
// @return string 返回生成的随机字符串
func randomStr(n int, randChars ...[]rune) string {
	if n <= 0 {
		return ""
	}

	var letters []rune

	if len(randChars) == 0 {
		letters = defaultLetters
	} else {
		letters = randChars[0]
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
