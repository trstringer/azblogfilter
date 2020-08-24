package config

import (
	yaml "gopkg.in/yaml.v3"

	"github.com/trstringer/azblogfilter/internal/blog"
)

// Retriever is the contract to get config from source.
type Retriever interface {
	Fetch() (string, error)
}

// Config represents application config.
type Config struct {
	Keywords   []blog.Keyword  `yaml:"keywords"`
	Categories blog.Categories `yaml:"categories"`
}

// GetConfig gets and parses application config.
func GetConfig(retriever Retriever) (*Config, error) {
	rawConfig, err := retriever.Fetch()
	if err != nil {
		return nil, err
	}

	rawConfigBytes := []byte(rawConfig)
	config := Config{}
	yaml.Unmarshal(rawConfigBytes, &config)

	return &config, nil
}
