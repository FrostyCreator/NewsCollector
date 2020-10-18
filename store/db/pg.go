package db

import (
	"fmt"
	"log"

	"github.com/FrostyCreator/NewsCollector/config"
	"github.com/FrostyCreator/NewsCollector/model"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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

	err := createSchema(pgDB)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB work")
	return &PgDB{pgDB}, nil
}


func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.OneNews)(nil),
	}

	for _, m := range models {

		//b, err := db.Model(m).Exists();
		//if  err != nil {
		//	log.Fatal(err)
		//}
		//if b {
		//	continue
		//}

		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}