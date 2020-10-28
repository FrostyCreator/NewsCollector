package controller

import (
	"context"
	"log"

	"github.com/FrostyCreator/NewsCollector/model"
	"github.com/FrostyCreator/NewsCollector/store"
)

type NewsController struct {
	ctx      context.Context
	newsRepo store.NewsRepository
}

func NewNewsController(ctx context.Context, newsRep store.NewsRepository) *NewsController{
	return &NewsController{
		ctx: 		ctx,
		newsRepo: 	newsRep,
	}
}

// GetAllNews Вернуть все новости из бд
func (ctr *NewsController) GetAllNewsFromDB() (*[]model.OneNews, error) {
	news, err := ctr.newsRepo.GetAllNews(ctr.ctx);
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return news, err
}

// UpdateAllNews Обновить все новости
func (ctr *NewsController) UpdateAllNews() error {
	news, err := ctr.getAllNewsFromSites()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = ctr.newsRepo.UpdateSliceNews(ctr.ctx, news)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil;
}

// DeleteNewsById Удалить новость по id
func (ctr *NewsController) DeleteNewsById(id int) (error) {
	err := ctr.newsRepo.DeleteNewsById(ctr.ctx, id)
	if err != nil {
		log.Println("Ошибка при удалении новости с id - ", id )
		return err
	}

	return nil
}