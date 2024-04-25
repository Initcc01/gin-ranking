package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

func (u UserController) GetUserInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.Param("name")
	//ReturnSuccess(ctx, 0, "success", "user info", 1)
	//ReturnSuccess(ctx, 0, "success", id, 1)
	ReturnSuccess(ctx, 0, name, id, 1)
}

func (u UserController) GetList(ctx *gin.Context) {
	ReturnError(ctx, 4004, "没有相关信息list")
}
