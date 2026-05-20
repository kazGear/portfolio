import React, { forwardRef, useEffect, useState } from "react";
import styled from "styled-components";
import { COLORS, SIZE } from "../../lib/Constants";


const Sinput = styled.input`
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
    labelTitle: string;
    inputType: string;
    accept?: string;
    placeholder?: string;
    id?: string
    name?: string;
    alertMessage?: string;
    showMessage?: boolean;
    disabled?: boolean;
    styleObj?: React.CSSProperties;
    onChange?: React.ChangeEventHandler<HTMLInputElement> | undefined;
    onClick?: React.MouseEventHandler<HTMLInputElement> | undefined;
    value?: string;
    defaultValue?: string | number;
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
    onClick,
    value,
    defaultValue
}, ref) => {
    const [show, setShow] = useState(false);

    useEffect(() => {
        setShow(!showMessage);
    }, [showMessage]);

    return (
        <div style={{margin: "5px 0 5px 0"}}>
            <label style={{marginRight: "10px"}}>
                {labelTitle}
            </label>
            <Sinput type={inputType}
                    accept={accept}
                    style={styleObj}
                    id={id}
                    name={name}
                    placeholder={placeholder}
                    onChange={onChange}
                    onClick={onClick}
                    disabled={disabled}
                    ref={ref}
                    value={value}
                    defaultValue={defaultValue}
                    />
            {
                show ? <Sspan>{alertMessage}</Sspan> : ""
            }
        </div>
    );
});

export default Input;