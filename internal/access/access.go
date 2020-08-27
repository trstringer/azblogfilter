package access

import (
	"encoding/xml"
	"time"

	"github.com/trstringer/azblogfilter/internal/blog"
)

// DateTimeLayout is the standard for layout of Time.
const DateTimeLayout string = "Mon, 02 Jan 2006 15:04:05 Z"

// Retriever is the interface for retrieving blog data.
type Retriever interface {
	Fetch() (string, error)
}

// GetPosts gets all blog post data.
func GetPosts(retriever Retriever) ([]blog.Post, error) {
	rawData, err := retriever.Fetch()
	if err != nil {
		return nil, err
	}

	rss := parsePostsXML(rawData)

	blogPosts, err := extractPostsFromRss(rss)
	if err != nil {
		return nil, err
	}

	return blogPosts, nil
}

func parsePostsXML(rawXML string) *Rss {
	rss := Rss{}
	rawXMLBytes := []byte(rawXML)
	xml.Unmarshal(rawXMLBytes, &rss)
	return &rss
}

func extractPostsFromRss(rss *Rss) ([]blog.Post, error) {
	blogPosts := []blog.Post{}
	for _, blogPostRaw := range rss.Channel.Items {
		blogPostPublishedTime, err := time.Parse(DateTimeLayout, blogPostRaw.PublicationDate)
		if err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, blog.Post{
			Title:       blogPostRaw.Title,
			Categories:  blogPostRaw.Categories,
			Description: blogPostRaw.Description,
			ID:          blogPostRaw.ID,
			Link:        blogPostRaw.Link,
			PublishedAt: blogPostPublishedTime,
		})
	}

	return blogPosts, nil
}
