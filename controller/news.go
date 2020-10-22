package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/FrostyCreator/NewsCollector/model"
	"github.com/FrostyCreator/NewsCollector/service"

	"github.com/gocolly/colly/v2"
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

func (ctr *NewsController) AddNews() error {
	news, err := ctr.getAllNewsFromSites()
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = ctr.newsRepo.CreateSliceNews(ctr.ctx, news)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

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

func (ctr *NewsController) DeleteNewsById(id int) (error) {
	err := ctr.newsRepo.DeleteNewsById(ctr.ctx, id)
	if err != nil {
		log.Fatal("Ошибка при удалении новости с id - ", id )
		return err
	}

	return nil
}

// getNewsFromPerm59 получить все новости с сайта https://59.ru
func getNewsFromPerm59(id *int) (*[]model.OneNews, error) {
	var newsFromPerm59 *model.NewsFromPerm59
	url := "https://newsapi.59.ru/v1/public/jtnews/records/?page=1&pagesize=40&text=ПГАТУ&sort=weight&pageType=search&regionId=59"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = json.Unmarshal(body, &newsFromPerm59)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if newsFromPerm59.ResultData.StatusCode != 200 {
		log.Fatal("url:", url, " status code -", newsFromPerm59.ResultData.StatusCode, ", with error:", newsFromPerm59.ResultData.Error)
	}

	return newsFromPerm59.ConvertToSliceOneNews(id), nil;
}

// getNewsFromProperm получить новости с сайта https://properm.ru/
func getNewsFromProperm(id *int) (*[]model.OneNews, error) {
	url := "https://properm.ru/news/search/?searchString=%D0%9F%D0%93%D0%90%D0%A2%D0%A3"
	newsFromProperm := new([]model.OneNews)
	c := colly.NewCollector()

	c.OnHTML("a.new-news-piece__link", func(e *colly.HTMLElement) {
		*newsFromProperm = append(*newsFromProperm, model.OneNews{
			ID: *id,
			Header: e.Text,
			URL:    e.Attr("href"),
			Site:   "https://properm.ru",
		})
		(*id)++
	})


	if err := c.Visit(url); err != nil {
		log.Fatal("Ошибка во время парсинга сайта -", url)
		return nil, err
	}

	return newsFromProperm, nil
}

func (ctr *NewsController) getAllNewsFromSites() (*[]model.OneNews, error) {
	news := new([]model.OneNews)
	var id *int = new(int)
	*id = 1

	news, err := getNewsFromPerm59(id);
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//	Добавить новости с сайта https://properm.ru/

	sliceNews, err := getNewsFromProperm(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, oneNews := range *sliceNews {
		*news = append(*news, oneNews)
	}

	*id = 1
	return news, err;
}

