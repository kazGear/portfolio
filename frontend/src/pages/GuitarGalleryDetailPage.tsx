import { Link, useParams } from "react-router-dom";
import CommonFrame from "../components/common/CommonFrame";
import { useEffect, useState } from "react";
import { api } from "../lib/apiClient";
import { Guitar, GuitarsResponse } from "../types/Guitar";
import useApiErrorHandler from "../hooks/useApiErrorHandler";
import { PUBLIC_API_BASE_URL } from "../config/env";
import GuitarSpec from "../components/guitarGalleryPage/GuitarSpec";
import CommonZoomableImage from "../components/common/CommonZoomableImage";
import { parseGuitarPrice } from "../components/guitarGalleryPage/GuitarFuncs";
import styled from "styled-components";

const P = styled.p`
    overflow-y: auto;
    font-size: 14px;
    width: 100%;
    height: 30%;
    margin-bottom: 0px;
`;

const GuitarGalleryPage = () => {
    const { makerCd, name, color } = useParams();
    const [ guitar, setGuitar ]    = useState<Guitar | null>(null);

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
    }, []);

    return (
        <CommonFrame styleObj={{margin: 0, borderRadius: 0, height: "92vh", overflowY: "hidden"}}>
            <Link to={"/GuitarGalleryPage"} target="_blank">Guitar Gallery へ</Link>
            {
                guitar !== null ? (
                    <div style={{display: "flex", justifyContent: "space-evenly"}}>
                        <GuitarSpec selectedGuitars={guitar}/>

                        <div style={{width: "50%", margin: "20px"}}>
                            <p style={{marginTop: 0}}>最終更新日：{guitar?.updated}</p>

                            <CommonZoomableImage
                                imgURL={guitar?.src}
                                alt={guitar?.name}
                                width={450}
                                height={300}
                                zoomRate={300}/>

                            <h3 style={{margin: "0px"}}>
                                price:&emsp;{parseGuitarPrice(guitar?.price!)}
                            </h3>

                            <P>{guitar?.comment}</P>
                        </div>
                    </div>
                ) : (
                    <h1>ギター情報の取得に失敗しました。</h1>
                )
            }
        </CommonFrame>
    );
}

export default GuitarGalleryPage;