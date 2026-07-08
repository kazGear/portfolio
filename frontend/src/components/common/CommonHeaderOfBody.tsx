import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import CommonButton from "./CommonButton";

const EditHeader = styled.div`
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

const CommonHeaderOfBody = ({message, buttonText, buttonWidth, callback}: ArgProps) => {
    const width = buttonWidth ?? 140;

    return (
        <EditHeader>
            <h2 style={{marginLeft: "20px"}}>
                {message}
            </h2>
            <CommonButton text={buttonText}
                    width={width}
                    onClick={callback}
                    styleObj={{margin: "20px"}}/>
        </EditHeader>
    );
}

export default CommonHeaderOfBody;