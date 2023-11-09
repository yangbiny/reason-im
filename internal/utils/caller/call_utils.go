package caller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	apierror "github.com/yangbiny/reason-commons/err"
	"reason-im/internal/utils/logger"
	"reflect"
)

var validObj = validator.New()

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
	function func(c *gin.Context, req A) (B, *apierror.ApiError),
	c *gin.Context,
	req A,
) {
	err2 := c.ShouldBindUri(req)
	if err2 != nil {
		logger.Error(c, "bind uri has failed", "req", req)
		c.JSON(wrapWithParamError(err2))
		return
	}
	if err := c.Bind(req); err != nil {
		logger.Error(c, "bind req has failed", "req", req)
		c.JSON(wrapWithParamError(err))
		return
	}

	renderLoginUserId(c, req)

	err2 = validateReq(req)
	if err2 != nil {
		logger.Error(c, "缺少必要参数", "req", req, "error", err2.Error())
		c.JSON(wrapWithParamError(err2))
		return
	}

	data, err := function(c, req)
	if err != nil {
		if err.ApiStatus == apierror.RequestParamError {
			c.JSON(wrapWithParamError(err.Err))
			return
		}
		logger.ErrorWithErr(c, "execute has failed : ", err.Err)
		c.JSON(wrapWithServiceError(err.Err))
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

func wrapWithParamError(err error) (int, any) {
	return 200, ApiResp{
		Code: ParamInvalid,
		Msg:  err.Error(),
		Data: nil,
	}
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
				v2 := v.Elem().Field(i)
				if v2.Kind() == reflect.Ptr {
					ptr := reflect.New(v2.Type().Elem())
					ptr.Elem().SetInt(value.(int64))
					v2.Set(ptr)
				} else {
					v2.SetInt(reflect.ValueOf(value).Int())
				}

				break
			}
		}
	}
}

func validateReq[A any](req A) error {
	err := validObj.Struct(req)
	if err != nil {
		return err
	}
	return nil
}
