package router

import (
	"gin-ranking/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		//user.GET("/info", controllers.UserController{}.GetUserInfo)
		//user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.GET("/info/:id/:name", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)
		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
