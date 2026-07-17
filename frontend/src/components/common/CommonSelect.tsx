import React, { ChangeEventHandler } from "react";
import styled from "styled-components";
import { COLORS, SIZE } from "../../lib/Constants";

const Frame = styled.div`
    width: ${SIZE.INPUT_WIDTH};
    margin: 10px;
`;

const Select = styled.select`
    width: ${SIZE.INPUT_WIDTH};
    margin: 0 10px 0 10px;
    height: ${SIZE.INPUT_HEIGHT};
    border: 1px solid ${COLORS.BORDER};
    box-shadow: ${COLORS.DIALOG_SHADOW} 1px 1px;
    color: ${COLORS.MAIN_FONT};
`;

interface ArgProps {
    title?: string;
    children: |React.ReactNode[] | React.ReactNode;
    onChange?: ChangeEventHandler<HTMLSelectElement>;
    styleObj?: React.CSSProperties;
    defaultValue?: string | number;
}

const CommonSelect = ({title, children, onChange, styleObj, defaultValue}: ArgProps) => {
    return (
        <Frame>
            <label>{title}
                <Select onChange={onChange} style={styleObj} defaultValue={defaultValue}>
                    {children}
                </Select>
            </label>
        </Frame>
    );
}

export default CommonSelect;