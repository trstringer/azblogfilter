package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/trstringer/go-systemd-time/systemdtime"

	"github.com/trstringer/azblogfilter/internal/access"
	"github.com/trstringer/azblogfilter/internal/blog"
	"github.com/trstringer/azblogfilter/internal/cache"
	"github.com/trstringer/azblogfilter/internal/config"
)

func getBlogPosts(since time.Time, keywords, categories string, blogUrls []string) ([]blog.Post, error) {
	blogPosts := []blog.Post{}
	for _, blogUrl := range blogUrls {
		blogPostsFromUrl, err := access.GetPostsFromWeb(blogUrl)
		if err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, blogPostsFromUrl...)
	}

	blogPostsFiltered := filterBlogPostsByTime(blogPosts, since)
	config := config.GetConfigFromCLI(keywords, categories)
	blogPostsFiltered = config.FilterBlogPosts(blogPostsFiltered)

	return blogPostsFiltered, nil
}

func filterBlogPostsByTime(blogPosts []blog.Post, filter time.Time) []blog.Post {
	blogPostsFiltered := []blog.Post{}
	for _, blogPost := range blogPosts {
		if blogPost.PublishedAt.After(filter) {
			blogPostsFiltered = append(blogPostsFiltered, blogPost)
		}
	}

	return blogPostsFiltered
}

func effectiveSinceTime() (time.Time, error) {
	// --since takes precendence over --cache for determining the effective
	// since time for filtering blog posts.
	if sinceFilter != "" {
		now := time.Now()
		since, err := systemdtime.AdjustTime(&now, sinceFilter)
		if err != nil {
			return time.Time{}, err
		}
		return since, nil
	}

	cachePath, err := realCachePath()
	if err != nil {
		return time.Time{}, err
	}
	lastCachedTime, err := cache.LastCachedTimeFromFileSystem(cachePath)
	if err != nil {
		errorMsg := err.Error()
		if strings.HasSuffix(errorMsg, "no such file or directory") {
			err = nil
		}
		return time.Time{}, err
	}
	return lastCachedTime, nil
}

func realCachePath() (string, error) {
	cachePath, err := homedir.Expand(cacheLocation)
	if err != nil {
		return "", err
	}
	return cachePath, nil
}

func formatBlogPosts(format string, posts []blog.Post) (string, error) {
	var formatFunc func([]blog.Post) (string, error)
	switch format {
	case "json":
		formatFunc = formatBlogPostsJSON
	case "csv":
		formatFunc = formatBlogPostsCSV
	}

	return formatFunc(posts)
}

func formatBlogPostsJSON(posts []blog.Post) (string, error) {
	output, err := json.Marshal(posts)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func formatBlogPostsCSV(posts []blog.Post) (string, error) {
	var output strings.Builder
	for _, post := range posts {
		output.WriteString(fmt.Sprintf(
			"%s,%s,%s\n",
			post.PublishedAt.String(),
			post.Title,
			post.Link,
		))
	}
	return output.String(), nil
}

func validateArgs() error {
	if !useCache && sinceFilter == "" {
		return fmt.Errorf("You must specify either --cache or --since")
	}

	if !contains(outputOptions, outputFormat) {
		return fmt.Errorf(
			"Unknown output '%s'. Possible: %v",
			outputFormat,
			outputOptions,
		)
	}

	return nil
}

func contains(collection []string, search string) bool {
	for _, item := range collection {
		if item == search {
			return true
		}
	}
	return false
}
