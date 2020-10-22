package server

import (
	"log"
	"net/http"
	"strconv"

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

	r.router.GET("/update", func(context *gin.Context) {
		err := r.controller.UpdateAllNews()
		if err != nil {
			log.Fatal(err)
			context.JSON(500, gin.H{
				"message": "Произошла ошибка при обновлении данных",
			})
		}
		context.JSON(200, gin.H{
			"message": "Данные обновлены",
		})
	})

	r.router.GET("/getNews", func(context *gin.Context) {
		news, err := r.controller.GetAllNewsFromDB()
		if err != nil {
			log.Fatal(err)
			context.JSON(500, gin.H{
				"message": "Произошла ошибка при получении данных",
			})
		}

		context.JSON(200, *news)
	})

	r.router.POST("/delete/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if (err != nil) {
			log.Fatal(err)
			context.JSON(500, gin.H{
				"message": "Введен неверный id",
			})
		}

		err = r.controller.DeleteNewsById(id)
		if (err != nil) {
			log.Fatal(err)
			context.JSON(500, gin.H{
				"message": "Новости с таким id не существует",

			})
		}
		context.JSON(200, gin.H{
			"message": "Новость удалена",
		})
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

