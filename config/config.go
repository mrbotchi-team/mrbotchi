package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Host      string         `toml:"host"`
		Port      int            `toml:"port"`
		PasetoKey string         `toml:"paseto_key"`
		DB        DatabaseConfig `toml:"database"`
		User      UserConfig     `toml:"user"`
		Argon2    Argon2Config   `toml:"argon2"`
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
	Argon2Config struct {
		Memory      uint32 `toml:"memory_cost"`
		Iterations  uint32 `toml:"iteration_cost"`
		Parallelism uint8  `toml:"parallelism"`
		SaltLength  uint32 `toml:"salt_length"`
		KeyLength   uint32 `toml:"key_length"`
	}
)

func LoadConfig() *Config {
	var config Config

	if _, err := toml.DecodeFile("/.config/mrbotchi.toml", &config); nil != err {
		log.Fatalln("File not found!")
	}

	return &config
}
