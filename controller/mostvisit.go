package controller

import (
	"errors"
	"net/http"
	"strconv"
	"zlink/utils"

	"github.com/gin-gonic/gin"
)

// provide top count against count path variable.
func (c *Controller) MostVisit(ctx *gin.Context) {
	count, err := strconv.Atoi(ctx.Param("count"))
	if err != nil {
		utils.BindError(ctx, errors.New("wrong path"))

		return
	}

	countUrlMap := c.service.MostVisitUrl(ctx, count)
	if countUrlMap == nil {
		utils.ErrorResponse(ctx, errors.New("no data available"), http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, countUrlMap)
}
