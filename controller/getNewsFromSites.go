package controller

import (
	"log"
	"strings"
	"time"

	"github.com/FrostyCreator/NewsCollector/model"

	"github.com/gocolly/colly/v2"
)

// getAllNewsFromSites Спарсить новости с сайтов
func (ctr *NewsController) getAllNewsFromSites() (*[]model.OneNews, error) {
	news	:= new([]model.OneNews)
	var id	= new(int)
	*id		= 1

	sliceFunc := []func(chan *[]model.OneNews, *int) {
		getNewsFromPerm59,
		getNewsFromPermkrai,
		getNewsFromAif,
	}
	chNews := make(chan *[]model.OneNews, len(sliceFunc))

	for _, f := range sliceFunc {
		go f(chNews, id);
	}

	var sliceNews *[]model.OneNews
	for i := 0; i < len(sliceFunc); i++ {
		sliceNews = <-chNews
		if  sliceNews != nil {
			for _, n := range *sliceNews {
				*news = append(*news, n)
			}
		} else {
			log.Println("Ошибка при парсинге сайта")
		}
	}

	*id = 1
	return news, nil;
}

// getNewsFromPerm59 получить все новости с сайта https://59.ru
func getNewsFromPerm59(ch chan *[]model.OneNews, id *int) {
	url		:= "https://59.ru/search/?keywords=%D0%9F%D0%93%D0%90%D0%A2%D0%A3&sort=weight"
	c		:= colly.NewCollector()
	news	:= new([]model.OneNews)

	c.OnHTML(".central-column-container > div > article > a > img", func(e *colly.HTMLElement) {
		*news = append(*news, model.OneNews{
			ID:			*id,
			ImageSrc:	e.Attr("src"),
		})
		*id++
	})
	c.OnHTML(".central-column-container > div > article > div > div > h2 > a", func(e *colly.HTMLElement) {
		(*news)[e.Index].Header	= e.Attr("title")
		(*news)[e.Index].Site	= "https://59.ru"
		(*news)[e.Index].URL	= e.Attr("href")
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

// getNewsFromAif Парсинг новостей с сайта https://perm.aif.ru
func getNewsFromAif(ch chan *[]model.OneNews, id *int) {
	url := "https://perm.aif.ru/search?text=%D0%9F%D0%93%D0%90%D0%A2%D0%A3"
	c := colly.NewCollector()
	news := new([]model.OneNews)

	c.OnHTML(".list_item > div > a", func(e *colly.HTMLElement) {
		*news = append(*news, model.OneNews{
			ID:			*id,
			Site:		"https://perm.aif.ru",
			URL: 		strings.Replace(e.Attr("href"), "https://perm.aif.ru", "", 1),
		})
		*id++
	})
	c.OnHTML(".list_item > div > a > h3", func(e *colly.HTMLElement) {
		(*news)[e.Index].Header = e.Text
	})
	c.OnHTML(".list_item > div > .text_box__date", func(e *colly.HTMLElement) {
		d := e.Text[6:10] + "-" + e.Text[3:5] + "-" + e.Text[:2] + "T" + e.Text[11:]
		(*news)[e.Index].Date, _ = time.Parse("2006-01-02T15:04", d)
	})
	c.OnHTML(".list_item > div > span:last-child", func(e *colly.HTMLElement) {
		(*news)[e.Index].Description = e.Text
	})
	c.OnHTML(".list_item > a > img", func(e *colly.HTMLElement) {
		(*news)[e.Index].ImageSrc = e.Attr("src")
	})

	if err := c.Visit(url); err != nil {
		log.Println("Ошибка во время парсинга сайта -", url)
		ch <- nil
	}

	ch <- news
}