package domain

import (
	"go.uber.org/zap"
)

type Store struct {
	log *zap.Logger
}

type IStore interface {
}

func NewStore(logger *zap.Logger) *Store {
	return &Store{
		log: logger,
	}
}
