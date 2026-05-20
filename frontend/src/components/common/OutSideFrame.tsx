import styled from "styled-components";
import { COLORS } from "../../lib/Constants";


const Sdiv = styled.div`
    margin: 20px;
    border: solid 1px ${COLORS.BORDER_COLOR};
    box-shadow: 4px 4px ${COLORS.SHADOW};
    overflow: overlay;
    // opacityは全て透過させてしまう
    background: rgba(255, 255, 255, 0.85);
    border-radius: 10px 0 10px 0
`;

interface ArgProps {
    children: React.ReactNode;
    styleObj?: React.CSSProperties;
}

const OutSideFrame = ({children, styleObj}: ArgProps) => {
    return (
        <Sdiv style={styleObj}>
            {children}
        </Sdiv>
    );
}

export default OutSideFrame;