package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Host string         `toml:"host"`
		Port int            `toml:"port"`
		DB   DatabaseConfig `toml:"database"`
		User UserConfig     `toml:"user"`
	}
	DatabaseConfig struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		DBname   string `toml:"dbname"`
	}
	UserConfig struct {
		Name        string `toml:"name"`
		DisplayName string `toml:"display_name"`
		PublicKey   string `toml:"public_key"`
		PrivateKey  string `toml:"private_key"`
	}
)

func LoadConfig() *Config {
	var config Config

	if _, err := toml.DecodeFile(".config/mrbotchi.toml", &config); nil != err {
		log.Fatalln("File not found!")
	}

	return &config
}
