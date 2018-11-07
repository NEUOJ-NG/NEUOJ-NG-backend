package router

import (
	c "github.com/NEUOJ-NG/NEUOJ-NG-backend/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	r.GET("/ping", c.Ping)
}
