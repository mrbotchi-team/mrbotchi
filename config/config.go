package config

import (
	"errors"
	"log"
	"net/url"

	"github.com/BurntSushi/toml"
)

type (
	urlItem struct{ *url.URL }
	// Config は設定を表現する構造体。
	Config struct {
		Host urlItem `toml:"host"`
		Port int     `toml:"port"`

		Actor struct {
			PreferredUsername string `toml:"preferredUsername"`
			Name              string `toml:"name"`
			Summary           string `toml:"summary"`
		} `toml:"actor"`
	}
)

// LoadConfig は設定ファイルを読み込む関数
func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); nil != err {
		log.Println("error: Failed to decode config file!")
		return nil, err
	}

	return &config, nil
}

func (u *urlItem) UnmarshalText(text []byte) error {
	var err error
	u.URL, err = url.Parse(string(text))
	if nil != err || "" == u.URL.Scheme || "" == u.URL.Host {
		return errors.New("host config isnt correct. Please check the config file")
	}

	return nil
}
