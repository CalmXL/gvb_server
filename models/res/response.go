package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(0, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithMessage(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMessage(data any, msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailWithCode(code int, c *gin.Context) {
	msg, ok := ErrorMap[ErrorCode(code)]
	if ok {
		Result(code, map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, "未知错误", c)
}
