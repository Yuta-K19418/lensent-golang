package server

import (
	"lensent/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Get)
		u.POST("", ctrl.Post)
		u.GET("/:sub", ctrl.GetBySub)
		u.PUT("/:sub", ctrl.Put)
		u.DELETE("/:sub", ctrl.Delete)
	}

	return r
}
