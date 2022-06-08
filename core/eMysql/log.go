package eMysql

import (
	"github.com/mittacy/gin-toy/core/log"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
	"time"
)

var l zapgorm2.Logger

func initLog(options ...GormConfigOption) {
	gc := logConf{
		Name:                 "mysql",
		SlowThreshold:        100 * time.Millisecond,
		IgnoreRecordNotFound: false,
	}

	for _, option := range options {
		option(&gc)
	}

	l = zapgorm2.New(log.New(gc.Name).Log())
	l.SlowThreshold = gc.SlowThreshold
	l.LogLevel = logger.Info
	l.IgnoreRecordNotFoundError = gc.IgnoreRecordNotFound
	l.SetAsDefault()
}

type logConf struct {
	Name                 string
	SlowThreshold        time.Duration
	IgnoreRecordNotFound bool
}

type GormConfigOption func(conf *logConf)

// WithName 慢日志名，默认为 mysql
func WithName(name string) GormConfigOption {
	return func(conf *logConf) {
		conf.Name = name
	}
}

// WithSlowThreshold 慢日志时间阈值，默认为 100毫秒
func WithSlowThreshold(duration time.Duration) GormConfigOption {
	return func(conf *logConf) {
		conf.SlowThreshold = duration
	}
}

// WithIgnoreRecordNotFound 是否忽略notFound错误，默认为 false
func WithIgnoreRecordNotFound(isIgnore bool) GormConfigOption {
	return func(conf *logConf) {
		conf.IgnoreRecordNotFound = isIgnore
	}
}
