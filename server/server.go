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
		u.GET("/:id", ctrl.GetByID)
		u.PUT("/:id", ctrl.Put)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
