package config

import (
	"candy/pkg/loger"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Address string `toml:"address"`
	Port    uint16 `toml:"port"`
}

func Get() Config {
	var cfg Config

	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		loger.New().Error().Err(err)
	}

	return cfg
}
