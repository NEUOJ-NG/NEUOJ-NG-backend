package util

import "github.com/gin-gonic/gin"

type HTTPError struct {
	// in order to extend basic http status code, we multiply it by 10
	// for example, 404 will be extended into 40400
	Code    int    `json:"code" example:"40000"`
	Message string `json:"message" example:"status bad request"`
}

func NewStandardError(ctx *gin.Context, standardStatus int, err error) {
	er := HTTPError{
		Code:    standardStatus * 10,
		Message: err.Error(),
	}
	ctx.JSON(standardStatus, er)
}

func NewExtendedError(ctx *gin.Context, standardStatus int, extStatus int, err error) {
	er := HTTPError{
		Code:    extStatus,
		Message: err.Error(),
	}
	ctx.JSON(standardStatus, er)
}
