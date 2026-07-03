import { Guitar } from "../../types/Guitar";

interface ArgProps {
    selectedGuitars: Guitar | null;
}

const DetailLeftSide = ({selectedGuitars}: ArgProps) => {
    const guitar = selectedGuitars;

    return (
        <div style={{width: "50%", margin: "0px 40px"}}>
            <img src={guitar?.src}
                 alt={guitar?.name}
                 style={{
                    width:"100%",
                    height:"40%",
                    objectFit: "contain",
                    marginTop: "20px"
                }}/>
            <h2 style={{margin: "0px"}}>
                price:&emsp;{guitar?.price} 円
            </h2>
            <p>comment.</p>
            <p style={{
                overflowY: "auto",
                fontSize: "14px",
                height: "30%"
                }}>
                {guitar?.comment}
            </p>
        </div>
    );
}
export default DetailLeftSide;