package access

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// WebRetriever is the Retriever implementation to fetch the blog
// resources from the web.
type WebRetriever struct{}

// Fetch gets the raw blog feed content.
func (w WebRetriever) Fetch(blogUrl string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	randomCacheNumber := rand.Intn(10000)
	blogUrlModified := fmt.Sprintf(
		"%s?nocache=%d",
		blogUrl,
		randomCacheNumber,
	)
	response, err := http.Get(blogUrlModified)
	fmt.Printf("Blog URL: %s\n", blogUrlModified)
	fmt.Printf("Response content length: %d\n", response.ContentLength)
	fmt.Printf("Response status: %d %s\n", response.StatusCode, response.Status)
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
