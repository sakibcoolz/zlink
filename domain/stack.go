package domain

import (
	"slices"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Store) SetStack(ctx *gin.Context, path string) {
	if _, err := s.GetUrl(ctx, path); err != nil {
		s.log.Error(ctx, "path not found in NewMemoryStore", zap.Error(err))

		return
	}

	s.collections.Mt.Lock()
	defer s.collections.Mt.Unlock()
	if !slices.Contains(s.collections.Collections.URLs, path) {
		s.collections.Collections.URLs = append(s.collections.Collections.URLs, path)
		s.collections.Collections.Counts = append(s.collections.Collections.Counts, 1)

		return
	}

	idx := slices.Index(s.collections.Collections.URLs, path)
	s.collections.Collections.Counts[idx] += 1
}

func (s *Store) GetMostUrl(ctx *gin.Context, top int) map[string]int {
	urlMap := make(map[string]int)
	s.collections.Mt.Lock()
	data := s.collections.Collections
	s.collections.Mt.Unlock()

	urls := data.URLs
	countData := data.Counts

	if len(urls) != len(countData) {
		return nil
	}

	for i := 0; i < len(countData); i++ {
		for j := 0; j < len(countData); j++ {
			if countData[i] > countData[j] {
				countData[j], countData[i] = countData[i], countData[j]
				urls[j], urls[i] = urls[i], urls[j]
			}
		}
	}

	for i := 0; i < len(urls[:top]); i++ {
		urlMap[urls[i]] = countData[i]
	}

	return urlMap
}
