package server

import (
	"lensent/controllers"
	"time"

	"github.com/gin-contrib/cors"
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

	// Setting cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"https://lensent.tk/",
		"http://lensent.tk/",
		"https://lensent-nextjs-alb-*.ap-northeast-1.elb.amazonaws.com/",
		"http://lensent-nextjs-alb-*.ap-northeast-1.elb.amazonaws.com/",
	}
	corsConfig.AllowMethods = []string{
		"POST",
		"GET",
		"PUT",
		"DELETE",
	}
	corsConfig.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Origin",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 24 * time.Hour

	r.Use(cors.New(corsConfig))

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
