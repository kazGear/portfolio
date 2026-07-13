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
    cursor: pointer;
    background: none;
    pointer-events: ${Props => Props.enable ? "" : "none"};
    opacity: ${Props => Props.enable ? 1.0 : 0.2};
    cursor: ${Props => Props.enable ? "" : "not-allowed"};
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
    children:          React.ReactNode;
    styleObj?:         React.CSSProperties;
    hasPrev:           boolean;
    hasNext:           boolean;
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
                {/* インライン要素が望ましい */}
                {children}
            <Button onClick={changeNextPageHandler} enable={hasNext}><Span> ▶ </Span></Button>
        </Div>
    );
}

export default CommonPagination;