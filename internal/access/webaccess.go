package access

import (
	"io/ioutil"
	"net/http"
)

// WebRetriever is the Retriever implementation to fetch the blog
// resources from the web.
type WebRetriever struct{}

// Fetch gets the raw blog feed content.
func (w WebRetriever) Fetch() (string, error) {
	response, err := http.Get("https://azurecomcdn.azureedge.net/en-us/updates/feed/")
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
