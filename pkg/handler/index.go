package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 首页
func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

// api交互通用返回结构
type RespInfo struct {
	Code int8        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回赋值
func result(code int8, msg string, data interface{}) *RespInfo {
	return &RespInfo{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
