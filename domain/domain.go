package domain

import (
	"sync"
	"zlink/log"
	"zlink/model"

	"github.com/gin-gonic/gin"
)

type Store struct {
	log         *log.Log
	ms          *model.MemoryStore
	sc          *model.CountStore
	mr          *model.MappingRev
	collections *model.URLCountCollections
}

type IStore interface {
	GetCounter(ctx *gin.Context) int
	UrlStore(ctx *gin.Context, mk map[string]string)
	GetUrl(ctx *gin.Context, path string) (string, error)
	SetUrlMapping(ctx *gin.Context, url, path string)
	// returns short url against actual url
	GetUrlMapping(ctx *gin.Context, url string) string
	SetStack(ctx *gin.Context, path string)
	GetMostUrl(ctx *gin.Context, top int) map[string]int
}

func NewStore(logger *log.Log, ms *model.MemoryStore, sc *model.CountStore,
	mr *model.MappingRev, collections *model.URLCountCollections) *Store {
	return &Store{
		log:         logger,
		ms:          ms,
		sc:          sc,
		mr:          mr,
		collections: collections,
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
		Mt:    mt,
	}
}

func NewMappingRev(urLRevMapping map[string]string,
	mt *sync.Mutex) *model.MappingRev {
	return &model.MappingRev{
		URLRevMapping: urLRevMapping,
		Mt:            mt,
	}
}

func NewUrlCollectionCount(urlCollection model.Collections,
	mt *sync.Mutex) *model.URLCountCollections {
	return &model.URLCountCollections{
		Mt: mt,
		Collections: model.Collections{
			URLs:   urlCollection.URLs,
			Counts: urlCollection.Counts,
		},
	}
}
