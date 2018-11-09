package config

import "time"

type jwtConfig struct {
	Realm           string        `toml:"realm"`
	Key             string        `toml:"key"`
	Timeout         time.Duration `toml:"timeout"`
	MaxRefreshDelay time.Duration `toml:"max_refresh_delay"`
}
