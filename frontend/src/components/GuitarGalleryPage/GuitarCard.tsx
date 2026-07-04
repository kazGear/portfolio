import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import { Guitar } from "../../types/Guitar";
import { getColorString, parsePrice } from "./GuitarFuncs";

const ScardFrame = styled.div`
    font-weight: 900;
    width: 220px;
    height: 250px;
    overflow-y: hidden;
    margin: 10px 10px 20px 5px;
    background: ${COLORS.BASE_BACKGROUND};
    border-radius: 15px;
    box-shadow:
        inset 0 1px 0 rgba(255,255,255,0.1),
        inset 0 -2px 10px rgba(0,0,0,0.4);
`;

const Sbutton = styled.button`
    padding: 0;
    background: none;
    border: none;
    cursor: pointer;
`

interface ArgProps {
    guitar:   Guitar | null;
    callback: (guitar: Guitar | null) => void;
}

const GuitarCard = ({guitar, callback}: ArgProps) => {
    const color = getColorString(guitar?.colorCd);

    return (
        <Sbutton onClick={() => callback(guitar)}>
            <ScardFrame>
                <div style={{textAlign: "center", margin: "10px", height: "40%"}}>
                    {/* モーダルにギター情報を渡す */}
                    <img style={{width:"90%", height:"90%", objectFit: "contain"}}
                         src={guitar?.src}
                         alt={guitar?.name + " " + guitar?.color}
                         />
                </div>
                <div style={{textAlign: "center", height: "60%"}}>
                    <p style={{margin: "0px 20px"}}>{guitar?.makerName}</p>
                    <h3 style={{
                        margin: "0px 20px",
                        color: color,
                        textShadow:
                            "-1px -1px 0 black, 1px -1px 0 black, -1px  1px 0 black, 1px  1px 0 black"
                        }}>
                        {guitar?.name}
                    </h3>
                    <p style={{margin: "0px 20px"}}>{guitar?.color}</p>
                    <p>{parsePrice(guitar?.price)}</p>
                </div>
            </ScardFrame>
        </Sbutton>
    );
}
export default GuitarCard;