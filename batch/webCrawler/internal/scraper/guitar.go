package scraper

import (
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

func ScrapeGuitars() ([]model.Guitar, error) {
    guitars := []model.Guitar{}

    c := colly.NewCollector()

    c.OnHTML(".item", func(e *colly.HTMLElement) {
        g := model.Guitar{
            Model: e.ChildText(".model"),
            Maker: e.ChildText(".maker"),
            Price: utils.ParsePrice(e.ChildText(".price")),
            URL:   e.ChildAttr("a", "href"),
            Image: e.ChildAttr("img", "src"),
            Shop:  "Ishibashi",
        }
        guitars = append(guitars, g)
    })

    err := c.Visit("https://kazapp-trial.com/")
    return guitars, err
}
