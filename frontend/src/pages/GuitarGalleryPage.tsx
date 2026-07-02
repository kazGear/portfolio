import { useCallback, useEffect, useState } from "react";
import styled from "styled-components";
import { Code } from "../types/Code";
import { api } from "../lib/apiClient";
import { GuitarParams, GuitarsResponse } from "../types/Guitar";
import { useGuitarParams } from "../hooks/useGuitarParams";
import { createQueryParams } from "../components/GuitarGalleryPage/GuitarFuncs";
import OutSideFrame from "../components/common/OutSideFrame";
import GuitarCards from "../components/GuitarGalleryPage/GuitarCards";

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

    // プルダウンデータ等取得
    useEffect(() => {
        api.GET<Code[]>("https://localhost:7170/api/v1/makers").then(result => setMakers(result));
        api.GET<Code[]>("https://localhost:7170/api/v1/Colors").then(result => setColors(result));
        api.GET<Code[]>("https://localhost:7170/api/v1/bodyMaterials").then(result => setBodyMaterials(result));
        api.GET<GuitarsResponse>("https://localhost:7170/api/v1/guitars?").then(result => setGuitars(result));
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
    // TODO: tmp
    console.log(makers);
    console.log(series);
    console.log(colors);
    console.log(bodyMaterials);
    console.log(guitars);

    return (
        <div style={{display: "flex"}}>
            <OutSideFrame styleObj={{width: "20%", minWidth: "240px", height: "85vh"}}>
                <h1>検索部</h1>
            </OutSideFrame>
            <OutSideFrame styleObj={{width: "80%", minWidth: "280px",height: "85vh"}}>
                <GuitarCards guitarsRes={guitars}></GuitarCards>
            </OutSideFrame>
            <OutSideFrame styleObj={{width: "85vh", height: "85vh", display: "none"}}>
                <h1>モーダル部</h1>
            </OutSideFrame>
        </div>
    );
}

export default GuitarGalleryPage;