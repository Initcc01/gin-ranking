package controllers

import "github.com/gin-gonic/gin"

type JsonStruct struct {
	//首字母大写无法返给前段，若想保留字段名的大小写形式，可以使用json标签来定义字段的JSON序列化名称
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

// ReturnSuccess 正确返回
func ReturnSuccess(ctx *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	ctx.JSON(200, json)
}

// ReturnError 错误返回1
func ReturnError(ctx *gin.Context, code int, msg interface{}) {
	json := &JsonStruct{
		Code: code,
		Msg:  msg,
	}
	ctx.JSON(200, json)
}
