package scraper

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

// ギター構造体の構築フレームワーク
func buildGuitarFrame(spec map[string]string, url string) (*model.Guitar) {
	guitar := model.Guitar{}
    trim   := utils.TrimSpace()

    var errMaker error
	guitar.Maker, errMaker = strconv.Atoi(spec[C.Maker])
	guitar.Name            = trim(spec[C.Name])

    if errMaker != nil {
        log.Printf("[Maker convert error]: %v", errMaker)
        return &model.Guitar{}
	}
	guitar.BodyFinish = trim(spec[C.BodyFinish])

	// ボディ材構築
    if len(spec[C.BodyMaterialBack]) <= 0 { // バック材がない
        guitar.BodyMaterial = trim(spec[C.BodyMaterialTop])
    } else if len(spec[C.BodyMaterialTop]) <= 0 { // トップ材が無い
        guitar.BodyMaterial = trim(spec[C.BodyMaterialBack])
    } else {
        guitar.BodyMaterial = trim(spec[C.BodyMaterialTop]) + " / " + trim(spec[C.BodyMaterialBack])
    }

    guitar.BodyMaterialBack  = searchWoodCode(spec[C.BodyMaterialBack])
	guitar.BodyMaterialTop   = searchWoodCode(spec[C.BodyMaterialTop])
    guitar.Bridge            = trim(spec[C.Bridge])
	guitar.Color             = trim(spec[C.Color])
    guitar.ColorCd           = utils.ConvertColorCd(guitar.Color)
	guitar.Comment           = trim(spec[C.Comment])
	guitar.Controls          = trim(spec[C.Controls])
    guitar.Fingerboard       = searchWoodCode(spec[C.Fingerboard])

	var errFretCount error
	fretCount                     := trim(spec[C.FretCount])
    guitar.FretCount, errFretCount = utils.GetFretCount(fretCount)
    if errFretCount != nil {
        // log.Println(errFretCount)
    }
	guitar.Inlays       = trim(spec[C.Inlays])
	guitar.Joint        = trim(spec[C.Joint])
    guitar.NeckMaterial = searchWoodCode(spec[C.NeckMaterial])

    // ピックアップ構築
    if len(spec[C.CenterPickup]) > 0 { // センターピックアップあり
        guitar.Pickups = trim(spec[C.NeckPickup])   + " / " +
                         trim(spec[C.CenterPickup]) + " / " +
                         trim(spec[C.BridgePickup])
    } else if len(spec[C.NeckPickup]) <= 0 && len(spec[C.CenterPickup]) <= 0 { // ブリッジピックアップのみ
        guitar.Pickups = trim(spec[C.BridgePickup])
    } else {
        guitar.Pickups = trim(spec[C.NeckPickup]) + " / " + trim(spec[C.BridgePickup])
    }

	guitar.Price = utils.ParsePrice(spec[C.Price])

    scaleLengthMM       := trim(spec[C.ScaleLengthMM])
    guitar.ScaleLengthMM = int(utils.ParseScale(scaleLengthMM))
	guitar.Series        = trim(spec[C.Series])

    guitar.Src           = trim(spec[C.Src])

	// 画像の相対パスをフルパスへ
    if strings.HasPrefix(guitar.Src, "/") {
        fullPass, err := utils.CreateImagePath(url, guitar.Src)

        if err != nil {
            log.Println(err)
        }
        guitar.Src = fullPass
    }

    guitar.Weight = utils.ParseWeight(trim(spec[C.Weight]))

	return &guitar
}

