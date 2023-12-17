package model

import "sync"

type AddUrl struct {
	URL string `json:"url" validate:"required,url"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	ErrCode int    `json:"err_code"`
}

type MemoryStore struct {
	Data map[string]string
	Mt   *sync.Mutex
}
