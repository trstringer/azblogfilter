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
	httpClient := &http.Client{}
	request, err := http.NewRequest("GET", blogUrlModified, nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0")
	response, err := httpClient.Do(request)
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
