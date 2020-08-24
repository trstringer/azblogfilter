package access

import (
	"testing"
	"time"
)

type fakeRetriever struct{}

func (t fakeRetriever) Fetch() (string, error) {
	rssData := `
	<rss version='2.0'>
		<channel>
			<item>
				<guid>post-id-1</guid>
				<title>title1</title>
				<description>description1</description>
				<link>link1</link>
				<category>category1</category>
				<category>category2</category>
				<pubDate>Mon, 17 Aug 2020 09:00:03 Z</pubDate>
			</item>
			<item>
				<guid>post-id-2</guid>
				<title>title2</title>
				<description>description2</description>
				<link>link1</link>
				<category>category1</category>
				<category>category3</category>
				<category>category4</category>
				<pubDate>Thu, 06 Aug 2020 09:00:11 Z</pubDate>
			</item>
		</channel>
	</rss>`
	return rssData, nil
}

func TestGetPosts(t *testing.T) {
	dataRetriever := fakeRetriever{}
	blogPosts, err := GetPosts(dataRetriever)
	if err != nil {
		t.Errorf("Expected no error: %v", err)
	}

	desiredItemCount := 2
	actualItemCount := len(blogPosts)
	if actualItemCount != desiredItemCount {
		t.Errorf("Expected %d items, got %d", desiredItemCount, actualItemCount)
	}
}

func TestDateTimeStringToTime(t *testing.T) {
	input := "Thu, 06 Aug 2020 09:00:11 Z"
	desiredOutput := time.Date(2020, time.August, 6, 9, 0, 11, 0, time.UTC)
	actualOutput, err := datetimeStringToTime(input)
	if err != nil {
		t.Errorf("Error converting time: %v", err)
	}
	if !actualOutput.Equal(desiredOutput) {
		t.Errorf("Expected %v, got %v", desiredOutput, actualOutput)
	}
}
