package store

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/FrostyCreator/NewsCollector/model"

	"github.com/go-pg/pg/v10"
)

type NewsPgRepo struct {
	db *PgDB
}

func NewNewsRepo(db *PgDB) *NewsPgRepo {
	return &NewsPgRepo{db: db}
}

// GetAllNews получить все новости из БД
func (repo *NewsPgRepo) GetAllNews(ctx context.Context) (*[]model.OneNews, error) {
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

// GetOneNewsById Получить одну новость по id
func (repo *NewsPgRepo) GetOneNewsById(ctx context.Context, id int) (*model.OneNews, error) {
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

// CreateNews Добавить одну новость в БД
func (repo *NewsPgRepo) CreateNews(ctx context.Context, oneNews *model.OneNews) (*model.OneNews, error) {
	_, err := repo.db.Model(oneNews).
		Insert()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return oneNews, nil
}

// CreateSliceNews Добавить список новостей в БД
func (repo *NewsPgRepo) CreateSliceNews(ctx context.Context, news *[]model.OneNews) (*[]model.OneNews, error) {
	_, err := repo.db.Model(news).WherePK().Insert()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return news, nil
}

// UpdateNews Изменить новость
func (repo *NewsPgRepo) UpdateNews(ctx context.Context, oneNews *model.OneNews) (*model.OneNews, error) {

	// Существует ли текущая запись
	exist, err := repo.db.Model(oneNews).WherePK().Exists()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Если существует, то изменить
	if exist {
		_, err = repo.db.Model(oneNews).WherePK().Update()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	// Или добавить
	} else {
		_, err := repo.CreateNews(ctx, oneNews)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return oneNews, nil
}

// UpdateSliceNews Изменить список новостей
func (repo *NewsPgRepo) UpdateSliceNews(ctx context.Context, news *[]model.OneNews) (*[]model.OneNews, error) {
	for _, n := range *news {
		_, err := repo.UpdateNews(ctx, &n)
		if err != nil {
			return nil, err
		}
	}
	return news, nil
}

// DeleteNewsById Удалить новость по id
func (repo *NewsPgRepo) DeleteNewsById(ctx context.Context, id int) error {
	exist, err := repo.db.Model(new(model.OneNews)).Where("id = ?", id).Exists()
	if err != nil {
		log.Println(err)
		return err
	}

	if exist {
		_, err := repo.db.Model(new(model.OneNews)).
			Where("id = ?", id).
			Delete()
		if err != nil {
			if err == pg.ErrNoRows {
				return nil
			}
			return err
		}

		return nil
	} else {
		str := fmt.Sprintf("Записи в id - %s не существует", id)
		log.Println(str)
		return errors.New(str)
	}
}