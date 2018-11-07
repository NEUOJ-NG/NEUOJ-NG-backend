package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"sync"
)

// singleton mode for Config
var (
	cfg     *Config
	cfgLock sync.RWMutex
	once    sync.Once
)

type Config struct {
	App appConfig `toml:"app"`
}

func GetConfig() *Config {
	once.Do(ReloadConfig)
	cfgLock.RLock()
	defer cfgLock.RUnlock()
	return cfg
}

func ReloadConfig() {
	filePath, err := filepath.Abs("./config.toml")
	if err != nil {
		panic(err)
	}
	log.Info("parsing config.toml")
	config := new(Config)
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}
	cfgLock.Lock()
	defer cfgLock.Unlock()
	cfg = config
	return
}
