import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const MenuTitle = styled.h2`
    height: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 10px;
    border-collapse: collapse;
    background: ${COLORS.LOGINED};
    box-shadow: 4px 4px ${COLORS.SHADOW};
`;

interface ArgProps {
    title: string
    styleObj?: React.CSSProperties;
    className?: string
}

const CommonMenuTitle = ({title, styleObj, className}: ArgProps) => {
    return <MenuTitle className={className} style={styleObj}>
               {title}
           </MenuTitle>;
};

export default CommonMenuTitle;
