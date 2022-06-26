package thirdHttp

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

// Get
// @param c
// @param uri 路由
// @param result data数据结果赋值
// @return error
func (ctl *Client) Get(c context.Context, uri string, result interface{}) (int, error) {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respData := ctl.reply()
	respBody, err := client.R().SetResult(&respData).ForceContentType("application/json").Get(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "respBody", respBody, "err", err)
		return respData.GetUnknownCode(), err
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "respBody", respBody)

	// http请求是否成功
	if !respBody.IsSuccess() {
		return respData.GetUnknownCode(), errors.New(respBody.String())
	}

	// 业务响应是否成功
	if !respData.IsSuccess() {
		return respData.GetCode(), errors.New(respData.GetMsg())
	}

	// 解析data数据
	if err = respData.UnmarshalData(result); err != nil {
		return respData.GetUnknownCode(), err
	}

	return respData.GetCode(), respData.UnmarshalData(result)
}

// GetParams
// @param c
// @param uri 路由
// @param params 请求参数
// @param result data数据结果赋值
// @return error
func (ctl *Client) GetParams(c context.Context, uri string, params map[string]string, result interface{}) (int, error) {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respData := ctl.reply()
	respBody, err := client.R().SetQueryParams(params).SetResult(&respData).ForceContentType("application/json").Get(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "params", params, "respBody", respBody, "err", err)
		return respData.GetUnknownCode(), err
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "params", params, "respBody", respBody)

	// http请求是否成功
	if !respBody.IsSuccess() {
		return respData.GetUnknownCode(), errors.New(respBody.String())
	}

	// 业务响应是否成功
	if !respData.IsSuccess() {
		return respData.GetCode(), errors.New(respData.GetMsg())
	}

	// 解析data数据
	if err = respData.UnmarshalData(result); err != nil {
		return respData.GetUnknownCode(), err
	}

	return respData.GetCode(), respData.UnmarshalData(result)
}

// Post
// @param c
// @param uri 路由
// @param body 请求体
// @param result data数据结果赋值
// @return error
func (ctl *Client) Post(c context.Context, uri string, body interface{}, result interface{}) (int, error) {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respData := ctl.reply()
	respBody, err := client.R().SetHeader("Content-Type", "application/json").SetBody(body).SetResult(&respData).Post(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "params", body, "respBody", respBody, "err", err)
		return respData.GetUnknownCode(), err
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "params", body, "respBody", respBody)

	// http请求是否成功
	if !respBody.IsSuccess() {
		return respData.GetUnknownCode(), errors.New(respBody.String())
	}

	if !respData.IsSuccess() {
		return respData.GetCode(), errors.New(respData.GetMsg())
	}

	// 解析data数据
	if err = respData.UnmarshalData(result); err != nil {
		return respData.GetUnknownCode(), err
	}

	return respData.GetCode(), respData.UnmarshalData(result)
}

// SimpleGet 简单请求，不处理结果
// @param c
// @param uri
// @param result
// @return error
func (ctl *Client) SimpleGet(c context.Context, uri string, result interface{}) error {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respBody, err := client.R().SetResult(result).ForceContentType("application/json").Get(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "respBody", respBody, "err", err)
		return err
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "respBody", respBody, "result", result)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	return nil
}

// SimpleGetParams 简单带参数请求，不处理结果
// @param c
// @param uri
// @param params
// @param result
// @return error
func (ctl *Client) SimpleGetParams(c context.Context, uri string, params map[string]string, result interface{}) error {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respBody, err := client.R().SetQueryParams(params).SetResult(result).ForceContentType("application/json").Get(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "params", params, "respBody", respBody, "err", err)
		return err
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "params", params, "respBody", respBody)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	return nil
}

// SimplePost 简单请求，不处理结果
// @param c
// @param uri
// @param body
// @param result
// @return error
func (ctl *Client) SimplePost(c context.Context, uri string, body interface{}, result interface{}) error {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respBody, err := client.R().SetHeader("Content-Type", "application/json").SetBody(body).SetResult(result).Post(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "params", body, "respBody", respBody, "err", err)
		return err
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "params", body, "respBody", result)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	return nil
}
