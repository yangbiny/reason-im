package caller

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/logger"
	"reflect"
)

type ApiRespCode int

const (
	Success      ApiRespCode = 2
	ParamInvalid ApiRespCode = 4
	ServiceError ApiRespCode = 5
)

type BasicParam interface {
	checkParam() string
}

type ApiResp struct {
	Code ApiRespCode `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Call[A, B any](
	function func(req A) (B, error),
	c *gin.Context,
	req A,
) {
	if err := c.BindJSON(req); err != nil {
		logger.Error(c, "bind req has failed", "req", req)
		ResponseWithParamInvalid(c, err.Error())
		return
	}
	renderLoginUserId(c, req)
	data, err := function(req)
	if err != nil {
		logger.ErrorWithErr(c, "execute has failed : ", err)
		c.JSON(wrapWithServiceError(err))
		return
	}
	c.JSON(wrapWithExecuteSuccess(data))
}

func CallWithContext[A, B any](
	function func(c *gin.Context, req A) (B, error),
	c *gin.Context,
	req A,
) {
	if err := c.BindJSON(req); err != nil {
		logger.Error(c, "bind req has failed", "req", req)
		ResponseWithParamInvalid(c, err.Error())
		return
	}
	renderLoginUserId(c, req)
	data, err := function(c, req)
	if err != nil {
		logger.ErrorWithErr(c, "execute has failed : ", err)
		c.JSON(wrapWithServiceError(err))
		return
	}
	c.JSON(wrapWithExecuteSuccess(data))
}

func CallWithParam[A, B any](
	function func(req A) (B, error),
	c *gin.Context,
	req A,
) {
	renderLoginUserId(c, req)
	data, err := function(req)
	if err != nil {
		logger.Error(c, "call function has failed", "req", req, "err", err)
		c.JSON(wrapWithServiceError(err))
		return
	}
	c.JSON(wrapWithExecuteSuccess(data))
}

func wrapWithExecuteSuccess(data any) (int, any) {
	return 200, ApiResp{
		Code: Success,
		Msg:  "",
		Data: data,
	}
}

func wrapWithServiceError(err error) (int, any) {
	return 200, ApiResp{
		Code: ServiceError,
		Msg:  err.Error(),
		Data: nil,
	}
}

func wrapWithParamsInvalid(msg string) (int, any) {
	return 200, ApiResp{
		Code: ParamInvalid,
		Msg:  msg,
		Data: nil,
	}
}

func ResponseWithParamInvalid(c *gin.Context, msg string) {
	c.JSON(400, msg)
}

func renderLoginUserId[A any](c *gin.Context, req A) {
	value, exists := c.Get("login_user_id")
	if exists {
		of := reflect.TypeOf(req)
		var elem = of
		if of.Kind() == reflect.Pointer {
			elem = of.Elem()
		}
		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
			tag := field.Tag.Get("login_user_id")
			if len(tag) > 0 {
				v := reflect.ValueOf(req)
				v.Elem().Field(i).SetInt(reflect.ValueOf(value).Int())
			}
		}
	}
}
