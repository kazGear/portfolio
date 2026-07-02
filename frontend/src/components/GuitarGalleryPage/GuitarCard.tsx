import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import { Guitar } from "../../types/Guitar";

const ScardFrame = styled.span`
    font-weight: 900;
    width: 220px;
    height: 250px;
    margin: 10px 10px 20px 5px;
    background: ${COLORS.BASE_BACKGROUND};
    border-radius: 15px;
    box-shadow:
        inset 0 1px 0 rgba(255,255,255,0.1),
        inset 0 -2px 10px rgba(0,0,0,0.4);
`;

interface ArgProps {
    guitar: Guitar | null;
}

const GuitarCard = ({guitar}: ArgProps) => {
    return (
        <ScardFrame>
            <div style={{textAlign: "center", margin: "10px"}}>
                <img style={{width:"200px", height:"100px"}}
                        src={guitar?.src}
                        alt={guitar?.name + " " + guitar?.color}
                        />
            </div>
            <div style={{textAlign: "center"}}>
                <h3 style={{marginBottom: "0", color: COLORS.ACCENT_FONT_PINK}}>
                    {guitar?.name}
                </h3>
                <p style={{marginTop: "0"}}>{guitar?.color}</p>
                <p>{guitar?.price} 円</p>
            </div>
        </ScardFrame>
    );
}
export default GuitarCard;