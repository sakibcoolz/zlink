package service

import (
	"errors"
	"fmt"
	"strconv"
	"zlink/model"
	"zlink/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Service) AddUrl(ctx *gin.Context, addUrl model.AddUrl) (string, error) {
	urlStr := make(chan string)
	counter := make(chan int)

	s.log.Info(ctx, "add url", zap.String("sakib", "sakib"))

	if utils.UrlNotAllowed(addUrl.URL) {
		err := errors.New("this url not allowed")
		s.log.Error(ctx, err.Error())

		return "", err
	}

	addUrl.URL = utils.ModifiyUrl(addUrl.URL)

	if path := s.store.GetUrlMapping(ctx, addUrl.URL); path != "" {
		return utils.UrlMaker(path), nil
	}

	go func(urlStr chan string) {
		urlStr <- utils.UrlPath()
	}(urlStr)

	go func(counter chan int) {
		counter <- s.store.GetCounter(ctx)
	}(counter)

	url := fmt.Sprintf("%s%s", <-urlStr, strconv.Itoa(<-counter))

	go s.store.UrlStore(ctx, map[string]string{url: addUrl.URL})

	return utils.UrlMaker(url), nil
}

func (s *Service) GetUrl(ctx *gin.Context, path string) (string, error) {

	return s.store.GetUrl(ctx, path)
}

func (s *Service) MostVisitUrl(ctx *gin.Context, count int) map[string]int {
	mostUrl := make(map[string]int)
	urlMap := s.store.GetMostUrl(ctx, count)

	for key, value := range urlMap {
		url, err := s.store.GetUrl(ctx, key)
		if err != nil {
			s.log.Log().Error("url not found")

			return nil
		}

		mostUrl[url] = value
	}

	return mostUrl
}
