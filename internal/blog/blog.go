package blog

import (
	"strings"
	"time"
)

// Post is the base struct for holding data for a single blog post.
type Post struct {
	ID          string
	Description string
	PublishedAt time.Time
	Link        string
	Title
	Categories
}

// Category represents a blog post category. Categories can be used
// for blog post filtering, alongside Keywords.
type Category string

// Categories is a collection of multiple Category instances.
type Categories []Category

// Keyword represents a blog post keyword that can be used to search
// through the title.
type Keyword string

// Title represents the blog post title that can be used to search
// for a keyword.
type Title string

// String returns the string representation of the category.
func (c Category) String() string {
	return string(c)
}

// Equals checks if two Category instances are the same.
func (c Category) Equals(c2 Category) bool {
	return strings.EqualFold(c.String(), c2.String())
}

// ContainsCategory checks if the post contains a category.
func (c Categories) ContainsCategory(category Category) bool {
	for _, postCategory := range c {
		if postCategory.Equals(category) {
			return true
		}
	}
	return false
}

// HasKeyword checks if the Title has a keyword match.
func (t Title) HasKeyword(keyword Keyword) bool {
	lowerTitle := strings.ToLower(string(t))
	lowerKeyword := strings.ToLower(string(keyword))
	return strings.Contains(lowerTitle, lowerKeyword)
}
