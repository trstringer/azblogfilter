package cmd

import (
	"testing"
	"time"

	"github.com/trstringer/azblogfilter/internal/blog"
)

func TestFilterBlogPostsByTime(t *testing.T) {
	blogPosts := []blog.Post{
		{
			Title:       "Title1",
			PublishedAt: time.Now().AddDate(0, 0, -7),
		},
		{
			Title:       "Title2",
			PublishedAt: time.Now().AddDate(0, 0, -3),
		},
		{
			Title:       "Title3",
			PublishedAt: time.Now().AddDate(0, 0, -9),
		},
	}
	sinceFilter := time.Now().AddDate(0, 0, -8)

	filteredBlogPosts := filterBlogPostsByTime(blogPosts, sinceFilter)

	expectedResultLength := 2
	actualResultLength := len(filteredBlogPosts)
	if actualResultLength != expectedResultLength {
		t.Fatalf("Expected %d posts, received %d", expectedResultLength, actualResultLength)
	}
}
