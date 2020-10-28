package store

import (
	"fmt"
	"github.com/FrostyCreator/NewsCollector"
	"github.com/FrostyCreator/NewsCollector/model"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PgDB struct {
	*pg.DB
}

func Dial(cfg NewsCollector.Config) (*PgDB, error) {
	pgDB := pg.Connect(&pg.Options{
		Addr: cfg.PgAddr,
		User:  cfg.PgUser,
		Password: cfg.PgPassword,
		Database: cfg.PgDb,
	})

	if _, err := pgDB.Exec("SELECT 1"); err != nil{
		log.Println(err)
		return nil, err
	}

	err := createSchema(pgDB)
	if err != nil {
		log.Println(err)
	}

	db := &PgDB{pgDB}

	//go KeepAlivePg(db, cfg)
	fmt.Println("DB work")

	return db, nil
}


func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.OneNews)(nil),
	}

	for _, m := range models {

		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

//func KeepAlivePg(pgDB *PgDB, cfg NewsCollector.Config) {
//	var err error
//	for {
//		// Check if PostgreSQL is alive every 3 seconds
//		time.Sleep(time.Second * 3)
//		lostConnect := false
//		if pgDB == nil {
//			lostConnect = true
//		} else if _, err = pgDB.Exec("SELECT 1"); err != nil {
//			lostConnect = true
//		}
//		if !lostConnect {
//			continue
//		}
//		log.Println("[store.KeepAlivePg] Lost PostgreSQL connection. Restoring...")
//		pgDB, err = Dial(cfg)
//		if err != nil {
//			log.Println(err)
//			continue
//		}
//		log.Println("[store.KeepAlivePg] PostgreSQL reconnected")
//	}
//}