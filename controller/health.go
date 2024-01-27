package controller

import (
	"net/http"
	"os"
	"zlink/literals"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Health(ctx *gin.Context) {
	c.log.Info(ctx, "health status API")

	ctx.JSON(http.StatusOK, gin.H{"status": os.Getenv(literals.VERSION)})
}
