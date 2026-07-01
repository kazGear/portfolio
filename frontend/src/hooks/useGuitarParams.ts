import { useState } from "react";

export const useGuitarParams = () => {
    const [makerCd, setMakerCd]                       = useState<number>(0);
    const [name, setName]                             = useState<string>("");
    const [series, setSeries]                         = useState<string>("");
    const [colorCd, setColorCd]                       = useState<number>(0);
    const [bodyMaterialTopCd, setBodyMaterialTopCd]   = useState<number>(0);
    const [bodyMaterialBackCd, setBodyMaterialBackCd] = useState<number>(0);
    const [minPrice, setMinPrice]                     = useState<number>(0);
    const [maxPrice, setMaxPrice]                     = useState<number>(0);
    const [order, setOrder]                           = useState<string>("ASC");
    const [sort, setSort]                             = useState<string>("");
    const [page, setPage]                             = useState<number>(1);
    const [pageSize, setPageSize]                     = useState<number>(25);

    const params: GuitarParams = {
        makerCd:            makerCd,
        name:               name,
        series:             series,
        colorCd:            colorCd,
        bodyMaterialTopCd:  bodyMaterialTopCd,
        bodyMaterialBackCd: bodyMaterialBackCd,
        minPrice:           minPrice,
        maxPrice:           maxPrice,
        order:              order,
        sort:               sort,
        page:               page,
        pageSize:           pageSize,

        setMakerCd:            setMakerCd,
        setName:               setName,
        setSeries:             setSeries,
        setColorCd:            setColorCd,
        setBodyMaterialTopCd:  setBodyMaterialTopCd,
        setBodyMaterialBackCd: setBodyMaterialBackCd,
        setMinPrice:           setMinPrice,
        setMaxPrice:           setMaxPrice,
        setOrder:              setOrder,
        setSort:               setSort,
        setPage:               setPage,
        setPageSize:           setPageSize,
    };
    return params;
}

type GuitarParams = {
    makerCd:            number;
    name:               string;
    series:             string;
    colorCd:            number;
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
    setSeries:             React.Dispatch<React.SetStateAction<string>>
    setColorCd:            React.Dispatch<React.SetStateAction<number>>;
    setBodyMaterialTopCd:  React.Dispatch<React.SetStateAction<number>>;
    setBodyMaterialBackCd: React.Dispatch<React.SetStateAction<number>>;
    setMinPrice:           React.Dispatch<React.SetStateAction<number>>;
    setMaxPrice:           React.Dispatch<React.SetStateAction<number>>;
    setOrder:              React.Dispatch<React.SetStateAction<string>>;
    setSort:               React.Dispatch<React.SetStateAction<string>>;
    setPage:               React.Dispatch<React.SetStateAction<number>>;
    setPageSize:           React.Dispatch<React.SetStateAction<number>>;
};