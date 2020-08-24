package access

import (
	"encoding/xml"

	"github.com/trstringer/azblogfilter/internal/blog"
)

// Rss defines the rss XML tag.
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	Version string   `xml:"version,attr"`
}

// Channel defines the channel XML tag.
type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Item   `xml:"item"`
}

// Item defines the item XML tag.
type Item struct {
	XMLName         xml.Name        `xml:"item"`
	ID              string          `xml:"guid"`
	Categories      []blog.Category `xml:"category"`
	Title           blog.Title      `xml:"title"`
	Description     string          `xml:"description"`
	PublicationDate string          `xml:"pubDate"`
	Link            string          `xml:"link"`
}
