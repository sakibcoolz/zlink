package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

func TestAddUrl(t *testing.T) {
	os.Setenv("SERVICEHOST", "localhost")
	os.Setenv("SERVICEPORT", "1000")
	logger := zap.NewExample()

	validate := validator.New(validator.WithRequiredStructEnabled())

	memStore := domain.NewMemoryStore(make(map[string]string), new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(logger, memStore, cntStore, mapRevStore, collectCount)

	service := service.NewService(logger, store)

	controller := NewController(logger, service, validate)

	t.Run("AddUrl_Success", func(t *testing.T) {
		jsonPayload := `{"url": "http://example.com"}`

		req := httptest.NewRequest("POST", "/addurl", strings.NewReader(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		controller.AddUrl(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, w.Body.String())
	})

}

func TestAddUrlWrong(t *testing.T) {
	os.Setenv("SERVICEHOST", "localhost")
	os.Setenv("SERVICEPORT", "1000")
	logger := zap.NewExample()

	validate := validator.New(validator.WithRequiredStructEnabled())

	memStore := domain.NewMemoryStore(make(map[string]string), new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(logger, memStore, cntStore, mapRevStore, collectCount)

	service := service.NewService(logger, store)

	controller := NewController(logger, service, validate)

	t.Run("AddUrl_Success", func(t *testing.T) {
		jsonPayload := `{"url": "http://localhost.com"}`

		req := httptest.NewRequest("POST", "/addurl", strings.NewReader(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		controller.AddUrl(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, w.Body.String())
	})

}

func TestAddUrlEmptyStruct(t *testing.T) {
	os.Setenv("SERVICEHOST", "localhost")
	os.Setenv("SERVICEPORT", "1000")
	logger := zap.NewExample()

	validate := validator.New(validator.WithRequiredStructEnabled())

	memStore := domain.NewMemoryStore(make(map[string]string), new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(logger, memStore, cntStore, mapRevStore, collectCount)

	service := service.NewService(logger, store)

	controller := NewController(logger, service, validate)

	t.Run("AddUrl_Success", func(t *testing.T) {
		jsonPayload := `{}`

		req := httptest.NewRequest("POST", "/addurl", strings.NewReader(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		controller.AddUrl(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, w.Body.String())
	})

}

func TestAddUrlNoBody(t *testing.T) {
	os.Setenv("SERVICEHOST", "localhost")
	os.Setenv("SERVICEPORT", "1000")
	logger := zap.NewExample()

	validate := validator.New(validator.WithRequiredStructEnabled())

	memStore := domain.NewMemoryStore(make(map[string]string), new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(logger, memStore, cntStore, mapRevStore, collectCount)

	service := service.NewService(logger, store)

	controller := NewController(logger, service, validate)

	t.Run("AddUrl_Success", func(t *testing.T) {
		jsonPayload := ``

		req := httptest.NewRequest("POST", "/addurl", strings.NewReader(jsonPayload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		controller.AddUrl(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotEmpty(t, w.Body.String())
	})

}
