package thirdHttp

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// Get
// @param c
// @param uri 路由
// @param result data数据结果赋值
// @return error
func (ctl *Client) Get(c context.Context, uri string, result interface{}) error {
	url := fullUrl(ctl.Host(), uri)
	client := resty.New().SetTimeout(ctl.Timeout())

	respData := ctl.reply
	respBody, err := client.R().SetResult(&respData).ForceContentType("application/json").Get(url)
	if err != nil {
		ctl.Log().ErrorwWithTrace(c, url, "respBody", respBody, "err", err)
		return errors.WithStack(err)
	}
	defer respBody.RawBody().Close()

	ctl.Log().InfowWithTrace(c, url, "respBody", respBody, "respData", respData)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	if !respData.IsSuccess() {
		return errors.New(fmt.Sprintf("code: %d, msg: %s", respData.GetCode(), respData.GetMsg()))
	}

	return respData.UnmarshalData(result)
}

// GetParams
// @param c
// @param uri 路由
// @param params 请求参数
// @param result data数据结果赋值
// @return error
func (ctl *Client) GetParams(c context.Context, uri string, params map[string]string, result interface{}) error {
	url := fullUrl(ctl.Host(), uri)
	client := resty.New().SetTimeout(ctl.Timeout())

	respData := ctl.reply
	respBody, err := client.R().SetQueryParams(params).SetResult(&respData).ForceContentType("application/json").Get(url)
	if err != nil {
		ctl.Log().ErrorwWithTrace(c, url, "params", params, "respBody", respBody, "err", err)
		return errors.WithStack(err)
	}
	defer respBody.RawBody().Close()

	ctl.Log().InfowWithTrace(c, url, "params", params, "respBody", respBody, "respData", respData)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	if !respData.IsSuccess() {
		return errors.New(fmt.Sprintf("code: %d, msg: %s", respData.GetCode(), respData.GetMsg()))
	}

	return respData.UnmarshalData(result)
}

// Post
// @param c
// @param uri 路由
// @param body 请求体
// @param result data数据结果赋值
// @return error
func (ctl *Client) Post(c context.Context, uri string, body interface{}, result interface{}) error {
	url := fullUrl(ctl.Host(), uri)
	client := resty.New().SetTimeout(ctl.Timeout())

	respData := ctl.reply
	respBody, err := client.R().SetHeader("Content-Type", "application/json").SetBody(body).SetResult(&respData).Post(url)
	if err != nil {
		ctl.Log().ErrorwWithTrace(c, url, "params", body, "respBody", respBody, "err", err)
		return errors.WithStack(err)
	}
	defer respBody.RawBody().Close()

	ctl.Log().InfowWithTrace(c, url, "params", body, "respBody", respBody, "respData", respData)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	if !respData.IsSuccess() {
		return errors.New(fmt.Sprintf("code: %d, msg: %s", respData.GetCode(), respData.GetMsg()))
	}

	return respData.UnmarshalData(result)
}
