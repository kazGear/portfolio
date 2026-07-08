import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import { Guitar } from "../../types/Guitar";
import { getColorString, parsePrice } from "./GuitarFuncs";

const CardFrame = styled.div`
    font-weight: 900;
    width: 240px;
    height: 280px;
    overflow-y: hidden;
    margin: 10px 10px 20px 5px;
    background: ${COLORS.BASE_BACKGROUND};
    border-radius: 15px;
    box-shadow:
        inset 0 1px 0 rgba(255,255,255,0.1),
        inset 0 -2px 10px rgba(0,0,0,0.4);
`;

const Button = styled.button`
    padding: 0;
    background: none;
    border: none;
    cursor: pointer;
`

const H3 = styled.h3`
    margin: 0px 20px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
`

const P = styled.p`
    margin: 0px 20px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
`

interface ArgProps {
    guitar:   Guitar | null;
    callback: (guitar: Guitar | null) => void;
}

const GuitarCard = ({guitar, callback}: ArgProps) => {
    const color = getColorString(guitar?.colorCd);

    let fontShadow = "";
    if (color === "Black") {
        fontShadow = "none"
    } else {
        fontShadow = "-1px -1px 0 #999999, 1px -1px 0 #999999, -1px  1px 0 #999999, 1px  1px 0 #999999"
    }

    return (
        <Button onClick={() => callback(guitar)}>
            <CardFrame>
                <div style={{textAlign: "center", margin: "20px 20px 0px 20px" ,height: "50%"}}>
                    {/* モーダルにギター情報を渡す */}
                    <img style={{width:"90%", height:"90%", objectFit: "contain"}}
                         src={guitar?.src}
                         alt={guitar?.maker + " " + guitar?.name}
                         loading="lazy"
                         />
                </div>
                <div style={{textAlign: "center", height: "50%"}}>
                    <p style={{margin: "0px 20px"}}>{guitar?.makerName}</p>
                    <H3 style={{color: color, textShadow: fontShadow}}>
                        {guitar?.name}
                    </H3>
                    <P>{guitar?.color}</P>
                    <p>{parsePrice(guitar?.price)}</p>
                </div>
            </CardFrame>
        </Button>
    );
}
export default GuitarCard;