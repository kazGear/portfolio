import { useCallback, useEffect, useState } from "react";
import { Code } from "../types/Code";
import { api } from "../lib/apiClient";
import { Guitar, GuitarParams, GuitarsResponse } from "../types/Guitar";
import { useGuitarParams } from "../hooks/useGuitarParams";
import { createQueryParams } from "../components/guitarGalleryPage/GuitarFuncs";
import CommonFrame from "../components/common/CommonFrame";
import GuitarCards from "../components/guitarGalleryPage/GuitarCards";
import SearchConditions from "../components/guitarGalleryPage/SearchConditions";
import DetailModal from "../components/guitarGalleryPage/DetailModal";
import { PUBLIC_API_BASE_URL } from "../config/env"

const GuitarGalleryPage = () => {
    // プルダウン用 params
    const [makers, setMakers]               = useState<Code[] | null>([]);
    const [series, setSeries]               = useState<Code[] | null>([]);
    const [colors, setColors]               = useState<Code[] | null>([]);
    const [bodyMaterials, setBodyMaterials] = useState<Code[] | null>([]);

    const [selectedGuitar, setSelectedGuitar] = useState<Guitar | null>(null);
    const [guitars, setGuitars]               = useState<GuitarsResponse | null>(null);

    const [isShowDetail, setIsShowDetail] = useState<boolean>(false);

    const gParams: GuitarParams = useGuitarParams();

    // プルダウンデータ等取得
    useEffect(() => {
        api.GET<Code[]>(`${PUBLIC_API_BASE_URL}/public/v1/makers`).then(result => setMakers(result));
        api.GET<Code[]>(`${PUBLIC_API_BASE_URL}/public/v1/Colors`).then(result => setColors(result));
        api.GET<Code[]>(`${PUBLIC_API_BASE_URL}/public/v1/bodyMaterials`).then(result => setBodyMaterials(result));
        // 初期画面用、条件なし検索
        api.GET<GuitarsResponse>(`${PUBLIC_API_BASE_URL}/public/v1/guitars?`).then(result => setGuitars(result));
    }, [])

    // 変動プルダウンデータ取得
    useEffect(() => {
        if (gParams.makerCd === 0) {
            setSeries([]);
            gParams.setSeries("")
            return;
        }
        api.GET<Code[]>(
            `${PUBLIC_API_BASE_URL}/public/v1/series?makerCd=${gParams.makerCd}`
        ).then(result => setSeries(result));

        gParams.setSeries("") // 初期化しないと、他メーカーのシリーズを選択したままになってしまう。
    }, [gParams.makerCd])

    // ギターデータ取得
    const guitarSearchHandler = useCallback( async (gParams: GuitarParams) => {
        const queryParams = createQueryParams(gParams);
        const resGuitars  = await api.GET<GuitarsResponse>(
            `${PUBLIC_API_BASE_URL}/public/v1/guitars?${queryParams.toString()}`
        );
        setGuitars(resGuitars);
    }, []);

    // 選択ギターpk取得
    const getSelectedGuitarHandler = useCallback((guitar: Guitar | null) => {
        setSelectedGuitar(guitar)
        setIsShowDetail(true)
    }, []);

    return (
        <div style={{display: "flex"}}>
            <CommonFrame styleObj={{width: "20%", minWidth: "280px", height: "85vh", marginLeft: "20px"}}>
                <SearchConditions guitarRes={guitars}
                                  guitarParams={gParams}
                                  makers={makers}
                                  colors={colors}
                                  series={series}
                                  bodyMaterials={bodyMaterials}
                                  callback={guitarSearchHandler}
                                  />
            </CommonFrame>
            <CommonFrame styleObj={{width: "80%", minWidth: "280px",height: "85vh", margin: "20px 20px 0px 10px"}}>
                <GuitarCards guitarsRes={guitars}
                             callback={getSelectedGuitarHandler}>
                </GuitarCards>
            </CommonFrame>

            <DetailModal selectedGuitars={selectedGuitar}
                         isShow={isShowDetail}
                         callback={setIsShowDetail}>
            </DetailModal>
        </div>
    );
}

export default GuitarGalleryPage;