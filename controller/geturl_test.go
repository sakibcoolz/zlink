package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"zlink/domain"
	"zlink/model"
	"zlink/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetUrl(t *testing.T) {
	os.Setenv("SERVICEHOST", "localhost")
	os.Setenv("SERVICEPORT", "1000")
	logger := zap.NewExample()

	validate := validator.New(validator.WithRequiredStructEnabled())

	memStore := domain.NewMemoryStore(map[string]string{"valid_path": "valid_path"}, new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(logger, memStore, cntStore, mapRevStore, collectCount)

	service := service.NewService(logger, store)

	controller := NewController(logger, service, validate)

	router := gin.Default()

	router.GET("/geturl/:path", controller.GetUrl)

	t.Run("GetUrl_Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/geturl/valid_path", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusFound, w.Code)
	})

	t.Run("GetUrl_EmptyPath", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/geturl/", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "404 page not found")
	})

	t.Run("GetUrl_InvalidPath", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/geturl/invalid_path", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "no url mapped for path")
	})
}
