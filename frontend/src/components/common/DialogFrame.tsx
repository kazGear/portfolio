import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const SdivDialogFrame = styled.div`
    background-color: ${COLORS.BASE_BACKGROUND};
    background-size: cover;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    min-width: 50%;
    max-width: 70%;
    min-height: 20%;
    max-height: 70%;
    border-radius: 50px;
    z-index: 3000;
`;

const SdivMessageArea = styled.div`
    margin: 20px auto;
    width: 80%;
    height: 80%;
    z-index: 1000;
`;
const SimgL = styled.img`
    width: 100px;
    height: 100px;
    position: absolute;
    left: 0;
    bottom: 0;
    z-index: 2000;
`;
const SimgR = styled.img`
    width: 150px;
    height: 170px;
    position: absolute;
    right: 0;
    top: 0;
    z-index: 2000;
`;
const SdivFilter = styled.div`
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.8);
    position: absolute;
    top: 0;
    left: 0;
    z-index: 500;
`;

interface DialogFrameProps {
    children: React.ReactNode;
    showDialog: boolean;
    showFilter?: boolean;
}

const DialogFrame = (
    {children, showDialog = true, showFilter}: DialogFrameProps
) => {
    return (
        <>
            <SdivDialogFrame style={{display: showDialog ? "block" : "none"}}>
                {/* <SimgL src={imageL} alt="ツタ"></SimgL> */}
                <SdivMessageArea>
                    {children}
                </SdivMessageArea>
                {/* <SimgR src={imageR} alt="ツタ"></SimgR> */}
            </SdivDialogFrame>

            <SdivFilter style={{display: showDialog ? "block" : "none"}}/>
        </>
    );
}
export default DialogFrame;