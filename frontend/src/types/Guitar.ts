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
    makerName:        string;
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
    fingerboardName:  string;
    fretCount:        number;
    inlays:           string;
    joint:            string;
    neckMaterial:     number;
    neckMaterialName: number;
    pickups:          string;
    price:            number;
    scaleLengthMm:    number;
    series:           string;
    src:              string;
    weight:           number;
    updated:          string;
};

export type GuitarParams = {
    makerCd:            number;
    name:               string;
    colorCd:            number;
    series:             string;
    bodyMaterialTopCd:  number;
    bodyMaterialBackCd: number;
    minPrice:           number;
    maxPrice:           number;
    order:              string;
    sort:               string;
    page:               number;
    pageSize:           number;

    setMakerCd:            React.Dispatch<React.SetStateAction<number>>;
    setName:               React.Dispatch<React.SetStateAction<string>>;
    setColorCd:            React.Dispatch<React.SetStateAction<number>>;
    setSeries:             React.Dispatch<React.SetStateAction<string>>
    setBodyMaterialTopCd:  React.Dispatch<React.SetStateAction<number>>;
    setBodyMaterialBackCd: React.Dispatch<React.SetStateAction<number>>;
    setMinPrice:           React.Dispatch<React.SetStateAction<number>>;
    setMaxPrice:           React.Dispatch<React.SetStateAction<number>>;
    setOrder:              React.Dispatch<React.SetStateAction<string>>;
    setSort:               React.Dispatch<React.SetStateAction<string>>;
    setPage:               React.Dispatch<React.SetStateAction<number>>;
    setPageSize:           React.Dispatch<React.SetStateAction<number>>;
};