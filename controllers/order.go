package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

// Search 结构体中要指定json字段，否则首字母大写的话是匹配不上前端数据的，前端返回的字段也不会是首字母大写的
type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetList(ctx *gin.Context) {
	//cid := ctx.PostForm("cid")
	//name := ctx.DefaultPostForm("name", "wangwu") //空值时会自动赋一个默认值
	//ReturnSuccess(ctx, 0, cid, name, 1)

	//param := make(map[string]any)
	//err := ctx.BindJSON(&param)

	search := &Search{}
	err := ctx.BindJSON(&search)
	if err == nil {
		ReturnSuccess(ctx, 0, search.Name, search.Cid, 1)
		return
	}
	ReturnError(ctx, 4001, gin.H{"err": err})
}
