package config

import (
	"strings"

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

// FilterBlogPosts filters all blog posts by config.
func (c Config) FilterBlogPosts(posts []blog.Post) []blog.Post {
	filteredPosts := []blog.Post{}
	for _, post := range posts {
		postAdded := false
		for _, keyword := range c.Keywords {
			if post.Title.HasKeyword(keyword) {
				filteredPosts = append(filteredPosts, post)
				postAdded = true
				break
			}
		}
		if postAdded {
			continue
		}
		for _, postCategory := range post.Categories {
			if c.Categories.ContainsCategory(postCategory) {
				filteredPosts = append(filteredPosts, post)
				break
			}
		}
	}
	return filteredPosts
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

// GetConfigFromCLI parses config passed from CLI.
func GetConfigFromCLI(keywords, categories string) Config {
	var configKeywords []blog.Keyword
	for _, keyword := range strings.Split(keywords, ",") {
		configKeywords = append(configKeywords, blog.Keyword(keyword))
	}
	var configCategories blog.Categories
	for _, category := range strings.Split(categories, ",") {
		configCategories = append(configCategories, blog.Category(category))
	}

	return Config{
		Keywords:   configKeywords,
		Categories: configCategories,
	}
}
