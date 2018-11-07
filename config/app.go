package config

type appConfig struct {
	Addr          string   `toml:"addr"`
	AllowOrigin   []string `toml:"allow_origin"`
	EnableSwagger bool     `toml:"enable_swagger"`
	LogFile       string   `toml:"log_file"`
}
