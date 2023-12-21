package app

import (
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"zlink/config"
	"zlink/controller"
	"zlink/domain"
	"zlink/literals"
	"zlink/model"
	"zlink/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var (
	StopService = make(chan os.Signal, 1)
)

func Apps(config *config.Config, logger *zap.Logger, router *gin.Engine) *gin.Engine {
	go TerminateService(StopService, logger)

	signal.Notify(StopService, syscall.SIGINT, syscall.SIGTERM)

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

	controller := controller.NewController(logger, service, validate)

	router.Use(gin.Recovery())

	router.Use(CustomMiddleware(store))

	router.GET("/:path", controller.GetUrl)

	preapproute := router.Group("/v1")

	preapproute.GET(literals.HEALTH, controller.Health)

	preapproute.POST(literals.ADDURL, controller.AddUrl)

	preapproute.GET(literals.MOSTVISIT+"/:count", controller.MostVisit)

	return router
}

func CustomMiddleware(store *domain.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		path := c.Request.RequestURI

		UrlCountCollection(store, path)

	}
}

func UrlCountCollection(store *domain.Store, path string) {

	if strings.Contains(path, literals.HEALTH) || strings.Contains(path, literals.ADDURL) || strings.Contains(path, literals.MOSTVISIT) {
		return
	}

	paths := strings.Split(path, "/")[1]

	store.SetStack(paths)
}

func TerminateService(stopService chan os.Signal, log *zap.Logger) {
	log.Info("Service Started")

	if _, ok := <-stopService; ok {

		log.Info("Terminating service")

		os.Exit(0)
	}
}
