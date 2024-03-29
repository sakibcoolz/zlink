package middleware

import (
	"strings"
	"zlink/domain"
	"zlink/literals"

	"github.com/gin-gonic/gin"
)

func FilterURL(store *domain.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		urlCountCollection(store, c.Request.RequestURI)
	}
}

func urlCountCollection(store *domain.Store, path string) {
	if strings.Contains(path, literals.HEALTH) || strings.Contains(path, literals.ADDURL) || strings.Contains(path, literals.MOSTVISIT) {
		return
	}

	paths := strings.Split(path, "/")[1]

	store.SetStack(paths)
}
