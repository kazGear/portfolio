import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import CommonZoomableImage from "../common/CommonZoomableImage";
import { parseGuitarPrice } from "./GuitarFuncs";

const P = styled.p`
    overflow-y: auto;
    font-size: 14px;
    width: 100%;
    height: 30%;
`;

interface ArgProps {
    selectedGuitars: Guitar | null;
}

const DetailLeftSide = ({selectedGuitars}: ArgProps) => {
    const guitar = selectedGuitars;

    return (
        <div style={{width: "50%", margin: "0px 20px 0px 40px"}}>
            <p>最終更新日：{guitar?.updated}</p>

            <CommonZoomableImage imgURL={guitar?.src}
                                 alt={guitar?.name}
                                 width={400}
                                 height={250}
                                 zoomRate={400}/>

            <h3 style={{margin: "0px"}}>
                price:&emsp;{parseGuitarPrice(guitar?.price!)}
            </h3>

            <p>comment.</p>

            <P>{guitar?.comment}</P>
        </div>
    );
}
export default DetailLeftSide;