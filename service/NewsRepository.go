package service

import (
	"context"

	"github.com/FrostyCreator/NewsCollector/model"
)

type NewsRepository interface {
	// GetNews Получить все новости
	GetAllNews(context.Context) (*[]model.OneNews, error)

	// GetOneNewaById Получить одну новость по id
	GetOneNewsById(context.Context, int) (*model.OneNews, error)

	// CreateNews Добавить одну новость в бд
	CreateNews(context.Context, *model.OneNews) (*model.OneNews, error)

	// CreateSliceNews Добавить список новостей в бд
	CreateSliceNews(context.Context, *[]model.OneNews) (*[]model.OneNews, error)

	// UpdateNews Обновить запись в бд
	UpdateNews(context.Context, *model.OneNews) (*model.OneNews, error)

	// UpdateNews Обновить запись в бд
	UpdateSliceNews(context.Context, *[]model.OneNews) (*[]model.OneNews, error)

	// DeteleOneNewById Удалить одну запись из бд по id
	DeleteNewsById(context.Context, int) error
}