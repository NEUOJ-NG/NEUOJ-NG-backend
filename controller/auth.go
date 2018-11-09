package controller

import (
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/middleware"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	middleware.GetGinJWTMiddleware().LoginHandler(ctx)
}

func RefreshToken(ctx *gin.Context) {
	middleware.GetGinJWTMiddleware().RefreshHandler(ctx)
}
