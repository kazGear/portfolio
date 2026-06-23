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
			label: "BODY", want: "BodyMaterialBack",
		},{
			label: "NECK", want: "NeckMaterial",
		},{
			label: "xxx", want: "",
		},
	}

	for _, item := range items {
		actual, _ := utils.ConvertLabel(item.label, specFieldMap)
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

	var mutex = &sync.Mutex{}
	visited  := map[string]struct{}{}
	guitar   := guitarScraper{ logger: utils.NewLogger("")}

	for _, url := range urls {
		url    := url
		actual := guitar.isFirstVisit(mutex, url.url, visited)
		assert.True(t, url.want == actual)
	}
}

func TestSearchWoodCode(t *testing.T) {
	var maple int = 6
	var hardMaple int = 1

	woods := []string{
		"hardMaple Paduak 7P",
		"HardMaple, Walnut, Paduak 7P",
		"Maple, Paduak 7P HardMaple", // 具体名が拾われる（マスタ順サーチ）
		"HardMaple Maple",
		"Hard Maple Maple",
	}

	assert.Equal(t, hardMaple, searchWoodCode(woods[0]))
	assert.Equal(t, hardMaple, searchWoodCode(woods[1]))
	assert.Equal(t, hardMaple, searchWoodCode(woods[2]))
	assert.Equal(t, hardMaple, searchWoodCode(woods[3]))
	assert.Equal(t, maple, searchWoodCode(woods[4]))
}