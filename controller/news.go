package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FrostyCreator/NewsCollector/model"
	"github.com/FrostyCreator/NewsCollector/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
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



func (ctr *NewsController) Test(ctx *gin.Context) error{
	//var oneNews model.OneNews
	ctx.JSON(200, gin.H{
		"message": "test",
	})

	return nil
}

func (ctr *NewsController) UpdateNews(ctx *gin.Context) error {
	news, err := getNewsFromPerm59();
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _, n := range *news {
		ctr.newsRepo.CreateNews(ctr.ctx, &n)
	}

	return nil;
}

func (ctr * NewsController) GetAllNews(ctx  *gin.Context) (*[]model.OneNews, error) {
	news, err := ctr.newsRepo.GetNews(ctr.ctx);
	if err != nil {
		return nil, err
	}

	fmt.Println(news)

	return news, err
}

func getNewsFromPerm59() (*[]model.OneNews, error) {
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

	return newsFromPerm59.ConvertToSliceOneNews(), nil;
}

