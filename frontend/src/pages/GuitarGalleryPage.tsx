import { useCallback, useEffect, useState } from "react";
import { Code } from "../types/Code";
import { api } from "../lib/apiClient";
import { Guitar, GuitarParams, GuitarsResponse } from "../types/Guitar";
import { useGuitarParams } from "../hooks/useGuitarParams";
import { createQueryParamsGuitar } from "../components/guitarGalleryPage/GuitarFuncs";
import CommonFrame from "../components/common/CommonFrame";
import GuitarCards from "../components/guitarGalleryPage/GuitarCards";
import SearchConditionsGuitar from "../components/guitarGalleryPage/SearchConditionsGuitar";
import GuitarDetailModal from "../components/guitarGalleryPage/GuitarDetailModal";
import { PUBLIC_API_BASE_URL } from "../config/env"
import useApiErrorHandler from "../hooks/useApiErrorHandler";
import CommonNowLoading from "../components/common/CommonNowLoading";
import { COLORS, HTML_HEAD_DATA, SIZE } from "../lib/Constants"
import styled from "styled-components";
import CommonButton from "../components/common/CommonButton";
import SelectorPage from "../components/guitarGalleryPage/SelectorPage";
import SearchConditionsGuitarMobile from "../components/guitarGalleryPage/SearchConditionsGuitarMobile";

const GuitarGalleryWide = styled.div`
    display: flex;

    @media (max-width: ${SIZE.MOBILE_LAYOUT_BREAKPOINT}) {
        display: none;
    }
`;
const GuitarGalleryMobile = styled.div`
    display: none;

    @media (max-width: ${SIZE.MOBILE_LAYOUT_BREAKPOINT}) {
        display: block;
    }
`;
const SearchViewMobile = styled.div`
    background: ${COLORS.BASE_BACKGROUND};
`;

