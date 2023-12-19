package domain

import (
	"sync"
	"zlink/model"

	"go.uber.org/zap"
)

type Store struct {
	log *zap.Logger
	ms  *model.MemoryStore
	sc  *model.CountStore
}

type IStore interface {
	GetCounter() int
	UrlStore(map[string]string)
	GetUrl(path string) (string, error)
}

func NewStore(logger *zap.Logger, ms *model.MemoryStore, sc *model.CountStore) *Store {
	return &Store{
		log: logger,
		ms:  ms,
		sc:  sc,
	}
}

func NewMemoryStore(data map[string]string, mt *sync.Mutex) *model.MemoryStore {

	return &model.MemoryStore{
		Data: data,
		Mt:   mt,
	}
}

func NewCountStore(count int, mt *sync.Mutex) *model.CountStore {
	return &model.CountStore{
		Count: count,
		Mt:    *mt,
	}
}
