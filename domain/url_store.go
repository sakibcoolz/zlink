package domain

import "go.uber.org/zap"

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
