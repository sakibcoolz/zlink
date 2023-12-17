package domain

import (
	"sync"
	"zlink/model"

	"go.uber.org/zap"
)

type Store struct {
	log *zap.Logger
	ms  *model.MemoryStore
}

type IStore interface {
}

func NewStore(logger *zap.Logger, ms *model.MemoryStore) *Store {
	return &Store{
		log: logger,
		ms:  ms,
	}
}

func NewMemoryStore(data map[string]string, mt *sync.Mutex) *model.MemoryStore {

	return &model.MemoryStore{
		Data: data,
		Mt:   mt,
	}
}
