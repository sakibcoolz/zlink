package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Health(ctx *gin.Context) {
	c.log.Info("health status API")

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
