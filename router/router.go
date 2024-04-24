package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})
		user.POST("/list", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user list")
		})
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}
	return r
}