package service

import (
	"context"
	"github.com/FrostyCreator/NewsCollector/model"
)

type NewsRepository interface {
	GetNews(context.Context) (*[]model.OneNews, error)
	GetOneNewaById(context.Context, int) (*model.OneNews, error)
	CreateNews(context.Context, *model.OneNews) (*model.OneNews, error)
	UpdateNews(context.Context, *model.OneNews) (*model.OneNews, error)
	DeteleOneNewById(context.Context, int) error
}

