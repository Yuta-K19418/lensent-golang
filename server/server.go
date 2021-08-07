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
	v := r.Group("/api")
	v.BasePath()

	u := r.Group(v.BasePath() + "/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Get)
		u.POST("", ctrl.Post)
		u.GET("/:sub", ctrl.GetBySub)
		u.PUT("/:sub", ctrl.Put)
		u.DELETE("/:sub", ctrl.Delete)
	}

	s := r.Group(v.BasePath() + "/sentenses")
	{
		ctrl := controllers.SentenseController{}
		s.GET("", ctrl.Get)
		s.POST("", ctrl.Post)
		s.GET("/by-sub/:sub", ctrl.GetBySub)
		s.GET("/by-sentense-id/:sentense_id", ctrl.GetBySentenseId)
		s.PUT("/:sentense_id", ctrl.Put)
		s.DELETE("/:sentense_id", ctrl.Delete)
	}

	w := r.Group(v.BasePath() + "/words")
	{
		ctrl := controllers.WordController{}
		w.GET("", ctrl.Get)
		w.POST("", ctrl.Post)
		w.GET("/by-sentense-id/:sentense_id", ctrl.GetBySentenseId)
		w.GET("/by-word-id/:word_id", ctrl.GetByWordId)
		w.PUT("/:word_id", ctrl.Put)
		w.DELETE("/:word_id", ctrl.Delete)
	}

	return r
}
