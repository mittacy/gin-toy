package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mittacy/gin-toy/core/bizerr"
	"github.com/mittacy/gin-toy/core/log"
	"github.com/pkg/errors"
	"net/http"
)

var NilData = struct{}{}

func Success(c *gin.Context, data interface{}) {
	Custom(c, http.StatusOK, bizerr.Success.Code, bizerr.Success.Msg, data)
}

func SuccessMsg(c *gin.Context, data interface{}, msg string) {
	Custom(c, http.StatusOK, bizerr.Success.Code, msg, data)
}

func SuccessNil(c *gin.Context) {
	Custom(c, http.StatusOK, bizerr.Success.Code, bizerr.Success.Msg, NilData)
}

func FailMsg(c *gin.Context, msg string) {
	Custom(c, http.StatusOK, bizerr.Request.Code, msg, NilData)
}

func FailErr(c *gin.Context, err error) {
	if !bizerr.IsBizErr(err) {
		Unknown(c, err)
		return
	}

	Custom(c, http.StatusOK, bizerr.Code(err), err.Error(), NilData)
}

// FailCheckBizErr 检查错误是否为指定的业务错误，否则记录日志并响应
func FailCheckBizErr(c *gin.Context, title string, sourceErr error, possibleErrs ...error) {
	for _, v := range possibleErrs {
		if bizerr.Is(sourceErr, v) {
			Custom(c, http.StatusOK, bizerr.Code(v), v.Error(), NilData)
			return
		}
	}

	Unknown(c, errors.WithMessage(sourceErr, title))
	return
}

func Unknown(c *gin.Context, err error) {
	log.ErrorfWithTrace(c, "err: %+v", err)
	code := bizerr.Unknown.Code
	msg := bizerr.Unknown.Error()

	if gin.Mode() != gin.ReleaseMode {
		msg = err.Error()
	}

	Custom(c, http.StatusOK, code, msg, NilData)
}

// Unauthorized 未认证
func Unauthorized(c *gin.Context) {
	Custom(c, http.StatusOK, bizerr.Unauthorized.Code, bizerr.Unauthorized.Error(), NilData)
}

// Forbidden 权限不足
func Forbidden(c *gin.Context) {
	Custom(c, http.StatusOK, bizerr.Forbidden.Code, bizerr.Forbidden.Error(), NilData)
}

func Custom(c *gin.Context, httpCode, bizCode int, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code":       bizCode,
		"msg":        msg,
		"data":       data,
		"request_id": log.RequestIdVal(c),
	})
}

// ValidateErr 表单解析错误响应
func ValidateErr(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator错误
		Unknown(c, err)
		return
	}

	// validator错误进行翻译
	details := removeTopStruct(errs.Translate(trans))

	// 随机返回校验错误中的一条到 msg 字符串
	msg := "param error"
	for _, v := range details {
		msg = v
		break
	}

	Custom(c, http.StatusOK, bizerr.Param.Code, msg, details)
	return
}
