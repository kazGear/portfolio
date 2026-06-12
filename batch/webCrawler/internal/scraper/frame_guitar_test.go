package scraper

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var (
	ctx, _cancel = chromedp.NewContext(context.Background())
)

func TestCollectURLsEsp(t *testing.T) {
	guitar  := NewScraperEsp(nil)
	urls, _ := guitar.CollectLinks(ctx)
	fmt.Printf("urlsCount: %v", len(urls))

	assert.GreaterOrEqual(t, len(urls), 350)
}

func TestCollectURLsEspSig(t *testing.T) {
	guitar  := NewScraperEspSig(nil)
	urls, _ := guitar.CollectLinks(ctx)
	fmt.Printf("urlsCount: %v", len(urls))

	assert.LessOrEqual(t, len(urls), 100)
}

func TestCollectURLsStrandberg(t *testing.T) {
	guitar  := NewScraperStrandberg(nil)
	urls, _ := guitar.CollectLinks(ctx)
	fmt.Printf("urlsCount: %v", len(urls))

	assert.GreaterOrEqual(t, len(urls), 30)
}

func TestConvertLabelEsp(t *testing.T) {
	items := []struct{
		label  string
		want   string
	}{
		{
			label: "BODY", want: "BodyMaterial",
		},{
			label: "NECK", want: "NeckMaterial",
		},{
			label: "xxx", want: "",
		},
	}

	for _, item := range items {
		actual := utils.ConvertLabel(item.label, fieldMapEsp)
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