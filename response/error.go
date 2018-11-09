package response

import "github.com/gin-gonic/gin"

type HTTPError struct {
	// in order to extend basic http status code, we multiply it by 10
	// for example, 404 will be extended into 40400
	Code    int    `json:"code" example:"40000"`
	Message string `json:"message" example:"status bad request"`
}

func NewStandardError(ctx *gin.Context, standardStatus int, err interface{}) {
	var er HTTPError
	if v, ok := err.(error); ok {
		er = HTTPError{
			Code:    standardStatus * 100,
			Message: v.Error(),
		}
	} else if v, ok := err.(string); ok {
		er = HTTPError{
			Code:    standardStatus * 100,
			Message: v,
		}
	} else {
		panic("err must be type of error or string")
	}
	ctx.JSON(standardStatus, er)
}

func NewExtendedError(ctx *gin.Context, standardStatus int, extStatus int, err interface{}) {
	var er HTTPError
	if v, ok := err.(error); ok {
		er = HTTPError{
			Code:    extStatus,
			Message: v.Error(),
		}
	} else if v, ok := err.(string); ok {
		er = HTTPError{
			Code:    extStatus,
			Message: v,
		}
	} else {
		panic("err must be type of error or string")
	}
	ctx.JSON(standardStatus, er)
}
