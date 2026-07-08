import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const DialogFrame = styled.div`
    background-color: ${COLORS.BASE_BACKGROUND};
    background-size: cover;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    min-width: 60%;
    max-width: 70%;
    min-height: 20%;
    max-height: 70%;
    border-radius: 10px;
    z-index: 3000;
`;

const MessageFrame = styled.div`
    margin: 20px auto;
    width: 80%;
    height: 80%;
    z-index: 1000;
`;
const BackFilter = styled.div`
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.8);
    position: absolute;
    top: 0;
    left: 0;
    z-index: 500;
`;

interface ArgProps {
    children: React.ReactNode;
    showDialog: boolean;
    showFilter?: boolean;
}

const CommonDialogFrame = (
    {children, showDialog = true, showFilter}: ArgProps
) => {
    return (
        <>
            <DialogFrame style={{display: showDialog ? "block" : "none"}}>
                <MessageFrame>
                    {children}
                </MessageFrame>
            </DialogFrame>

            <BackFilter style={{display: showDialog ? "block" : "none"}}/>
        </>
    );
}
export default CommonDialogFrame;