package service

import (
	"zlink/domain"
	"zlink/log"
	"zlink/model"

	"github.com/gin-gonic/gin"
)

type Service struct {
	log   *log.Log
	store domain.IStore
}

type IService interface {
	AddUrl(ctx *gin.Context, addUrl model.AddUrl) (string, error)
	GetUrl(ctx *gin.Context, path string) (string, error)
	MostVisitUrl(ctx *gin.Context, count int) map[string]int
}

func NewService(logger *log.Log, store domain.IStore) *Service {
	return &Service{
		log:   logger,
		store: store,
	}
}
