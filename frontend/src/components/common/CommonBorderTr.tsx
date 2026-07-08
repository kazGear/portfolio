import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import React from "react";

const Tr = styled.tr`
    border-top: 1px solid ${COLORS.BORDER_COLOR};
    border-bottom: 1px solid ${COLORS.BORDER_COLOR};
`;

interface ArgProps {
    children:  React.ReactNode;
    styleObj?: React.CSSProperties;
}

const CommonBorderTr = ({children, styleObj}: ArgProps) => {
    return (
        <Tr style={styleObj}>{children}</Tr>
    );
}

export default CommonBorderTr;