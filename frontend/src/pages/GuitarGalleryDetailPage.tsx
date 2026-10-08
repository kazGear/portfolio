import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { api } from "../lib/apiClient";
import { Guitar, GuitarsResponse } from "../types/Guitar";
import useApiErrorHandler from "../hooks/useApiErrorHandler";
import { PUBLIC_API_BASE_URL } from "../config/env";
import styled from "styled-components";
import { HTML_HEAD_DATA, SIZE } from "../lib/Constants";
import GuitarDetailPageWide from "../components/guitarGalleryPage/GuitarDetailPageWide";
import GuitarDetailPageMobile from "../components/guitarGalleryPage/GuitarDetailPageMobile";

const DetailPageMobile = styled.div`
    display: none;

    @media (max-width: ${SIZE.MOBILE_LAYOUT_BREAKPOINT}) {
        display: block;
    }

    background: rgba(255, 255, 255, 0.2);
    width: 100%;
    height: 100%;
`;
const DetailPageWide = styled.div`
    display: block;

    @media (max-width: ${SIZE.MOBILE_LAYOUT_BREAKPOINT}) {
        display: none;
    }
`;

const GuitarGalleryDetailPage = () => {
    let { makerCd, name, color } = useParams();
    const [ guitar, setGuitar ]  = useState<Guitar | null>(null);

    // 特定のギター情報を取得
    useEffect(() => {
        // 初期画面用、条件なし検索
        const url = `${PUBLIC_API_BASE_URL}/public/v1/guitars?makerCd=${makerCd}&name=${name}&color=${color}`;

        api.GET<GuitarsResponse>(url)
           .then(result => {
               if (result?.guitars !== undefined && result?.guitars.length >= 1) {
                   setGuitar(result?.guitars[0]); // pk検索なのでギターは１本しか返ってこない
               } else {
                   setGuitar(null);
               }
            })
           .catch(useApiErrorHandler);
    }, [makerCd, name, color]);

    // meta data
    useEffect(() => {
        const description = document.querySelector('meta[name="description"]');

        if (guitar) {
            const guitarFeatures = `${guitar.makerName} | ${guitar.name} | ${guitar.color}`;

            document.title = `${guitarFeatures}` + " | " + HTML_HEAD_DATA.DEFAULT_TITLE;

            description?.setAttribute(
                "content", `【${guitarFeatures}】` + `価格:${guitar.price.toLocaleString()}円。${guitar.comment}`
            );
        }

        // meta data 初期化
        return () => {
            document.title = HTML_HEAD_DATA.DEFAULT_TITLE;
            description?.setAttribute("content", HTML_HEAD_DATA.DEFAULT_DESCRIPTION);
        }
    }, [guitar]);

    return (
        <div>
            {/* PC, タブレット向け */}
            <DetailPageWide>
                <GuitarDetailPageWide guitar={guitar}/>
            </DetailPageWide>

            {/* モバイル向け */}
            <DetailPageMobile>
                <GuitarDetailPageMobile guitar={guitar}/>
            </DetailPageMobile>
        </div>
    );
}

export default GuitarGalleryDetailPage;