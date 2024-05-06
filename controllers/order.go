package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetList(ctx *gin.Context) {
	//cid := ctx.PostForm("cid")
	//name := ctx.DefaultPostForm("name", "wangwu") //空值时会自动赋一个默认值
	//ReturnSuccess(ctx, 0, cid, name, 1)

	param := make(map[string]any)
	err := ctx.BindJSON(&param)
	if err == nil {

	}
}
