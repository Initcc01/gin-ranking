package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, 0, "success", "user info", 1)
}
