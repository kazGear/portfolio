package model

import "fmt"

type Guitar struct {
    Maker             int     `db:"maker"               json:"maker"`
    Name              string  `db:"name"                json:"name"`
    BodyFinish        string  `db:"body_finish"         json:"body_finish"`
    BodyMaterial      string  `db:"body_material"       json:"body_material"`
    BodyMaterialFront int     `db:"body_material_front" json:"body_material_front"`
    BodyMaterialBack  int     `db:"body_material_back"  json:"body_material_back"`
    Bridge            string  `db:"bridge"              json:"bridge"`
    Color             string  `db:"color"               json:"color"`
    Controls          string  `db:"controls"            json:"controls"`
    Comment           string  `db:"comment"             json:"comment"`
    Fingerboard       int     `db:"fingerboard"         json:"fingerboard"`
    FretCount         int     `db:"fret_count"          json:"fret_count"`
    Inlays            string  `db:"inlays"              json:"inlays"`
    Joint             string  `db:"joint"               json:"joint"`
    NeckMaterial      int     `db:"neck_material"       json:"neck_material"`
    Pickups           string  `db:"pickups"             json:"pickups"`
    Price             int     `db:"price"               json:"price"`
    ScaleLengthMM     int     `db:"scale_length_mm"     json:"scale_length_mm"`
    Series            string  `db:"series"              json:"series"`
    Src               string  `db:"src"                 json:"src"`
    Weight            float64 `db:"weight"              json:"weight"`
}


func (g *Guitar) String() string {
	return fmt.Sprintf("maker: %v, name: %v, price: %d", g.Maker, g.Name, g.Price)
}
