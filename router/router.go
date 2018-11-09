package router

import (
	c "github.com/NEUOJ-NG/NEUOJ-NG-backend/controller"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	// test
	r.GET("/ping", c.Ping)

	// auth
	authMiddleware := middleware.GetGinJWTMiddleware()
	auth := r.Group("/auth")
	{
		auth.POST("/login", c.Login)
		auth.GET("/refresh_token",
			authMiddleware.MiddlewareFunc(),
			c.RefreshToken,
		)
		auth.GET("/ping",
			authMiddleware.MiddlewareFunc(),
			c.Ping,
		)
	}
}
