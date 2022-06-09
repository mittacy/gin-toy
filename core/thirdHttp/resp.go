package thirdHttp

import (
	"github.com/mitchellh/mapstructure"
	"time"
)

type IReply interface {
	GetCode() interface{}
	GetMsg() interface{}
	GetData() interface{}
	IsSuccess() bool
	UnmarshalData(result interface{}) error
}

func NewReply() IReply {
	return &Reply{
		successCode: 0,
		timeFormats: []string{"2006-01-02 15:04:05", time.RFC3339, time.RFC3339Nano},
		Code:        0,
		Msg:         "",
		Data:        nil,
	}
}

type Reply struct {
	successCode int
	timeFormats []string
	Code        int
	Msg         string
	Data        interface{}
}

func (r *Reply) GetCode() interface{} {
	return r.Code
}

func (r *Reply) GetMsg() interface{} {
	return r.Msg
}

func (r *Reply) GetData() interface{} {
	return r.Data
}

func (r *Reply) UnmarshalData(result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc(r.timeFormats...)),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(r.Data); err != nil {
		return err
	}
	return err
}

func (r *Reply) IsSuccess() bool {
	return r.Code == r.successCode
}
