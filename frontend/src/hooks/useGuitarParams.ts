import { useState } from "react";
import { GuitarParams } from "../types/Guitar";

export const useGuitarParams = () => {
    const [makerCd, setMakerCd]                       = useState<number>(0);
    const [name, setName]                             = useState<string>("");
    const [colorCd, setColorCd]                       = useState<number>(0);
    const [series, setSeries]                         = useState<string>("");
    const [bodyMaterialTopCd, setBodyMaterialTopCd]   = useState<number>(-1);
    const [bodyMaterialBackCd, setBodyMaterialBackCd] = useState<number>(-1);
    const [minPrice, setMinPrice]                     = useState<number>(-3);
    const [maxPrice, setMaxPrice]                     = useState<number>(100000000);
    const [order, setOrder]                           = useState<string>("ASC");
    const [sort, setSort]                             = useState<string>("price");
    const [page, setPage]                             = useState<number>(1);
    const [pageSize, setPageSize]                     = useState<number>(25);

    const params: GuitarParams = {
        makerCd:            makerCd,
        name:               name,
        colorCd:            colorCd,
        series:             series,
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
        setColorCd:            setColorCd,
        setSeries:             setSeries,
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