import styled from "styled-components";
import { COLORS, SIZE } from "../../lib/Constants";
import React from "react";

const Div = styled.div`
    margin: 0px;
    padding: 0px;
    width: ${SIZE.INPUT_WIDTH};
    height: ${SIZE.INPUT_HEIGHT};
    text-align: center;
`;
type Props = {
    enable: boolean;
};
const Button = styled.button<Props>`
    color: ${COLORS.ACCENT_FONT_GREEN};
    border: none;
    background: none;
    pointer-events: ${props => props.enable ? "auto" : "none"};
    opacity: ${props => props.enable ? 1.0 : 0.2};
    cursor: ${props => props.enable ? "pointer" : "not-allowed"};
`;
const Span = styled.span`
    font-size: 16px;
    text-shadow:
        -1px -1px 0 black,
        1px -1px 0 black,
        -1px  1px 0 black,
        1px  1px 0 black;
`;

interface ArgProps {
    children:              React.ReactNode;
    styleObj?:             React.CSSProperties;
    hasPrev:               boolean;
    hasNext:               boolean;
    changePrevPageHandler: React.MouseEventHandler<HTMLButtonElement> | undefined;
    changeNextPageHandler: React.MouseEventHandler<HTMLButtonElement> | undefined;
}

const CommonPagination = ({children,
                           styleObj,
                           hasPrev,
                           hasNext,
                           changePrevPageHandler,
                           changeNextPageHandler}: ArgProps
) => {
    return (
        <Div style={styleObj}>
            <Button onClick={changePrevPageHandler} enable={hasPrev}><Span> ◀ </Span></Button>
                {/* childrenはインライン要素が望ましい */}
                {children}
            <Button onClick={changeNextPageHandler} enable={hasNext}><Span> ▶ </Span></Button>
        </Div>
    );
}

export default CommonPagination;