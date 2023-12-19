package controller

import (
	"errors"
	"net/http"
	"zlink/utils"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetUrl(ctx *gin.Context) {
	path := ctx.Param("path")

	if len(path) == 0 {
		utils.BindError(ctx, errors.New("worng path"))

		return
	}

	url, err := c.service.GetUrl(path)
	if err != nil {
		utils.ErrorResponse(ctx, err, http.StatusInternalServerError)

		return
	}

	ctx.Redirect(http.StatusFound, url)
}
