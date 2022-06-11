package thirdHttp

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strings"
	"time"
)

func fullUrl(host, uri string) string {
	uri = strings.TrimLeft(uri, "/")
	host = strings.TrimRight(host, "/")
	return host + "/" + uri
}

func ToTimeHookFunc(timeFormat ...string) mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return ParseTime(data, timeFormat...)
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
	}
}

var ErrParseTime = errors.New("时间格式错误")

// ParseTime 解析时间字符串为time.Time
// @param data 时间
// @param format 可能的时间格式
// @return time.Time
// @return error 如果格式错误，将返回 ErrTimeFormat
func ParseTime(data interface{}, formats ...string) (time.Time, error) {
	for _, v := range formats {
		if result, err := time.Parse(v, data.(string)); err != nil {
			return result, nil
		}
	}

	return time.Time{}, ErrParseTime

}
