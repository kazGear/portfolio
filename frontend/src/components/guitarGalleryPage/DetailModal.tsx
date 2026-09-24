import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import { COLORS, SIZE } from "../../lib/Constants";
import CommonButton from "../common/CommonButton";
import GuitarSpec from "./GuitarSpec";
import CommonZoomableImage from "../common/CommonZoomableImage";
import { parseGuitarPrice } from "./GuitarFuncs";

const Background = styled.div`
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

const Modal = styled.div`
    width: 90%;
    height: 88%;
    display: flex;
    justify-content: space-evenly;
    border-radius: 20px;
    position: absolute;
    z-index: 2000;
    top: calc(50% + ${SIZE.HEADER_HEIGHT} / 2);
    left: 50%;
    transform: translate(-50%, -50%);
    background: ${COLORS.BASE_BACKGROUND};
    box-shadow:
        inset 0 4px 0 rgba(255,255,255,0.5),
        inset 0 -8px 40px rgba(0,0,0,0.8);
`;

const P = styled.p`
    overflow-y: auto;
    font-size: 14px;
    width: 100%;
    height: 30%;
`;

interface ArgProps {
    selectedGuitar : Guitar | null;
    isShow         : boolean;
    callback       : React.Dispatch<React.SetStateAction<boolean>>
}

const DetailModal = ({selectedGuitar, isShow, callback}: ArgProps) => {
    const isShowDetail = isShow ? "block" : "none"; // 詳細画面の表示制御
    const guitar = selectedGuitar;

    return (
        <Background style={{display: isShowDetail}}>
            <Modal>
                <GuitarSpec selectedGuitars={guitar}/>

                <div style={{width: "50%", margin: "20px"}}>
                    <p style={{marginTop: 0}}>最終更新日：{guitar?.updated}</p>

                    <CommonZoomableImage
                        imgURL={guitar?.src}
                        alt={guitar?.makerName + " | " + guitar?.name + " | " + guitar?.color}
                        width={450}
                        height={300}
                        zoomRate={300}/>

                    <h2 style={{margin: "0px"}}>
                        price:&emsp;{parseGuitarPrice(guitar?.price!)}
                    </h2>

                    <P>{guitar?.comment}</P>
                </div>

                <CommonButton
                        text="閉じる"
                        onClick={() => callback(false)}
                        styleObj={{
                            position: "absolute",
                            right: "0",
                            bottom: "0",
                            margin: "0px 20px 20px 0px"}}
                            />
            </Modal>
        </Background>
    );
}
export default DetailModal;