package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"zlink/domain"
	"zlink/log"
	"zlink/model"
	"zlink/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestMostVisit(t *testing.T) {
	os.Setenv("SERVICEHOST", "localhost")
	os.Setenv("SERVICEPORT", "1000")

	logger := log.New()

	validate := validator.New(validator.WithRequiredStructEnabled())

	memStore := domain.NewMemoryStore(map[string]string{"xyz": "xyz", "abc": "abc", "sfd": "sfd"}, new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{URLs: []string{"xyz", "abc", "sfd"}, Counts: []int{12, 34, 56}}, new(sync.Mutex))

	store := domain.NewStore(logger, memStore, cntStore, mapRevStore, collectCount)

	service := service.NewService(logger, store)

	controller := NewController(logger, service, validate)

	router := gin.Default()

	router.GET("/mostvisit/:count", controller.MostVisit)

	t.Run("MostVisit_Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/mostvisit/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.JSONEq(t, `{"sfd": 56}`, w.Body.String())
	})

	t.Run("MostVisit_InvalidCount", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/mostvisit/invalid_count", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "wrong path")
	})

}
