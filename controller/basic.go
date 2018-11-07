package controller

import (
	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Ping
// @Description just a ping test
// @ID ping
// @Tags basic
// @Accept json
// @Produce text/plain
// @Success 200 {string} string
// @Router /ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(200, "pong")
}
