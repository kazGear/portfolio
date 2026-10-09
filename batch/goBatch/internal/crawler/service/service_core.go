package service

import (
	"context"
	"runtime"

	"github.com/chromedp/chromedp"
)

type CrawlerService interface {
	RunCrawler()
}

// chromedp環境構築
func createChromedpCtx() (cancelAlloc context.CancelFunc,
                          cancelParent context.CancelFunc,
                          parentCtx context.Context,
) {
    opts := append(
        []chromedp.ExecAllocatorOption{},
        chromedp.DefaultExecAllocatorOptions[:]...,
    )
    // Linux（Docker）ではChromiumのパスを指定
    if runtime.GOOS == "linux" {
        opts = append(opts, chromedp.ExecPath("/usr/bin/chromium"))
    }

    // 動的ページ取得のためのchromedpコンテキスト構築
    allocCtx, allocCancel := chromedp.NewExecAllocator(
        context.Background(),
        opts...,
    )
    parentCtx, parentCancel := chromedp.NewContext(allocCtx)

    return allocCancel, parentCancel, parentCtx
}