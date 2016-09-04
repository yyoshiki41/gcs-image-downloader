package downloader

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var (
	tomlCredentials = "credentials.toml"
)

type Config struct {
	API APICredentials
}

type APICredentials struct {
	Cx  string `toml:"cx"`
	Key string `toml:"key"`
}

func loadConf(path string, conf *Config) error {
	credentials := filepath.Join(path, tomlCredentials)

	_, err := toml.DecodeFile(credentials, &conf)
	return err
}
