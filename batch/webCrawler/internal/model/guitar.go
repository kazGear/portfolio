package model

import "fmt"

type Guitar struct {
	Maker         int     `json:"maker"`
	Name          string  `json:"name"`
	SubName       string  `json:"sub_name"`
	BodyFinish    string  `json:"body_finish"`
	BodyMaterial  int     `json:"body_material"`
	Bridge        string  `json:"bridge"`
	Color         string  `json:"color"`
	Controls      string  `json:"controls"`
	Comment       string  `json:"comment"`
	Fingerboard   int     `json:"fingerboard"`
	FretCount     int     `json:"fret_count"`
	Inlays        string  `json:"inlays"`
	Joint         string  `json:"joint"`
	NeckMaterial  int     `json:"neck_material"`
	Pickups       string  `json:"pickups"`
	Price         int     `json:"price"`
	ScaleLengthMM int     `json:"scale_length_mm"`
	Series        string  `json:"series"`
	Weight        float64 `json:"weight"`
}

func (g *Guitar) String() string {
	return fmt.Sprintf("maker: %v, name: %v, price: %d", g.Maker, g.Name, g.Price)
}