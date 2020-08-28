package config

import (
	"testing"

	"github.com/trstringer/azblogfilter/internal/blog"
)

type fakeRetriever struct{}

func (f fakeRetriever) Fetch() (string, error) {
	testConfig := `
keywords:
  - keyword1
  - keyword2

categories:
  - category1
  - category2`

	return testConfig, nil
}

func TestGetConfig(t *testing.T) {
	retriever := fakeRetriever{}
	actualConfig, err := GetConfig(retriever)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if actualConfig == nil {
		t.Fatal("Unexpected nil config")
	}

	desiredConfig := Config{
		Keywords:   []blog.Keyword{"keyword1", "keyword2"},
		Categories: []blog.Category{"category1", "category2"},
	}

	if !configsAreEqual(actualConfig, &desiredConfig) {
		t.Fatal("Configs do not match")
	}
}

func configsAreEqual(c1 *Config, c2 *Config) bool {
	if len(c1.Keywords) != len(c2.Keywords) ||
		len(c1.Categories) != len(c2.Categories) {
		return false
	}

	for _, keyword := range c1.Keywords {
		if !keywordsContainKeyword(c2.Keywords, keyword) {
			return false
		}
	}

	for _, category := range c1.Categories {
		if !c2.Categories.ContainsCategory(category) {
			return false
		}
	}

	return true
}

func keywordsContainKeyword(input []blog.Keyword, test blog.Keyword) bool {
	for _, instance := range input {
		if instance == test {
			return true
		}
	}

	return false
}

func TestGetConfigStrings(t *testing.T) {
	keywords := "keyword1,keyword2,keyword3"
	categories := "category1,category2"
	desiredConfig := Config{
		Keywords:   []blog.Keyword{"keyword1", "keyword2", "keyword3"},
		Categories: blog.Categories{"category1", "category2"},
	}
	actualConfig := GetConfigFromCLI(keywords, categories)

	if !configsAreEqual(&actualConfig, &desiredConfig) {
		t.Fatal("Configs from CLI string comparison are unexpectedly different")
	}
}
