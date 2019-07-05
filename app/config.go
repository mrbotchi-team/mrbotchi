package app

import (
	"log"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Host string         `toml:"host"`
		Port int            `toml:"port"`
		DB   DatabaseConfig `toml:"database"`
	}
	DatabaseConfig struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		DBname   string `toml:"dbname"`
	}
)

func loadConfig() *Config {
	var config Config

	if _, err := toml.DecodeFile("config.toml", &config); nil != err {
		log.Fatalln(err)
	}

	return &config
}
