package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrostyCreator/NewsCollector"
	"github.com/FrostyCreator/NewsCollector/controller"
	"github.com/FrostyCreator/NewsCollector/server"
	"github.com/FrostyCreator/NewsCollector/store"
)

func main(){
	if err := run(); err != nil{
		log.Fatal(err)
	}
}


func run() error {
	ctx := context.Background()

	// config
	cfg := NewsCollector.GetConfig()

	// connect to database
	pgDB, err := store.Dial(*cfg)
	if err != nil {
		return err
	}
	defer pgDB.Close()

	newsRepo := store.NewNewsRepo(pgDB)
	newsController := controller.NewNewsController(ctx, newsRepo)
	router := server.NewRouter(newsController)

	// Обновление новостей в бд
	err = newsController.UpdateAllNews()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// create new server instance and run http server
	addr := ":8080"
	_, err = server.Init(ctx, cfg, newsRepo, *router,  addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Running http server on %s\n", addr)

	return nil
}