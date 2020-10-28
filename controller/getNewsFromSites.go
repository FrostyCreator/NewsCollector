package controller

import (
	"github.com/FrostyCreator/NewsCollector/model"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
	"time"
)

// getAllNewsFromSites Спарсить все новости с сайтов
func (ctr *NewsController) getAllNewsFromSites() (*[]model.OneNews, error) {
	news := new([]model.OneNews)
	var id = new(int)
	*id = 1

	chPerm59 := make(chan *[]model.OneNews)
	go getNewsFromPerm59(chPerm59, id);

	chPermkrai := make(chan *[]model.OneNews)
	go getNewsFromPermkrai(chPermkrai, id)

	n := *<-chPerm59
	if n != nil {
		for _, n := range n {
			*news = append(*news, n)
		}
	} else {
		log.Println("Ошибка при парсинге сайта https://59.ru")
	}

	n = *<-chPermkrai
	if  n != nil {
		for _, n := range n {
			*news = append(*news, n)
		}
	} else {
		log.Println("Ошибка при парсинге сайта https://www.permkrai.ru")
	}

	*id = 1
	return news, nil;
}

// getNewsFromPerm59 получить все новости с сайта https://59.ru
func getNewsFromPerm59(ch chan *[]model.OneNews, id *int) {
	url := "https://59.ru/search/?keywords=%D0%9F%D0%93%D0%90%D0%A2%D0%A3&sort=weight"
	c := colly.NewCollector()

	news := new([]model.OneNews)
	c.OnHTML(".central-column-container > div > article > a > img", func(e *colly.HTMLElement) {
		*news = append(*news, model.OneNews{
			ID:       *id,
			ImageSrc: e.Attr("src"),
		})
		*id++
	})
	c.OnHTML(".central-column-container > div > article > div > div > h2 > a", func(e *colly.HTMLElement) {
		(*news)[e.Index].Header = e.Attr("title")
		(*news)[e.Index].Site = "https://59.ru"
		(*news)[e.Index].URL = e.Attr("href")
	})
	c.OnHTML(".central-column-container > div > article > div > div > div > time", func(e *colly.HTMLElement) {
		date := strings.Replace(e.Attr("datetime"), " ", "T", 1) + ".000Z"
		(*news)[e.Index].Date, _ = time.Parse("2006-01-02T15:04:05.000Z", date)
	})
	c.OnHTML(".central-column-container > div > article > div > div > p > a > span", func(e *colly.HTMLElement) {
		(*news)[e.Index].Description = e.Text
	})
	if err := c.Visit(url); err != nil {
		log.Println("Ошибка во время парсинга сайта -", url)
		ch <- nil
	}

	ch <- news
}

// getNewsFromPermkrai Парсинг новостей с сайта https://www.permkrai.ru
func getNewsFromPermkrai(ch chan *[]model.OneNews, id *int) {
	url := "https://www.permkrai.ru/search/?q=%D0%9F%D0%93%D0%90%D0%A2%D0%A3&category=NEWS"
	c := colly.NewCollector()
	news := new([]model.OneNews)

	c.OnHTML(".download-block_title", func(e *colly.HTMLElement) {
		*news = append(*news, model.OneNews{
			ID:			*id,
			Header:		e.Text,
			Site:		"https://www.permkrai.ru",
			URL: 		e.Attr("href"),
			ImageSrc:	"https://luxury-plitka.ru/img/noimage.png",
		})

		*id++
	})
	c.OnHTML(".download-block_header > .date", func(e *colly.HTMLElement) {
		date := e.Text[6:] + "-" + e.Text[3:5] + "-" + e.Text[:2]
		(*news)[e.Index].Date, _ = time.Parse("2006-01-02", date)
	})
	c.OnHTML(".download-block > p", func(e *colly.HTMLElement) {
		(*news)[e.Index].Description = e.Text
	})

	if err := c.Visit(url); err != nil {
		log.Println("Ошибка во время парсинга сайта -", url)
		ch <- nil
	}

	ch <- news
}