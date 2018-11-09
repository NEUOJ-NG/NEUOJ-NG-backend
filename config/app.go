package config

import "path/filepath"

const (
	AppDir    = "app"
	AvatarDir = "avatar"
)

type appConfig struct {
	Addr          string   `toml:"addr"`
	AllowOrigin   []string `toml:"allow_origin"`
	EnableSwagger bool     `toml:"enable_swagger"`
	LogFile       string   `toml:"log_file"`
	LogLevel      string   `toml:"log_level"`
	StoragePath   string   `toml:"storage_path"`
}

func GetAppStoragePath() string {
	return filepath.Join(GetConfig().App.StoragePath, AppDir)
}

func GetAvatarStoragePath() string {
	return filepath.Join(GetConfig().App.StoragePath, AppDir, AvatarDir)
}
