package utils

import (
	"net/http"
	"zlink/model"

	"github.com/gin-gonic/gin"
)

func BindError(ctx *gin.Context, err error) {
	errResp := model.ErrorResponse{
		Error:   err.Error(),
		ErrCode: http.StatusBadRequest,
	}

	ctx.JSON(errResp.ErrCode, errResp)
}

func ErrorResponse(ctx *gin.Context, err error, status int) {
	errResp := model.ErrorResponse{
		Error:   err.Error(),
		ErrCode: status,
	}

	ctx.JSON(errResp.ErrCode, errResp)
}
