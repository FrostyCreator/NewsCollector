package server

import (
	"context"
	"log"

	"github.com/FrostyCreator/NewsCollector/config"
	"github.com/FrostyCreator/NewsCollector/store/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	context 	context.Context
	config 		*config.Config
	router		*gin.Engine
	db 			*db.PgDB
}

// Init returns new server instance
func Init(ctx context.Context, config *config.Config, db *db.PgDB, addr string) (*Server, error) {
	router := gin.Default()
	s := &Server{
		context:	ctx,
		config:		config,
		router:		router,
		db:			db,
	}
	s.routes()
	if err := s.router.Run(addr); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return s, nil
}