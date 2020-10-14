package db

import (
	"log"

	"github.com/FrostyCreator/NewsCollector/config"

	"github.com/go-pg/pg/v10"
)

type PgDB struct {
	*pg.DB
}

func Dial(cfg config.Config) (*PgDB, error) {
	pgDB := pg.Connect(&pg.Options{
		Addr: cfg.PgAddr,
		User:  cfg.PgUser,
		Password: cfg.PgPassword,
		Database: cfg.PgDb,
	})

	if _, err := pgDB.Exec("SELECT 1"); err != nil{
		log.Fatal(err)
		return nil, err
	}

	return &PgDB{pgDB}, nil
}