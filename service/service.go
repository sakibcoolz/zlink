package service

import (
	"zlink/domain"

	"go.uber.org/zap"
)

type Service struct {
	log   *zap.Logger
	store domain.IStore
}

type IService interface {
}

func NewService(logger *zap.Logger, store domain.IStore) *Service {
	return &Service{
		log:   logger,
		store: store,
	}
}
