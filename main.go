package main

import (
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/config"
	_ "github.com/NEUOJ-NG/NEUOJ-NG-backend/docs"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/router"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/util"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title NEUOJ-NG-backend API
// @version 1.0
// @description This is API for NEUOJ-NG

// @host localhost:8080
// @BasePath /api/v1

func main() {
	// setup storage dir
	util.SetupStorage()

	// setup log
	util.SetupLog()

	// create Gin Engine with Logger and Recovery middleware
	app := gin.Default()

	// setup CORS
	if len(config.GetConfig().App.AllowOrigin) == 1 &&
		config.GetConfig().App.AllowOrigin[0] == "*" {
		// allow all origin
		app.Use(cors.Default())
	} else {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = config.GetConfig().App.AllowOrigin
		app.Use(cors.New(corsConfig))
	}

	// enable swagger doc if enabled
	if config.GetConfig().App.EnableSwagger {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// init router
	v1 := app.Group("/api/v1")
	router.InitRouter(v1)

	// start hot update handler
	// config will be reloaded with SYSUSR1 signal
	util.SetupConfigHotUpdate()

	// start server with endless
	// server will reload with HUP signal
	// server will stop with INT signal
	server := endless.NewServer(
		config.GetConfig().App.Addr,
		app,
	)
	server.BeforeBegin = func(add string) {
		log.Info("NEUOJ-NG-backend started")
	}
	server.ListenAndServe()
	log.Info("NEUOJ-NG-backend terminated")
}
