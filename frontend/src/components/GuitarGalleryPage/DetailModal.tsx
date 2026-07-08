import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import { COLORS } from "../../lib/Constants";
import CommonButton from "../common/CommonButton";
import DetailLeftSide from "./DetailLeftSide";
import DetailRightSide from "./DetailRightSide";

const Sbackground = styled.div`
    width: 100%;
    height: 100%;
    display: none;
    position: absolute;
    z-index: 1000;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: rgba(0, 0, 0, 0.7);
`;

const Smodal = styled.div`
    width: 80%;
    height: 80%;
    display: flex;
    justify-content: space-evenly;
    border-radius: 20px;
    position: absolute;
    z-index: 2000;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: ${COLORS.BASE_BACKGROUND};
    box-shadow:
        inset 0 4px 0 rgba(255,255,255,0.5),
        inset 0 -8px 40px rgba(0,0,0,0.8);
`;

interface ArgProps {
    selectedGuitars: Guitar | null;
    isShow         : boolean;
    callback       : React.Dispatch<React.SetStateAction<boolean>>
}

const DetailModal = ({selectedGuitars, isShow, callback}: ArgProps) => {
    const isShowDetail = isShow ? "block" : "none"; // 詳細画面の表示制御
    const guitar = selectedGuitars;

    return (
        <Sbackground style={{display: isShowDetail}}>
            <Smodal>
                <DetailLeftSide selectedGuitars={guitar}/>
                <DetailRightSide selectedGuitars={guitar}/>
                <CommonButton text="閉じる"
                        onClick={() => callback(false)}
                        styleObj={{
                            position: "absolute",
                            right: "0",
                            bottom: "0",
                            margin: "0px 40px 40px 0px"}}
                            />
            </Smodal>
        </Sbackground>
    );
}
export default DetailModal;