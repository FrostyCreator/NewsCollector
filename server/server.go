package server

import (
	"context"
	"log"

	"github.com/FrostyCreator/NewsCollector"
	"github.com/FrostyCreator/NewsCollector/store"
)

type Server struct {
	context 	context.Context
	config 		*NewsCollector.Config
	Router
	NewsPgRepo store.NewsRepository
}

// Init returns new server instance
func Init(ctx context.Context, config *NewsCollector.Config, db store.NewsRepository, r Router, addr string) (*Server, error) {
	s := &Server{
		context:	ctx,
		config:		config,
		Router:		r,
		NewsPgRepo:	db,
	}

	s.Router.routes()

	if err := s.Router.router.Start(addr); err != nil {
		log.Println(err)
		return nil, err
	}
	return s, nil
}