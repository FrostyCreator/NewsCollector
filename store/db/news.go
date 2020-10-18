package db

import (
	"context"
	"fmt"

	"github.com/FrostyCreator/NewsCollector/model"

	"github.com/go-pg/pg/v10"
)

type NewsPgRepo struct {
	db *PgDB
}

func NewNewsRepo(db *PgDB) * NewsPgRepo {
	return &NewsPgRepo{db: db}
}

func (repo *NewsPgRepo) GetNews(ctx context.Context) (*[]model.OneNews, error) {
	news := &[]model.OneNews{}
	err := repo.db.Model(news).
		Select()
	if err != nil {
		if err == pg.ErrNoRows { //not found
			return nil, nil
		}
		return nil, err
	}
	return news, nil
}

func (repo *NewsPgRepo) GetOneNewaById(ctx context.Context, id int) (*model.OneNews, error) {
	oneNews := &model.OneNews{}
	err := repo.db.Model(oneNews).
		Where("id = ?", id).
		Select()
	if err != nil {
		if err == pg.ErrNoRows { //not found
			return nil, nil
		}
		return nil, err
	}
	return oneNews, nil
}

func (repo *NewsPgRepo) CreateNews(ctx context.Context, oneNews *model.OneNews) (*model.OneNews, error) {
	_, err := repo.db.Model(oneNews).
		Insert()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("new news - ", oneNews)
	return oneNews, nil
}

func (repo *NewsPgRepo) UpdateNews(ctx context.Context, oneNews *model.OneNews) (*model.OneNews, error) {
	_, err := repo.db.Model(oneNews).
		WherePK().
		Returning("*").
		Update()
	if err != nil {
		if err == pg.ErrNoRows { //not found
			return nil, nil
		}
		return nil, err
	}

	return oneNews, nil
}

func (repo *NewsPgRepo) DeteleOneNewById(ctx context.Context, id int) error {
	_, err := repo.db.Model((*model.OneNews)(nil)).
		Where("id = ?", id).
		Delete()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}