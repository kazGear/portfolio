import { Guitar } from "../../types/Guitar";
import { parsePrice } from "./GuitarFuncs";

interface ArgProps {
    selectedGuitars: Guitar | null;
}

const DetailLeftSide = ({selectedGuitars}: ArgProps) => {
    const guitar = selectedGuitars;

    return (
        <div style={{width: "50%", margin: "0px 40px"}}>
            <p>最終更新日：{guitar?.updated}</p>
            <img src={guitar?.src}
                 alt={guitar?.name}
                 style={{
                    width:"100%",
                    height:"35%",
                    objectFit: "contain",
                    marginTop: "10px"
                }}/>
            <h2 style={{margin: "0px"}}>
                price:&emsp;{parsePrice(guitar?.price!)}
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