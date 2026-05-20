import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import Button from "./Button";

const SdivEditHeader = styled.div`
    display: flex;
    position: sticky;
    justify-content: space-between;
    top: 0;
    background: ${COLORS.BASE_BACKGROUND};
`;

interface ArgProps {
    message?: string;
    buttonText: string;
    buttonWidth?: number;
    callback: () => void;
}

const HeaderOfBody = ({message, buttonText, buttonWidth, callback}: ArgProps) => {
    const width = buttonWidth ?? 140;

    return (
        <SdivEditHeader>
            <h2 style={{marginLeft: "20px"}}>
                {message}
            </h2>
            <Button text={buttonText}
                    width={width}
                    onClick={callback}
                    styleObj={{margin: "20px"}}/>
        </SdivEditHeader>
    );
}

export default HeaderOfBody;