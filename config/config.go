package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Kafka `yaml:"kafka"`
	Redis `yaml:"redis"`
}

type Kafka struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Topic  string `yaml:"topic"`
	Offset int64  `yaml:"offset"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	content, err := os.ReadFile(filepath.Join("config", "config.yaml"))
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal data for config: %w", err)
	}

	return cfg, nil
}
