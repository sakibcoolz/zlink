package controller

import (
	"errors"
	"net/http"
	"strconv"
	"zlink/utils"

	"github.com/gin-gonic/gin"
)

func (c *Controller) MostVisit(ctx *gin.Context) {
	count, err := strconv.Atoi(ctx.Param("count"))
	if err != nil {
		utils.BindError(ctx, errors.New("worng path"))

		return
	}

	countUrlMap := c.service.MostVisitUrl(count)

	ctx.JSON(http.StatusOK, countUrlMap)
}
