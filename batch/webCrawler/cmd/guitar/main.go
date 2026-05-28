package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(
        colly.Async(true),
        colly.MaxDepth(3),
    )

    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        Parallelism: 10,
    })

	// depth:0 ギターシリーズ一覧
    c.OnHTML("#item", func(e *colly.HTMLElement) {
		e.ForEach(".item_guitar", func(idx int, elem *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(elem.ChildAttr("a", "href"))
			fmt.Println("シリーズ:", elem.ChildText(".fig p.trim_name"))
			e.Request.Visit(link)
		})
    })

    // depth:1 ギター一覧
    c.OnHTML("div.items div .fig", func(e *colly.HTMLElement) {
		e.ForEach(".wrap", func(idx int, elem *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(elem.ChildAttr("a", "href"))
			// fmt.Println("ギター名:", elem.ChildText(".items_txt, h2.uppercase"))
			e.Request.Visit(link)
		})
    })

    // depth:2 ギター詳細(depth:2はギター一覧である可能性あり)
    c.OnHTML("#main", func(e *colly.HTMLElement) {
		// ギター詳細画面の処理
		name := e.ChildText(".overlay_bgr_header .header_content .header_title")

		if len(name) == 0 { return }

		subName := e.ChildText(".overlay_bgr_header .header_content .clr_name")
		price := e.ChildText("#specifications .detail_price")
		fmt.Printf("guitar詳細 --- ESP [ 名前 ]: %v %v [ 価格 ]: %v\n", name, subName, price)
    })
	// depth:2 ギター一覧 >> 詳細
	c.OnHTML("div.items", func(e *colly.HTMLElement) {
		// ギター一覧画面であった場合、再クロール
		e.ForEach(".fig .figcap", func(idx int, elem *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(elem.ChildAttr("a", "href"))
			e.Request.Visit(link)
		})
	})

	// depth:3 ギター詳細
    c.OnHTML("#main", func(e *colly.HTMLElement) {
		// ギター詳細画面の処理
		name := e.ChildText(".overlay_bgr_header .header_content .header_title")

		if len(name) == 0 { return }

		subName := e.ChildText(".overlay_bgr_header .header_content .clr_name")
		price := e.ChildText("#specifications .detail_price")
		fmt.Printf("guitar詳細 --- ESP [ 名前 ]: %v %v [ 価格 ]: %v\n", name, subName, price)
    })

    c.OnError(func(r *colly.Response, err error) {
		log.Fatal("Error:", err)
    })

    c.Visit("https://espguitars.co.jp/products/esp")
    c.Wait()
}
