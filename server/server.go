package server

import (
	"lensent/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := Router()
	r.Run()
}

func Router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controllers.UsersController{}
		u.GET("", ctrl.Get)
		u.POST("", ctrl.Post)
		u.GET("/:sub", ctrl.GetBySub)
		u.PUT("/:sub", ctrl.Put)
		u.DELETE("/:sub", ctrl.Delete)
	}

	s := r.Group("/sentenses")
	{
		ctrl := controllers.SentensesController{}
		s.GET("", ctrl.Get)
		s.POST("", ctrl.Post)
		s.GET("/bysub/:sub", ctrl.GetBySub)
		s.GET("/bysentenseid/:sentense_id", ctrl.GetBySentenseId)
		s.PUT("/:sentense_id", ctrl.Put)
		s.DELETE("/:sentense_id", ctrl.Delete)
	}

	return r
}
