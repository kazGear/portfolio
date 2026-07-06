import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import React from "react";

const Str = styled.tr`
    border-top: 1px solid ${COLORS.BORDER_COLOR};
    border-bottom: 1px solid ${COLORS.BORDER_COLOR};
`;

interface ArgProps {
    children:  React.ReactNode;
    styleObj?: React.CSSProperties;
}

const BorderTr = ({children, styleObj}: ArgProps) => {
    return (
        <Str style={styleObj}>{children}</Str>
    );
}

export default BorderTr;