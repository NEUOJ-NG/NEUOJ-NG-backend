package util

import (
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/config"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func SetupConfigHotUpdate() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			config.ReloadConfig()
			log.Info("config reloaded")
		}
	}()
}
