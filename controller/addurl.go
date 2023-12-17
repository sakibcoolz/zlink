package controller

import (
	"zlink/model"
	"zlink/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

func (c *Controller) AddUrl(ctx *gin.Context) {
	var addUrl model.AddUrl

	if err := ctx.ShouldBindBodyWith(&addUrl, binding.JSON); err != nil {
		c.log.Error("incorrect json bind with struct", zap.Error(err))

		utils.StructBindError(ctx, err)

		return
	}

	if err := c.validate.Struct(&addUrl); err != nil {
		c.log.Error("invalid with struct", zap.Error(err))

		utils.StructBindError(ctx, err)

		return
	}

}
