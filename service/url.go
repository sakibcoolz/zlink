package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"zlink/model"
	"zlink/utils"

	"github.com/gin-gonic/gin"
)

func (s *Service) AddUrl(ctx *gin.Context, addUrl model.AddUrl) (string, error) {
	urlStr := make(chan string)
	counter := make(chan int)

	if utils.UrlNotAllowed(addUrl.URL) {
		err := errors.New("this url not allowed")
		s.log.Error(err.Error())

		return "", err
	}

	addUrl.URL = utils.ModifiyUrl(addUrl.URL)

	go func(urlStr chan string) {
		urlStr <- utils.UrlPath()
	}(urlStr)

	go func(counter chan int) {
		counter <- s.store.GetCounter()
	}(counter)

	url := fmt.Sprintf("%s%s", <-urlStr, strconv.Itoa(<-counter))

	go s.store.UrlStore(map[string]string{url: addUrl.URL})

	return fmt.Sprintf("http://%s:%s/%s", os.Getenv("SERVICEHOST"),
		os.Getenv("SERVICEPORT"),
		url), nil
}

func (s *Service) GetUrl(path string) (string, error) {

	return s.store.GetUrl(path)
}