var specFieldMap = map[string]string{
	"Top":                     C.BodyMaterialTop,
	"Top Wood":                C.BodyMaterialTop,
    "Body Top":                C.BodyMaterialTop,

    "Back Wood":               C.BodyMaterialBack,
	"Body":                    C.BodyMaterialBack,
	"BODY":                    C.BodyMaterialBack,
	"Body Material":           C.BodyMaterialBack,
	"Body Wood":               C.BodyMaterialBack,
    "Body Back":               C.BodyMaterialBack,
    "Back & Sides":            C.BodyMaterialBack,

	"Finish":                  C.BodyFinish,
	"Finish Type":             C.BodyFinish,
    "Body Finish":             C.BodyFinish,

	"Bridge":                  C.Bridge,
	"BRIDGE":                  C.Bridge,

	"COLOR":                   C.Color,
    "Color":                   C.Color,
    "Body Color":              C.Color,

	"Controls":                C.Controls,
	"CONTROLS":                C.Controls,
    "CONTROL":                 C.Controls,

	"Fingerboard Material":    C.Fingerboard,
	"FINGERBOARD":             C.Fingerboard,
    "FINGER BOARD":            C.Fingerboard,
	"Fretboard Wood":          C.Fingerboard,
    "Fingerboard":             C.Fingerboard,
    "Fingerboard & Bridge":    C.Fingerboard,

	"FRET":                    C.FretCount,
	"FRETS":                   C.FretCount,
    "Frets":                   C.FretCount,
	"Number Of Frets":         C.FretCount,
	"Number of Frets":         C.FretCount,

	"INLAY":                   C.Inlays,
	"Inlays":                  C.Inlays,
	"Fretboard Inlay":         C.Inlays,
    "Position Inlays":         C.Inlays,
    "Fret Markers":            C.Inlays,

	"CONSTRUCTION":            C.Joint,
	"Neck Joint":              C.Joint,
    "Joint":                   C.Joint,
    "JOINT":                   C.Joint,
	"Neck/Body Assembly Type": C.Joint,

	"Material":                C.NeckMaterial,
	"NECK":                    C.NeckMaterial,
    "Neck":                    C.NeckMaterial,
	"Neck Wood":               C.NeckMaterial,
    "Neck Material":           C.NeckMaterial,

	"PICKUPS":                 C.Pickups,
    "Pickups":                 C.Pickups,
    "Pickup":                  C.Pickups,
    "Pickup(Neck, Middle, Bridge)":C.Pickups,
	"Bass Pickup":             C.NeckPickup,
    "Neck Pickup":             C.NeckPickup,
    "Neck P ickup":            C.NeckPickup,
	"Middle Pickup":           C.CenterPickup,
	"Treble Pickup":           C.BridgePickup,
    "Bridge Pickup":           C.BridgePickup,

	"Price":                   C.Price,
    "PRICE":                   C.Price,

	"SCALE":                   C.ScaleLengthMM,
    "Scale":                   C.ScaleLengthMM,
	"Scale Length":            C.ScaleLengthMM,

    "Series":                  C.Series,
}

var regWood = regexp.MustCompile(`\s+`)
// 木材コードを探しだす
func searchWoodCode(s string) int {
	trimed := regWood.ReplaceAllString(s, "")

	for _, wood := range GetWoods() {
		if strings.Contains(strings.ToLower(trimed), strings.ToLower(wood.Name)) {
			return wood.Code
		}
	}
	return 99 // unknown
}
type wood struct {
	Name string
	Code int
}

// wood materials
func GetWoods() []wood {
	woods := []wood{
		{"HardMaple", 1},
		{"FlameMaple", 2},
		{"FlamedMaple", 2},
		{"QuiltedMaple", 3},
		{"BirdseyeMaple", 4},
		{"RoastedMaple", 5},
		{"Maple", 6},
		{"HonduranMahogany", 7},
		{"Mahogany", 8},
		{"Sapele", 9},
		{"Korina", 10},
		{"WhiteKorina", 11},
		{"Alder", 12},
		{"Ash", 13},
		{"Basswood", 14},
		{"Linden", 14},
		{"Poplar", 15},
		{"Spruce", 16},
		{"Cedar", 17},
		{"IndianRosewood", 18},
		{"BrazilianRosewood", 19},
		{"Rosewood", 20},
		{"PauFerro", 21},
		{"Ovangkol", 22},
		{"Ebony", 23},
		{"Walnut", 24},
		{"Padauk", 25},
		{"Koa", 26},
		{"Nato", 27},
		{"Agathis", 28},
		{"Bubinga", 29},
		{"Wenge", 30},
		{"Purpleheart", 31},
		{"Zebrawood", 32},
		{"Okoume", 33},
		{"Meranti", 34},
		{"Sakura", 35},
		{"Tochi", 36},
	}
	return woods
}