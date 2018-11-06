package main

import (
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/config"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/router"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/util"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// create Gin Engine with Logger and Recovery middleware
	app := gin.Default()
	router.InitRouter(app)

	// start hot update handler
	// config will be reloaded with SYSUSR1 signal
	util.SetupConfigHotUpdate()

	// start server with endless
	// server will reload with HUP signal
	// server will stop with INT signal
	err := endless.ListenAndServe(
		config.GetConfig().App.Addr,
		app,
	)
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
