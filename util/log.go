package util

import (
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func SetupLog() {
	level, err := log.ParseLevel(config.GetConfig().App.LogLevel)
	if err != nil {
		log.Fatal("parse log_level failed")
		log.Fatal("please check config.toml")
		panic(err)
	}
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.Info("setting up log file")
	CreateDirOrPanic(filepath.Dir(config.GetConfig().App.LogFile))
	logFile, err := os.OpenFile(
		config.GetConfig().App.LogFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)
}
