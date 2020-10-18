package server

import (
	"log"
	"net/http"

	"github.com/FrostyCreator/NewsCollector/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	controller	*controller.NewsController
	router		*gin.Engine
}

func NewRouter(ctrl *controller.NewsController) *Router {
	return &Router{
		controller: ctrl,
		router: gin.Default(),
	}
}

//routes lists routes for our HTTP server
func (r Router) routes() {
	r.router.Use(LiberalCORS)

	r.router.GET("/", func(context *gin.Context) {
		r.controller.Test(context)
	})

	r.router.GET("/update", func(context *gin.Context) {
		err := r.controller.UpdateNews(context)
		if err != nil {
			log.Fatal(err)
			context.JSON(500, "")
		}
		context.JSON(200, "")
	})

	r.router.GET("/getNews", func(context *gin.Context) {
		news, err := r.controller.GetAllNews(context)
		if err != nil {
			log.Fatal(err)
			context.JSON(500, "")
		}

		context.JSON(200, *news)
	})
}

// LiberalCORS CORS settings
func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}

