import { Guitar } from "../../types/Guitar";
import CommonZoomableImage from "../common/CommonZoomableImage";
import { parseGuitarPrice } from "./GuitarFuncs";

interface ArgProps {
    selectedGuitars: Guitar | null;
}

const DetailLeftSide = ({selectedGuitars}: ArgProps) => {
    const guitar = selectedGuitars;

    return (
        <div style={{width: "50%", margin: "0px 40px"}}>
            <p>最終更新日：{guitar?.updated}</p>

            <CommonZoomableImage imgURL={guitar?.src}
                                 alt={guitar?.name}
                                 width={400}
                                 height={200}
                                 zoomRate={400}/>

            <h2 style={{margin: "0px"}}>
                price:&emsp;{parseGuitarPrice(guitar?.price!)}
            </h2>

            <p>comment.</p>

            <p style={{
                overflowY: "auto",
                fontSize: "14px",
                width: "100%",
                height: "30%"
                }}>
                {guitar?.comment}
            </p>
        </div>
    );
}
export default DetailLeftSide;