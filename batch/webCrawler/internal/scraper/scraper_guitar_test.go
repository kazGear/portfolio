package scraper

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/stretchr/testify/assert"
)

func TestCollectURLsESP(t *testing.T) {
    _, cancel := chromedp.NewContext(context.Background())

	guitar := NewEspScraper(cancel)
	urls   := guitar.CollectLinks()
	fmt.Printf("urlsCount: %v", len(urls))

	assert.GreaterOrEqual(t, len(urls), 350)
}

func TestConvertLabel(t *testing.T) {
	items := []struct{
		label  string
		want   string
	}{
		{
			label: "BODY", want: "BodyMaterial",
		},{
			label: "NECK", want: "NeckMaterial",
		},{
			label: "PARTS COLOR", want: "Comment",
		},{
			label: "xxx", want: "",
		},
	}

	for _, item := range items {
		actual := convertLabel(item.label)
		assert.Equal(t, item.want, actual)
	}
}

func TestIsFirstVisit(t *testing.T) {
	urls := []struct{
		url  string
		want bool
	}{
		{
			url: "https://example.com/1", want: true,
		}, {
			url: "https://example.com/1", want: false,
		}, {
			url: "https://example.com/2", want: true,
		}, {
			url: "https://example.com/3", want: true,
		}, {
			url: "https://example.com/3", want: false,
		}, {
			url: "https://example.com/1", want: false,
		},
	}

	var mutex   = &sync.Mutex{}
	visited    := map[string]struct{}{}

	for _, url := range urls {
		url    := url
		actual := isFirstVisit(mutex, url.url, visited)
		assert.True(t, url.want == actual)
	}
}