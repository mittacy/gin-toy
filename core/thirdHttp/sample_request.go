package thirdHttp

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

func (ctl *SimpleClient) Get(c context.Context, uri string, result interface{}) error {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respBody, err := client.R().SetResult(result).ForceContentType("application/json").Get(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "respBody", respBody, "err", err)
		return errors.WithStack(err)
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "respBody", respBody, "result", result)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	return nil
}

func (ctl *SimpleClient) GetParams(c context.Context, uri string, params map[string]string, result interface{}) error {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respBody, err := client.R().SetQueryParams(params).SetResult(result).ForceContentType("application/json").Get(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "params", params, "respBody", respBody, "err", err)
		return errors.WithStack(err)
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "params", params, "respBody", respBody, "result", result)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	return nil
}

func (ctl *SimpleClient) Post(c context.Context, uri string, body interface{}, result interface{}) error {
	url := fullUrl(ctl.host, uri)
	client := resty.New().SetTimeout(ctl.timeout)

	respBody, err := client.R().SetHeader("Content-Type", "application/json").SetBody(body).SetResult(result).Post(url)
	if err != nil {
		ctl.log.ErrorwWithTrace(c, url, "params", body, "respBody", respBody, "err", err)
		return errors.WithStack(err)
	}
	defer respBody.RawBody().Close()

	ctl.log.InfowWithTrace(c, url, "params", body, "respBody", result, "respData", result)

	if !respBody.IsSuccess() {
		return errors.New(respBody.String())
	}

	return nil
}
