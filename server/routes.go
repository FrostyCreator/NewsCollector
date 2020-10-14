package server

import (
	"github.com/gin-gonic/gin"
)

//routes lists routes for our HTTP server
func (s *Server) routes() {
	s.router = gin.Default()

	s.router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}