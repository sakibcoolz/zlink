package service

import (
	"zlink/domain"
	"zlink/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service struct {
	log   *zap.Logger
	store domain.IStore
}

type IService interface {
	AddUrl(ctx *gin.Context, addUrl model.AddUrl) (string, error)
	GetUrl(path string) (string, error)
}

func NewService(logger *zap.Logger, store domain.IStore) *Service {
	return &Service{
		log:   logger,
		store: store,
	}
}
