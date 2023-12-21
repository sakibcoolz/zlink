package model

import "sync"

type AddUrl struct {
	URL string `json:"url" validate:"required"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	ErrCode int    `json:"err_code"`
}

type MemoryStore struct {
	Data map[string]string
	Mt   *sync.Mutex
}

type CountStore struct {
	Count int
	Mt    *sync.Mutex
}

type MappingRev struct {
	URLRevMapping map[string]string
	Mt            *sync.Mutex
}

type URLCountCollections struct {
	Mt          *sync.Mutex
	Collections Collections
}

type Collections struct {
	URLs   []string
	Counts []int
}
