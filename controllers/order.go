package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetList(ctx *gin.Context) {
	cid := ctx.PostForm("cid")
	name := ctx.DefaultPostForm("name", "wangwu") //空值时会自动赋一个默认值
	ReturnSuccess(ctx, 0, cid, name, 1)
}
