package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func NewConfigFromFile() (*Config, error) {
	b, err := os.ReadFile(args.ConfigFilePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

type Config struct {
	Server  ConfigServer  `yaml:"server"`
	Session ConfigSession `yaml:"session"`
	Cookie  ConfigCookie  `yaml:"cookie"`
	Redis   ConfigRedis   `yaml:"redis"`
}

type ConfigServer struct {
	Listen string `yaml:"listen"`
}

type ConfigSession struct {
	Secret string             `yaml:"secret"`
	Store  ConfigSessionStore `yaml:"store"`
}
type ConfigSessionStore struct {
	Prefix string `yaml:"prefix"`
}

type ConfigCookie struct {
	Path   string `yaml:"path"`
	Domain string `yaml:"domain"`
	MaxAge int    `yaml:"max_age"`
	Secure bool   `yaml:"secure"`
}

type ConfigRedis struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	MaxIdle  int    `yaml:"max_idle"`
	DB       int    `yaml:"db"`
}
