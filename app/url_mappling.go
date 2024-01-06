package app

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"zlink/config"
	"zlink/controller"
	"zlink/domain"
	"zlink/literals"
	"zlink/middleware"
	"zlink/model"
	"zlink/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var (
	StopService = make(chan os.Signal, 1)
)

func Apps(config *config.Config, logger *zap.Logger, router *gin.Engine) *gin.Engine {
	go terminateService(StopService, logger)

	signal.Notify(StopService, syscall.SIGINT, syscall.SIGTERM)

	validate := validator.New(validator.WithRequiredStructEnabled())

	memStore := domain.NewMemoryStore(make(map[string]string), new(sync.Mutex))

	memRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectionStore := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(logger, memStore, cntStore, memRevStore, collectionStore)

	service := service.NewService(logger, store)

	controller := controller.NewController(logger, service, validate)

	configs := cors.DefaultConfig()
	configs.AllowOrigins = []string{"*"}
	configs.AllowHeaders = []string{"*"}
	configs.AllowMethods = []string{"GET", "POST"}

	router.Use(cors.New(configs))

	router.Use(gin.Recovery())

	router.Use(middleware.FilterURL(store))

	router.GET(literals.PATH, controller.GetUrl)

	preapproute := router.Group(literals.V1)

	preapproute.GET(literals.HEALTH, controller.Health)

	preapproute.POST(literals.ADDURL, controller.AddUrl)

	preapproute.GET(literals.MOSTVISIT+literals.COUNT, controller.MostVisit)

	return router
}

func terminateService(stopService chan os.Signal, log *zap.Logger) {
	log.Info("Service Started")

	if _, ok := <-stopService; ok {

		log.Info("Terminating service")

		os.Exit(0)
	}
}
