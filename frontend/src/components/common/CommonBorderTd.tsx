import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import React from "react";

const Td = styled.td`
    border-top: 1px solid ${COLORS.BORDER};
    border-bottom: 1px solid ${COLORS.BORDER};
`;

interface ArgProps {
    children:  React.ReactNode;
}

const CommonBorderTd = ({children}: ArgProps) => {
    return (
        <Td>{children}</Td>
    );
}

export default CommonBorderTd;