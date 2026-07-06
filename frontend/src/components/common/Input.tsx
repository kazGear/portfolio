import React, { forwardRef, useEffect, useState } from "react";
import styled from "styled-components";
import { COLORS, SIZE } from "../../lib/Constants";

const Sinput = styled.input`
    width: ${SIZE.INPUT_WIDTH};
    height: ${SIZE.INPUT_HEIGHT};
    padding: 0;
    border: ${COLORS.BORDER_COLOR} 1px solid;
`;
const Sspan = styled.span`
    display: inline;
    font-size: 9px;
    margin-left: 10px;
    color: ${COLORS.ALERT_MESSAGE_COLOR};
`;

interface ArgProps {
    labelTitle?: string;
    inputType: string;
    accept?: string;
    placeholder?: string;
    id?: string
    name?: string;
    alertMessage?: string;
    showMessage?: boolean;
    disabled?: boolean;
    styleObj?: React.CSSProperties;
    onChange?: React.ChangeEventHandler<HTMLInputElement>;
    onBlur?: React.FocusEventHandler<HTMLInputElement>;
    onClick?: React.MouseEventHandler<HTMLInputElement>;
    value?: string;
    defaultValue?: string | number;
    min?: string;
    max?: string;
}

// forwardRef: 親から参照されるため
const Input = forwardRef<HTMLInputElement, ArgProps>(({
    labelTitle,
    inputType,
    accept,
    placeholder,
    id,
    name,
    alertMessage,
    showMessage = false,
    disabled = false,
    styleObj = {},
    onChange,
    onBlur,
    onClick,
    value,
    defaultValue,
    min,
    max,
}, ref) => {
    const [show, setShow] = useState(false);

    useEffect(() => {
        setShow(!showMessage);
    }, [showMessage]);

    return (
        <>
            <label style={{marginRight: "10px"}}>
                {labelTitle}
                <Sinput type={inputType}
                        accept={accept}
                        style={styleObj}
                        id={id}
                        name={name}
                        placeholder={placeholder}
                        onChange={onChange}
                        onBlur={onBlur}
                        onClick={onClick}
                        disabled={disabled}
                        ref={ref}
                        value={value}
                        defaultValue={defaultValue}
                        min={min}
                        max={max}
                        onKeyDown={(e) => {
                            // Enter入力でフォーカスアウト = 入力済扱い
                            if (e.key === "Enter") {
                                e.currentTarget.blur();
                            }
                        }}
                        />
            </label>
            {
                show ? <Sspan>{alertMessage}</Sspan> : ""
            }
        </>
    );
});

export default Input;