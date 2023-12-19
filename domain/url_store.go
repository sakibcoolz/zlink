package domain

import (
	"errors"

	"go.uber.org/zap"
)

func (s *Store) UrlStore(data map[string]string) {
	defer func() {
		if rec := recover(); rec != nil {
			s.log.Error("Recovered from panic", zap.Any("err", rec))
		}
	}()
	for idx, val := range data {
		s.ms.Mt.Lock()
		s.ms.Data[idx] = val
		s.ms.Mt.Unlock()
	}
}

func (s *Store) GetUrl(path string) (string, error) {
	defer func() {
		if rec := recover(); rec != nil {
			s.log.Error("Recovered from panic", zap.Any("err", rec))
		}
	}()
	s.ms.Mt.Lock()
	defer s.ms.Mt.Unlock()
	val, ok := s.ms.Data[path]
	if !ok {
		err := errors.New("no url mapped for path")

		s.log.Error(err.Error())

		return val, err
	}

	return val, nil
}
