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
	Server ConfigServer `yaml:"server"`
}

type ConfigServer struct {
	Listen string `yaml:"listen"`
}
