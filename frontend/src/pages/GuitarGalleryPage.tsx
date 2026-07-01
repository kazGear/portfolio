import { useCallback, useEffect, useState } from "react";
import styled from "styled-components";
import { Code } from "../types/Code";
import { api } from "../lib/apiClient";
import { GuitarParams, GuitarsResponse } from "../types/Guitar";
import { useGuitarParams } from "../hooks/useGuitarParams";
import { createQueryParams } from "../components/GuitarGalleryPage/GuitarFuncs";

const SdivEditFrame = styled.div`
    width: 100%;
    margin-top: 80px;
`;

const GuitarGalleryPage = () => {
    const [makers, setMakers]               = useState<Code[] | null>([]);
    const [series, setSeries]               = useState<Code[] | null>([]);
    const [colors, setColors]               = useState<Code[] | null>([]);
    const [bodyMaterials, setBodyMaterials] = useState<Code[] | null>([]);
    const [guitars, setGuitars]             = useState<GuitarsResponse | null>(null);

    const gParams: GuitarParams = useGuitarParams();

    // プルダウンデータ取得
    useEffect(() => {
        api.GET<Code[]>("https://localhost:7170/api/v1/makers").then(result => setMakers(result));
        api.GET<Code[]>("https://localhost:7170/api/v1/Colors").then(result => setColors(result));
        api.GET<Code[]>("https://localhost:7170/api/v1/bodyMaterials").then(result => setBodyMaterials(result));
    }, [])

    // 変動プルダウンデータ取得
    useEffect(() => {
        if (gParams.makerCd === 0) {
            setSeries([]);
            return;
        }

        api.GET<Code[]>(
            `https://localhost:7170/api/v1/series?makerCd=${gParams.makerCd}`
        ).then(result => setSeries(result));
    }, [gParams.makerCd])

    // ギターデータ取得
    const guitarSearchHandler = useCallback( async (gParams: GuitarParams) => {
        const queryParams = createQueryParams(gParams);
        const resGuitars = await api.GET<GuitarsResponse>(
            `https://localhost:7170/api/v1/guitars?${queryParams.toString()}`
        );
        setGuitars(resGuitars);
    }, []);
    console.log(makers);
    console.log(series);
    console.log(guitars);
    console.log(colors);
    console.log(bodyMaterials);

    return (
        <>
            <h1>ギターページ</h1>
            <h1>工事中</h1>
            {
                guitars?.guitars.map(g => (<span key={g.maker + g.name + g.color}>{g.name}, {g.price}. </span>))
            }
        </>
    );
}

export default GuitarGalleryPage;