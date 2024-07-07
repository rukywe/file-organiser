package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SourceDir     string            `yaml:"source_dir"`
	LogLevel      string            `yaml:"log_level"`
	FileExtensions map[string]string `yaml:"file_extensions"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
