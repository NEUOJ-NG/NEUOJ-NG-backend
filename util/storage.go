package util

import (
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/config"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func CreateDirOrPanic(path string) {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("failed to parse path %v: %v", path, err)
		panic(err)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatalf("failed to create path %v: %v", path, err)
			panic(err)
		}
		log.Infof("path %v created successfully", path)
	}
}

func SetupStorage() {
	CreateDirOrPanic(config.GetConfig().App.StoragePath)
	CreateDirOrPanic(config.GetAppStoragePath())
	CreateDirOrPanic(config.GetAvatarStoragePath())
}
