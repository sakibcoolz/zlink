package app

import (
	"os"
	"os/signal"
	"syscall"
	"zlink/config"
	"zlink/controller"
	"zlink/domain"
	"zlink/literals"
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

	store := domain.NewStore(logger)

	service := service.NewService(logger, store)

	controller := controller.NewController(logger, service, validate)

	router.Use(gin.Recovery())

	preapproute := router.Group(literals.VERSIONONE)

	preapproute.GET("/health", controller.Health)

	return router
}

func TerminateService(stopService chan os.Signal, log *zap.Logger) {
	log.Info("Service Started")
	select {
	case <-stopService:
		log.Info("Terminating service")

		os.Exit(0)
	}
}
