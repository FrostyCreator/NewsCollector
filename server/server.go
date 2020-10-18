package server

import (
	"context"
	"github.com/FrostyCreator/NewsCollector/service"
	"log"

	"github.com/FrostyCreator/NewsCollector/config"
)

type Server struct {
	context 	context.Context
	config 		*config.Config
	Router
	NewsPgRepo	service.NewsRepository
}

// Init returns new server instance
func Init(ctx context.Context, config *config.Config, db service.NewsRepository, r Router, addr string) (*Server, error) {
	s := &Server{
		context:	ctx,
		config:		config,
		Router:		r,
		NewsPgRepo:	db,
	}

	s.Router.routes()

	if err := s.Router.router.Run(addr); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return s, nil
}