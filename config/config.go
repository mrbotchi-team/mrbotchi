package config

import (
	"crypto/rsa"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/mrbotchi-team/mrbotchi/utils"
)

type (
	Config struct {
		Host      string             `toml:"host"`
		Port      int                `toml:"port"`
		PasetoKey string             `toml:"paseto_key"`
		DB        *DatabaseConfig    `toml:"database"`
		Account   *AccountConfig     `toml:"account"`
		Argon2    *utils.Argon2Param `toml:"argon2"`
	}
	DatabaseConfig struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		DBname   string `toml:"dbname"`
	}
	AccountConfig struct {
		Name        string          `toml:"name"`
		DisplayName string          `toml:"display_name"`
		Summary     string          `toml:"summary"`
		PublicKey   *rsa.PublicKey  `toml:"-"`
		PrivateKey  *rsa.PrivateKey `toml:"-"`
	}
)

func LoadConfig() *Config {
	var config Config

	if _, err := toml.DecodeFile("/.config/mrbotchi.toml", &config); nil != err {
		log.Fatalln("File not found!")
	}

	privateKey, err := utils.ReadRSAPrivateKey("/.config/private.pem")
	if nil != err {
		log.Fatalln(err)
	}
	config.Account.PrivateKey = privateKey

	publicKey, err := utils.ReadRSAPublicKey("/.config/public.pem")
	if nil != err {
		log.Fatalln(err)
	}
	config.Account.PublicKey = publicKey

	return &config
}
