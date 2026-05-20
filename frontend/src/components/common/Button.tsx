import styled from "styled-components";
import { COLORS, SIZE } from "../../lib/Constants";
import { MouseEventHandler } from "react";

interface SbuttonProps {
    opacity: number;
    width: string;
    display: string;
}

const Sbutton = styled.button<SbuttonProps>`
    background: ${COLORS.BUTTON_COLOR};
    color: ${COLORS.BUTTON_FONT_COLOR};
    font-weight: 900;
    font-size: 14px;
    border: none;
    box-shadow: ${COLORS.SHADOW} 2px 2px;
    width: ${props => props.width};
    height: ${SIZE.INPUT_HEIGHT};
    margin: 0 5px 0 5px;
    border-radius: 10px;
    opacity: ${props => props.opacity};
    display: ${props => props.display};

    &:hover {
        cursor: pointer;
        color: ${COLORS.ACCENT_FONT_PINK};
    }

    &:active {
        transform: translate(2px, 2px);
        box-shadow: none;
    }
`;

interface ButtonProps {
    id?: string;
    text: string;
    width?: number;
    onClick: MouseEventHandler<HTMLButtonElement>;
    disabled?: boolean;
    display?: string;
    styleObj?: React.CSSProperties;
}

const Button = ({
    id,
    text,
    width,
    onClick,
    disabled = false,
    display = "inline",
    styleObj}: ButtonProps
) => {
    const range: string = width ? width + "px" : "100px"
    const opacity = disabled ? COLORS.BUTTON_DISABLED : 1.0;

    return (
        <Sbutton
            id={id}
            type="button"
            width={range}
            onClick={onClick}
            opacity={opacity}
            disabled={disabled}
            display={display}
            style={styleObj}
        >
            {text}
        </Sbutton>
    );
}

export default Button;