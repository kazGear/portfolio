package scraper

import (
	"log"
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type Scraper interface {
	Scrape() ([]model.Guitar, error)
}

// htmlページのリンクを次々に辿っていく
func navigateLinks(cssSelector string) func(*colly.HTMLElement) {
    return func(html *colly.HTMLElement) {
        html.ForEach(cssSelector, func(idx int, elem *colly.HTMLElement) {
            link := elem.Request.AbsoluteURL(elem.ChildAttr("a", "href"))
            elem.Request.Visit(link)
        })
    }
}

// ギター構造体の構築
func buildGuitar(spec map[string]string) (*model.Guitar) {

	guitar := model.Guitar{}

    var errMaker error
	guitar.Maker, errMaker = strconv.Atoi(spec["Maker"])
	guitar.Name            = spec["Name"]

    if guitar.Maker <= 0 || len(guitar.Name) == 0 || errMaker != nil {
		log.Printf("メーカー,名称は必須項目です。maker: %v, name: %v\n", guitar.Maker, guitar.Name)
		return &model.Guitar{}
	}

	guitar.BodyFinish        = spec["BodyFinish"]
	guitar.BodyMaterial      = spec["BodyMaterial"]
    guitar.BodyMaterialBack  = utils.SearchWoodCode(spec["BodyMaterialBack"])
	guitar.BodyMaterialFront = utils.SearchWoodCode(spec["BodyMaterialFront"])
    guitar.Bridge            = spec["Bridge"]
	guitar.Color             = spec["Color"]
	guitar.Comment           = spec["Comment"]
	guitar.Controls          = spec["Controls"]
    guitar.Fingerboard       = utils.SearchWoodCode(spec["Fingerboard"])

	var errFretCount error
	guitar.FretCount, errFretCount = utils.GetFretCount(spec["FretCount"])

    if errFretCount != nil {
        log.Println(errFretCount)
    }

	guitar.Inlays       = spec["Inlays"]
	guitar.Joint        = spec["Joint"]
    guitar.NeckMaterial = utils.SearchWoodCode(spec["NeckMaterial"])
    guitar.Pickups      = spec["Pickups"]

    var errPrice error
	guitar.Price, errPrice = utils.ParsePrice(spec["Price"])

    if errPrice != nil {
        log.Println(errPrice)
    }

    var errScaleLengthMM error
    guitar.ScaleLengthMM, errScaleLengthMM = utils.TrimScaleUnit(spec["ScaleLengthMM"])

    if errScaleLengthMM != nil {
        log.Println(errScaleLengthMM)
    }

	guitar.Series = spec["Series"]
	guitar.Src    = spec["Src"]

    weight, errWeight := strconv.Atoi(spec["Weight"])

    if errWeight != nil {
        log.Println(errWeight)
    }
	guitar.Weight = float64(weight)

	return &guitar
}
