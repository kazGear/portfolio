import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import React from "react";

const Std = styled.td`
    border-top: 1px solid ${COLORS.BORDER_COLOR};
    border-bottom: 1px solid ${COLORS.BORDER_COLOR};
`;

interface ArgProps {
    children: React.ReactNode;
    styleObj?: React.CSSProperties;
}

const BorderTd = ({children, styleObj}: ArgProps) => {
    return (
        <Std>{children}</Std>
    );
}

export default BorderTd;