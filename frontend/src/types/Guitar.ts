export interface GuitarsRequest {
    MakerCd?:            number;
    Name?:               string;
    Series?:             string;
    ColorCd?:            number;
    BodyMaterialTopCd?:  number;
    BodyMaterialBackCd?: number;
    MinPrice?:           number;
    MaxPrice?:           number;
    Order?:               string;
    Sort?:               string;
    Page:                number;
    PageSize:            number;
};
export interface GuitarsResponse {
    TotalCount: number;
    Page:       number;
    PageSize:   number;
    TotalPages: number;
    HasPrev:    boolean;
    HasNext:    boolean;
    Guitars:    Guitar[];
}

export interface Guitar {
    Maker:            number;
    Name:             string;
    Color:            string;
    ColorCd:          number;
    BodyFinish:       string;
    BodyMaterial:     string;
    BodyMaterialTop:  number;
    BodyMaterialBack: number;
    Bridge:           string;
    Controls:         string;
    Comment:          string;
    Fingerboard:      number;
    FretCount:        number;
    Inlays:           string;
    Joint:            string;
    NeckMaterial:     number;
    Pickups:          string;
    Price:            number;
    ScaleLengthMm:    number;
    Series:           string;
    Src:              string;
    Weight:           number;
};