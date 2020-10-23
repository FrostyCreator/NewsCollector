package controller

import (
	"context"
	"github.com/FrostyCreator/NewsCollector/model"
	"github.com/FrostyCreator/NewsCollector/service"
	"log"
)

type NewsController struct {
	ctx			context.Context
	newsRepo	service.NewsRepository
}

func NewNewsController(ctx context.Context, newsRep service.NewsRepository) *NewsController{
	return &NewsController{
		ctx: ctx,
		newsRepo: newsRep,
	}
}

// GetAllNews Вернуть все новости из бд
func (ctr *NewsController) GetAllNewsFromDB() (*[]model.OneNews, error) {
	news, err := ctr.newsRepo.GetAllNews(ctr.ctx);
	if err != nil {
		return nil, err
	}
	return news, err
}

// UpdateAllNews Обновить все новости
func (ctr *NewsController) UpdateAllNews() error {
	news, err := ctr.getAllNewsFromSites()
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = ctr.newsRepo.UpdateSliceNews(ctr.ctx, news)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil;
}

// DeleteNewsById Удалить новость по id
func (ctr *NewsController) DeleteNewsById(id int) (error) {
	err := ctr.newsRepo.DeleteNewsById(ctr.ctx, id)
	if err != nil {
		log.Fatal("Ошибка при удалении новости с id - ", id )
		return err
	}

	return nil
}