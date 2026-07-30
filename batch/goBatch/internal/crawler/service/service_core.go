package service

import (
	"context"

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
    // 動的ページ取得のためのchromedpコンテキスト構築
    allocCtx, allocCancel := chromedp.NewExecAllocator(
        context.Background(),
        chromedp.DefaultExecAllocatorOptions[:]...,
    )
    parentCtx, parentCancel := chromedp.NewContext(allocCtx)

    return allocCancel, parentCancel, parentCtx
}

// chromedp環境構築（実際にchromeを起動して挙動を確認できる。clickされているか？等）
func createChromedpCtxDebug() (cancelAlloc context.CancelFunc,
                               cancelParent context.CancelFunc,
                               parentCtx context.Context,
) {
    // 動的ページ取得のためのchromedpコンテキスト構築
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", false),
    )

    allocCtx, allocCancel := chromedp.NewExecAllocator(
        context.Background(),
        opts...
    )
    parentCtx, parentCancel := chromedp.NewContext(allocCtx)

    return allocCancel, parentCancel, parentCtx
}