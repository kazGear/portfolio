import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import { COLORS } from "../../lib/Constants";
import Button from "../common/Button";

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
                <div style={{width: "50%", margin: "0px 40px"}}>
                    <img src={guitar?.src} alt={guitar?.name} style={{width:"100%", height:"40%", objectFit: "contain", marginTop: "20px"}}/>
                    <h2 style={{margin: "0px"}}>price:&emsp;{guitar?.price} 円</h2>
                    <p>comment.</p>
                    <p style={{overflowY: "scroll", fontSize: "14px", height: "30%"}}>{guitar?.comment}</p>
                </div>
                <div style={{width: "50%", margin: "0px 40px"}}>
                    <h2>Guitars spec</h2>
                    <table style={{overflowY: "scroll"}}>
                        <tbody>
                            <tr>
                                <th>Name:&emsp;</th>
                                <td>{guitar?.name}</td>
                            </tr>
                            <tr>
                                <th>Color:&emsp;</th>
                                <td>{guitar?.color}</td>
                            </tr>
                            <tr>
                                <th>Series:&emsp;</th>
                                <td>{guitar?.series}</td>
                            </tr>
                            <tr>
                                <th>BodyMaterial:&emsp;</th>
                                <td>{guitar?.bodyMaterial}</td>
                            </tr>
                            <tr>
                                <th>BodyFinish:&emsp;</th>
                                <td>{guitar?.bodyFinish}</td>
                            </tr>
                            <tr>
                                <th>NeckMaterial:&emsp;</th>
                                <td>{guitar?.neckMaterial}</td>
                            </tr>
                            <tr>
                                <th>Fingerboard:&emsp;</th>
                                <td>{guitar?.fingerboard}</td>
                            </tr>
                            <tr>
                                <th>FretCount:&emsp;</th>
                                <td>{guitar?.fretCount} frets</td>
                            </tr>
                            <tr>
                                <th>Pickups:&emsp;</th>
                                <td>{guitar?.pickups}</td>
                            </tr>
                            <tr>
                                <th>Bridge:&emsp;</th>
                                <td>{guitar?.bridge}</td>
                            </tr>
                            <tr>
                                <th>Controls:&emsp;</th>
                                <td>{guitar?.controls}</td>
                            </tr>
                            <tr>
                                <th>Inlays:&emsp;</th>
                                <td>{guitar?.inlays}</td>
                            </tr>
                            <tr>
                                <th>Joint:&emsp;</th>
                                <td>{guitar?.joint}</td>
                            </tr>
                            <tr>
                                <th>ScaleLength:&emsp;</th>
                                <td>{guitar?.scaleLengthMm} mm</td>
                            </tr>
                            <tr>
                                <th>Weight:&emsp;</th>
                                <td>{guitar?.weight} mm</td>
                            </tr>

                        </tbody>
                    </table>
                </div>
                <Button text="閉じる" onClick={() => callback(false)} styleObj={{position: "absolute", right: "0", bottom: "0", margin: "0px 40px 40px 0px"}}/>
            </Smodal>
        </Sbackground>
    );
}
export default DetailModal;