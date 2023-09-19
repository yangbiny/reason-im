package caller

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/logger"
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
	data, err := function(req)
	if err != nil {
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
