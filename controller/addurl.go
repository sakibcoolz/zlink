package controller

import (
	"net/http"
	"zlink/model"
	"zlink/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// Add accept json body and filter for validating struct.
func (c *Controller) AddUrl(ctx *gin.Context) {
	var addUrl model.AddUrl

	if err := ctx.ShouldBindBodyWith(&addUrl, binding.JSON); err != nil {
		c.log.Error("incorrect json bind with struct", zap.Error(err))

		utils.BindError(ctx, err)

		return
	}

	if err := c.validate.Struct(&addUrl); err != nil {
		c.log.Error("invalid with struct", zap.Error(err))

		utils.BindError(ctx, err)

		return
	}

	url, err := c.service.AddUrl(ctx, addUrl)
	if err != nil {
		c.log.Error(err.Error())

		utils.ErrorResponse(ctx, err, http.StatusBadRequest)

		return
	}

	ctx.String(http.StatusOK, url)
}
