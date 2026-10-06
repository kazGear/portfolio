import { Guitar } from "../../types/Guitar";
import { SIZE } from "../../lib/Constants";
import GuitarSpec from "./GuitarSpec";
import CommonZoomableImage from "../common/CommonZoomableImage";
import { parseGuitarPrice } from "./GuitarFuncs";
import CommonFrame from "../common/CommonFrame";
import { Link } from "react-router-dom";

interface ArgProps {
    guitar : Guitar | null;
}

const GuitarDetailPageMobile = ({guitar}: ArgProps) => {
    return (
        <CommonFrame styleObj={{
            margin: 0,
            borderRadius: 0,
            height: `calc(100vh - ${SIZE.HEADER_HEIGHT} - 2px)`,
        }}>
            <Link to={"/GuitarGalleryPage"} target="_blank">Guitar Gallery へ</Link>
            {
                guitar !== null ? (
                    <div style={{margin: "20px"}}>
                        <p style={{marginTop: 0, textAlign: "right", fontSize: "14  px"}}>最終更新日：{guitar?.updated}</p>

                        <div style={{ textAlign: "center" }}>
                            <h2 style={{margin: 0, textAlign: "left"}}>{guitar.makerName}</h2>
                            <h1 style={{margin: 0}}>{guitar.name}</h1>
                            <h3 style={{margin: 0}}>{guitar.color}</h3>
                        </div>

                        <CommonZoomableImage
                            imgURL={guitar?.src}
                            alt={guitar?.makerName + " | " + guitar?.name + " | " + guitar?.color}
                            width={"100%"}
                            height={"300px"}
                            zoomRate={400}
                        />

                        <h2 style={{margin: "0px"}}>
                            price:&emsp;{parseGuitarPrice(guitar?.price!)}
                        </h2>

                        <GuitarSpec selectedGuitars={guitar}/>

                        <p>{guitar?.comment}</p>
                    </div>
                ) : (
                    <h1>ギター情報の取得に失敗しました。</h1>
                )
            }
            <Link to={"/GuitarGalleryPage"} target="_blank">Guitar Gallery へ</Link>
        </CommonFrame>
    );
}
export default GuitarDetailPageMobile;