const GuitarGalleryPage = () => {
    // プルダウン用 params
    const [makers, setMakers]               = useState<Code[]>([]);
    const [series, setSeries]               = useState<Code[]>([]);
    const [colors, setColors]               = useState<Code[]>([]);
    const [bodyMaterials, setBodyMaterials] = useState<Code[]>([]);

    const [selectedGuitar, setSelectedGuitar] = useState<Guitar | null>(null);
    const [guitars, setGuitars]               = useState<GuitarsResponse | null>(null);

    const [isShowDetail, setIsShowDetail]             = useState<boolean>(false);
    const [isShowMobileSearch, setIsShowMobileSearch] = useState<boolean>(false);

    const gParams: GuitarParams = useGuitarParams();
    const errorHandler          = useApiErrorHandler();

    // プルダウンデータ等取得
    useEffect(() => {
        api.GET<Code[]>(`${PUBLIC_API_BASE_URL}/public/v1/makers`)
           .then(result => setMakers(result ?? []))
           .catch(errorHandler);
        api.GET<Code[]>(`${PUBLIC_API_BASE_URL}/public/v1/Colors`)
           .then(result => setColors(result ?? []))
           .catch(errorHandler);
        api.GET<Code[]>(`${PUBLIC_API_BASE_URL}/public/v1/bodyMaterials`)
           .then(result => setBodyMaterials(result ?? []))
           .catch(errorHandler)

           // 初期画面用、条件なし検索
        api.GET<GuitarsResponse>(`${PUBLIC_API_BASE_URL}/public/v1/guitars?`)
           .then(result => setGuitars(result))
           .catch(errorHandler);
    }, [])

    // 変動プルダウンデータ取得
    useEffect(() => {
        if (gParams.makerCd === 0) {
            setSeries([]);
            gParams.setSeries("")
            return;
        }

        api.GET<Code[]>(`${PUBLIC_API_BASE_URL}/public/v1/series?makerCd=${gParams.makerCd}`)
           .then(result => setSeries(result ?? []))
           .catch(errorHandler);

        gParams.setSeries("") // 初期化しないと、他メーカーのシリーズを選択したままになってしまう。
    }, [gParams.makerCd])

    // ギターデータ取得
    const guitarSearchHandler = async (gParams: GuitarParams) => {
        const queryParams = createQueryParamsGuitar(gParams);
        const resGuitars  = await api.GET<GuitarsResponse>(
            `${PUBLIC_API_BASE_URL}/public/v1/guitars?${queryParams.toString()}`
        );
        setGuitars(resGuitars);
    };

    // 選択ギターpk取得
    const getSelectedGuitarHandler = useCallback((guitar: Guitar | null) => {
        setSelectedGuitar(guitar)
        setIsShowDetail(true)
    }, []);

    // meta data
    useEffect(() => {
        document.title = HTML_HEAD_DATA.GUITAR_GALLERY_TITLE;

        const description = document.querySelector('meta[name="description"]');
        description?.setAttribute("content", HTML_HEAD_DATA.GUITAR_GALLERY_DESCRIPTION);

        // meta data 初期化
        return () => {
            document.title = HTML_HEAD_DATA.DEFAULT_TITLE;
            description?.setAttribute("content", HTML_HEAD_DATA.DEFAULT_DESCRIPTION);
        }
    }, []);

    // 検索条件を選択した時点で検索実行
    useEffect(() => {
        void guitarSearchHandler(gParams)
        gParams.setPage(1)
    }, [
        gParams.makerCd,
        gParams.colorCd,
        gParams.series,
        gParams.name,
        gParams.bodyMaterialTopCd,
        gParams.bodyMaterialBackCd,
        gParams.minPrice,
        gParams.maxPrice,
        gParams.pageSize,
    ])

    // ページ初期化なしの検索
    useEffect(() => {
        void guitarSearchHandler(gParams)
    }, [
        gParams.sort,
        gParams.order,
        gParams.page,
    ])

    // モバイル用検索画面の切り替え
    const toggleSearchViewMobile = () => {
        setIsShowMobileSearch(!isShowMobileSearch)
    }

    return (
        <div>
            {/* PC, タブレット向け */}
            <GuitarGalleryWide>
                <CommonFrame styleObj={{width: "20%", minWidth: "280px", height: "87vh", margin: "20px 0px 0px 20px"}}>
                    <SearchConditionsGuitar
                        guitarRes={guitars}
                        guitarParams={gParams}
                        makers={makers}
                        colors={colors}
                        series={series}
                        bodyMaterials={bodyMaterials}
                    />
                </CommonFrame>
                <CommonFrame styleObj={{width: "80%", minWidth: "280px",height: "87vh", margin: "20px 20px 0px 10px"}}>
                    {
                        guitars !== null ? (
                            <GuitarCards guitarsRes={guitars}
                                        callback={getSelectedGuitarHandler}>
                            </GuitarCards>
                        ) : (
                            <div style={{textAlign: "center", marginTop: "200px"}}>
                                <CommonNowLoading alt="guitar cards"/>
                            </div>
                        )}
                </CommonFrame>

                <GuitarDetailModal
                    selectedGuitar={selectedGuitar}
                    isShow={isShowDetail}
                    callback={setIsShowDetail}>
                </GuitarDetailModal>
            </GuitarGalleryWide>

            {/* モバイル向け */}
            <GuitarGalleryMobile>
                <CommonButton
                    text={isShowMobileSearch ? "閉じる" : "検索"}
                    onClick={() => toggleSearchViewMobile()}
                    styleObj={{
                        width: "100%",
                        height: "30px",
                        margin: 0,
                        borderRadius: 0,
                    }}
                    ></CommonButton>

                { isShowMobileSearch ? (
                        <div className={isShowMobileSearch ? "searchViewMobileMove" : ""}>
                            <SearchViewMobile>
                                <SearchConditionsGuitarMobile
                                    guitarParams={gParams}
                                    makers={makers}
                                    colors={colors}
                                    series={series}
                                />
                            </SearchViewMobile>
                        </div>
                    ) : ""
                }

                <CommonFrame styleObj={{width: "95%", margin: "10px", boxShadow: "none"}}>
                    {
                        guitars !== null ? (
                            <div>

                                <SelectorPage
                                    guitarRes={guitars}
                                    guitarParams={gParams}
                                    styleObj={{margin: "15px auto"}}
                                ></SelectorPage>

                                <GuitarCards guitarsRes={guitars}
                                             callback={getSelectedGuitarHandler}>
                                </GuitarCards>

                                <SelectorPage
                                    guitarRes={guitars}
                                    guitarParams={gParams}
                                    styleObj={{margin: "15px auto"}}
                                ></SelectorPage>
                            </div>
                        ) : (
                            <div style={{textAlign: "center", marginTop: "200px"}}>
                                <CommonNowLoading alt="guitar cards"/>
                            </div>
                        )}
                </CommonFrame>
            </GuitarGalleryMobile>
        </div>
    );
}

export default GuitarGalleryPage;