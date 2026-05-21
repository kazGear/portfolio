import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const SmenuTitle = styled.h2`
    height: 40px;
    display: flex;
    align-items: center;
    padding-left: 60px;
    border-radius: 10px;
    border-collapse: collapse;
    background: ${COLORS.LOGINED_COLOR};
    box-shadow: 4px 4px ${COLORS.SHADOW};
`;

interface ArgProps {
    title: string
    styleObj?: React.CSSProperties;
    className?: string
}

const MenuTitle = ({title, styleObj, className}: ArgProps) => {
    return <SmenuTitle className={className} style={styleObj}>
               {title}
           </SmenuTitle>;
};

export default MenuTitle;
