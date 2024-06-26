package config

import (
	"fmt"
	"xxl-job/tools"

	"gopkg.in/ini.v1"
)

type Config struct {
	Ceye `ini:"ceye"`
}

type Ceye struct {
	Identifier string `ini:"identifier"`
	Token      string `ini:"token"`
	Types      string `ini:"type"`
}

func ReturnConfig() (*Config, error) {
	path, err := tools.RootPath()

	if err != nil {
		return &Config{}, err
	}

	conf_ini := fmt.Sprintf("%s\\config\\config.ini", path)
	conf, err := ini.Load(conf_ini)
	if err != nil {
		return &Config{}, err
	}

	cfg := new(Config)
	conf.MapTo(cfg)

	// ini.MapTo(cfg, conf_ini)
	return cfg, nil
}
