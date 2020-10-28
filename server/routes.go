package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/FrostyCreator/NewsCollector/controller"
	"github.com/FrostyCreator/NewsCollector/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	controller	*controller.NewsController
	router		*echo.Echo
}

func NewRouter(ctrl *controller.NewsController) *Router {
	return &Router{
		controller: ctrl,
		router: 	echo.New(),
	}
}

//routes Lists routes for our HTTP server
func (r Router) routes() {
	r.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	r.router.GET("/update", func(context echo.Context) error {
		err := r.controller.UpdateAllNews()
		if err != nil {
			log.Println(err)
			return context.JSON(http.StatusInternalServerError, model.Message{
				"Произошла ошибка при обновлении данных",
			})
		}
		return context.JSON(http.StatusOK, model.Message{
			Message: "Данные обновлены",
		})
	})

	r.router.GET("/getNews", func(context echo.Context) error {
		news, err := r.controller.GetAllNewsFromDB()
		if err != nil {
			log.Println(err)
			return context.JSON(http.StatusInternalServerError, model.Message{
				Message: "Произошла ошибка при получении данных",
			})
		}

		return context.JSON(http.StatusOK, *news)
	})

	r.router.DELETE("/delete/:id", func(context echo.Context) error {
		v := context.Param("id")

		id, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
			return context.JSON(http.StatusBadRequest, model.Message {
				Message : "Введен неверный id",
			})
		}

		err = r.controller.DeleteNewsById(id)
		if err != nil {
			log.Println(err)
			return context.JSON(http.StatusBadRequest, model.Message {
				Message: "Новости с таким id не существует",
			})
		}

		return context.JSON(http.StatusOK, model.Message {
			Message: "Новость удалена",
		})
	})
}