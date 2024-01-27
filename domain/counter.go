package domain

import "github.com/gin-gonic/gin"

func (s *Store) GetCounter(ctx *gin.Context) int {
	counter := 0
	s.sc.Mt.Lock()
	s.sc.Count += 1
	counter = s.sc.Count
	s.sc.Mt.Unlock()
	return counter
}
