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
		defer s.ms.Mt.Unlock()
		s.ms.Data[idx] = val

		s.SetUrlMapping(val, idx)
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

func (s *Store) SetUrlMapping(url, path string) {
	s.mr.Mt.Lock()
	defer s.mr.Mt.Unlock()
	s.mr.URLRevMapping[url] = path
}

// returns short url against actual url
func (s *Store) GetUrlMapping(url string) string {
	var path string
	s.mr.Mt.Lock()
	defer s.mr.Mt.Unlock()
	path = s.mr.URLRevMapping[url]

	return path
}
