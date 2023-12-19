package utils

import (
	"net/http"
	"zlink/model"

	"github.com/gin-gonic/gin"
)

func StructBindError(ctx *gin.Context, err error) {
	errResp := model.ErrorResponse{
		Error:   err.Error(),
		ErrCode: http.StatusBadRequest,
	}

	ctx.JSON(errResp.ErrCode, errResp)
}
