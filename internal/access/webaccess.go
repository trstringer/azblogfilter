package access

import (
	"io/ioutil"
	"net/http"
)

// WebRetriever is the Retriever implementation to fetch the blog
// resources from the web.
type WebRetriever struct{}

// Fetch gets the raw blog feed content.
func (w WebRetriever) Fetch(blogUrl string) (string, error) {
	response, err := http.Get(blogUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	responseContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseContent), nil
}

// NewWebRetriever gets and instance of a WebRetriever.
func NewWebRetriever() WebRetriever {
	return WebRetriever{}
}
