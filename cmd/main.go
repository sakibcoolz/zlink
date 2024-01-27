package main

import (
	"fmt"
	"zlink/app"
	"zlink/config"
	"zlink/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger := log.New()

	config := config.NewConfig(logger.Log())

	router := gin.Default()

	router = app.Apps(config, logger, router)

	if err := router.Run(fmt.Sprintf(":%d", config.GetServiceConfig().Port)); err != nil {
		logger.Log().Fatal("Service Error", zap.Error(err))
	}
}
