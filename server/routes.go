package server

import (
	"log"
	"strconv"

	"github.com/FrostyCreator/NewsCollector/controller"

	"github.com/gin-contrib/cors"
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

//routes Lists routes for our HTTP server
func (r Router) routes() {
	r.router.Use(cors.Default())

	r.router.GET("/update", func(context *gin.Context) {
		err := r.controller.UpdateAllNews()
		if err != nil {
			log.Println(err)
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
			log.Println(err)
			context.JSON(500, gin.H{
				"message": "Произошла ошибка при получении данных",
			})
		}

		context.JSON(200, *news)
	})

	r.router.DELETE("/delete/:id", func(context *gin.Context) {
		v := context.Param("id")

		id, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
			context.JSON(500, gin.H {
				"message" : "Введен неверный id",
			})
			return
		}

		err = r.controller.DeleteNewsById(id)
		if err != nil {
			log.Println(err)
			context.JSON(500, gin.H {
				"message" : "Новости с таким id не существует",
			})
			return
		}

		context.JSON(200, gin.H {
			"message": "Новость удалена",
		})
	})
}