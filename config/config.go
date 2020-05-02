package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config は設定を表現する構造体。
type Config struct {
	Host string `toml:"host"`

	Actor struct {
		PreferredUsername string `toml:"preferredUsername"`
		Name              string `toml:"name"`
		Summary           string `toml:"summary"`
	} `toml:"actor"`
}

// LoadConfig は設定ファイルを読み込む関数
func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); nil != err {
		log.Println("error: Failed to decode config file!")
		return nil, err
	}

	return &config, nil
}
