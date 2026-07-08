import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import React from "react";

const Td = styled.td`
    border-top: 1px solid ${COLORS.BORDER_COLOR};
    border-bottom: 1px solid ${COLORS.BORDER_COLOR};
`;

interface ArgProps {
    children:  React.ReactNode;
    styleObj?: React.CSSProperties;
}

const CommonBorderTd = ({children, styleObj}: ArgProps) => {
    return (
        <Td>{children}</Td>
    );
}

export default CommonBorderTd;