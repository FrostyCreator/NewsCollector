package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/FrostyCreator/NewsCollector/model"

	"github.com/gocolly/colly/v2"
)

// getAllNewsFromSites Спарсить все новости с сайтов
func (ctr *NewsController) getAllNewsFromSites() (*[]model.OneNews, error) {
	news := new([]model.OneNews)
	var id *int = new(int)
	*id = 1

	news, err := getNewsFromPerm59(id);
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sliceNews, err := getNewsFromProperm(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, oneNews := range *sliceNews {
		*news = append(*news, oneNews)
	}

	sliceNews, err = getNewsFromPermkrai(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, oneNews := range *sliceNews {
		*news = append(*news, oneNews)
	}

	*id = 1
	return news, err;
}

// getNewsFromPerm59 получить все новости с сайта https://59.ru
func getNewsFromPerm59(id *int) (*[]model.OneNews, error) {
	var newsFromPerm59 *model.NewsFromPerm59
	url := "https://newsapi.59.ru/v1/public/jtnews/records/?page=1&pagesize=40&text=ПГАТУ&sort=weight&pageType=search&regionId=59"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(body, &newsFromPerm59)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if newsFromPerm59.ResultData.StatusCode != 200 {
		log.Println("url:", url, " status code -", newsFromPerm59.ResultData.StatusCode, ", with error:", newsFromPerm59.ResultData.Error)
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
		log.Println("Ошибка во время парсинга сайта -", url)
		return nil, err
	}

	return newsFromProperm, nil
}

// getNewsFromPermkrai Парсинг новостей с сайта https://www.permkrai.ru
func getNewsFromPermkrai(id *int) (*[]model.OneNews, error) {
	url := "https://www.permkrai.ru/search/?q=%D0%9F%D0%93%D0%90%D0%A2%D0%A3"
	newsFromProperm := new([]model.OneNews)
	c := colly.NewCollector()

	c.OnHTML(".download-block > a", func(e *colly.HTMLElement) {
		*newsFromProperm = append(*newsFromProperm, model.OneNews{
			ID: *id,
			Header: e.Text,
			URL:    e.Attr("href"),
			Site:   "https://www.permkrai.ru",
		})
		(*id)++
	})

	if err := c.Visit(url); err != nil {
		log.Println("Ошибка во время парсинга сайта -", url)
		return nil, err
	}

	return newsFromProperm, nil
}