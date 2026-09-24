import { Link, useParams } from "react-router-dom";
import CommonFrame from "../components/common/CommonFrame";
import { useEffect, useState } from "react";
import { api } from "../lib/apiClient";
import { Guitar, GuitarsResponse } from "../types/Guitar";
import useApiErrorHandler from "../hooks/useApiErrorHandler";
import { PUBLIC_API_BASE_URL } from "../config/env";

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
        <CommonFrame>
            <h1>req param makerCd {makerCd}</h1>
            <h1>req param guitar {name}</h1>
            <h1>req param color {color}</h1>
            <Link to={"/GuitarGalleryPage"} target="_blank">Guitar Gallery へ</Link>
            {
                guitar !== null ? (
                    <>
                        <h1>{guitar.name}</h1>
                        <img
                            style={{width:"50%", height:"50%", objectFit: "contain"}}
                            src={guitar.src}
                            alt={guitar.maker + " " + guitar.name}
                            />
                    </>
                ) : (
                    <h1>ギター情報の取得に失敗しました。</h1>
                )
            }
        </CommonFrame>
    );
}

export default GuitarGalleryPage;