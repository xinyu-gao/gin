package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResult struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

const(
	OK = http.StatusOK
	FAIL = http.StatusInternalServerError
)

func Ok(c *gin.Context, data interface{}){
	c.JSON(OK, HttpResult{
		Code: http.StatusOK,
		Msg: "操作成功",
		Data: data,
	})
}
func ParamError(c *gin.Context, msg string){
	c.JSON(OK, HttpResult{
		Code: http.StatusBadRequest,
		Msg: msg,
	})
}
func Error(c *gin.Context, msg string){
	c.JSON(OK, HttpResult{
		Code: http.StatusInternalServerError,
		Msg: msg,
	})
}

func MethodNotAllowed(c *gin.Context){
	c.JSON(OK, HttpResult{
		Code: http.StatusMethodNotAllowed,
		Msg: "method not allowed",
	})
}

func NotFound(c *gin.Context){
	c.JSON(OK, HttpResult{
		Code: http.StatusNotFound,
		Msg: "404 not found",
	})
}

func Unauthorized(c *gin.Context,){
	c.JSON(OK , HttpResult{
		Code: http.StatusUnauthorized ,
		Msg: "need unauthorized",
	})
}

func Forbidden(c *gin.Context){
	c.JSON(OK , HttpResult{
		Code: http.StatusForbidden ,
		Msg: "need permission",
	})
}