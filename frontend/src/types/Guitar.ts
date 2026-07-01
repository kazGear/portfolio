export interface GuitarsRequest {
    makerCd?:            number;
    name?:               string;
    series?:             string;
    colorCd?:            number;
    bodyMaterialTopCd?:  number;
    bodyMaterialBackCd?: number;
    minPrice?:           number;
    maxPrice?:           number;
    order?:              string;
    sort?:               string;
    page?:               number;
    pageSize?:           number;
};

export interface GuitarsResponse {
    totalCount: number;
    page:       number;
    pageSize:   number;
    totalPages: number;
    hasPrev:    boolean;
    hasNext:    boolean;
    guitars:    Guitar[];
}

export interface Guitar {
    maker:            number;
    name:             string;
    color:            string;
    colorCd:          number;
    bodyFinish:       string;
    bodyMaterial:     string;
    bodyMaterialTop:  number;
    bodyMaterialBack: number;
    bridge:           string;
    controls:         string;
    comment:          string;
    fingerboard:      number;
    fretCount:        number;
    inlays:           string;
    joint:            string;
    neckMaterial:     number;
    pickups:          string;
    price:            number;
    scaleLengthMm:    number;
    series:           string;
    src:              string;
    weight:           number;
};