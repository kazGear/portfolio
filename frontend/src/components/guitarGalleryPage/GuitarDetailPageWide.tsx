import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import { SIZE } from "../../lib/Constants";
import GuitarSpec from "./GuitarSpec";
import CommonZoomableImage from "../common/CommonZoomableImage";
import { parseGuitarPrice } from "./GuitarFuncs";
import CommonFrame from "../common/CommonFrame";
import { Link } from "react-router-dom";

const P = styled.p`
    overflow-y: auto;
    font-size: 14px;
    width: 100%;
    height: 30%;
`;

interface ArgProps {
    guitar : Guitar | null;
}

const GuitarDetailPageWide = ({guitar}: ArgProps) => {
    return (
        <CommonFrame styleObj={{
            margin: 0,
            borderRadius: 0,
            height: `calc(100vh - ${SIZE.HEADER_HEIGHT} - 2px)`,
            overflowY: "hidden"
        }}>
            <Link to={"/GuitarGalleryPage"} target="_blank">Guitar Gallery へ</Link>
            {
                guitar !== null ? (
                    <div style={{display: "flex", justifyContent: "space-evenly"}}>
                        <GuitarSpec selectedGuitars={guitar}/>

                        <div style={{width: "50%", margin: "20px"}}>
                            <p style={{marginTop: 0}}>最終更新日：{guitar?.updated}</p>

                            <CommonZoomableImage
                                imgURL={guitar?.src}
                                alt={guitar?.makerName + " | " + guitar?.name + " | " + guitar?.color}
                                width={"450px"}
                                height={"300px"}
                                zoomRate={300}
                            />

                            <h2 style={{margin: "0px"}}>
                                price:&emsp;{parseGuitarPrice(guitar?.price!)}
                            </h2>

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
export default GuitarDetailPageWide;