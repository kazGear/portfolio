import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import { GUITAR } from "../../lib/Constants";

const Sdiv = styled.div`
    width: 50%;
    height: 85%;
    margin: 20px 40px 40px 40px;
    overflow-y: auto;
`;
interface ArgProps {
    selectedGuitars: Guitar | null;
}

const DetailRightSide = ({selectedGuitars}: ArgProps) => {
    const guitar = selectedGuitars;

    return (
        <Sdiv>
            <h2>Guitars spec</h2>
            <table>
                <tbody>
                    <tr>
                        <th>Maker:&emsp;</th>
                        <td>{guitar?.makerName}</td>
                    </tr>
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
                        <td>{guitar?.neckMaterial === GUITAR.UNkNOWN ? ""
                                                                     : guitar?.neckMaterialName}</td>
                    </tr>
                    <tr>
                        <th>Fingerboard:&emsp;</th>
                        <td>{guitar?.fingerboard === GUITAR.UNkNOWN ? ""
                                                                    : guitar?.fingerboardName}</td>
                    </tr>
                    <tr>
                        <th>FretCount:&emsp;</th>
                        <td>{guitar?.fretCount! <= GUITAR.INVALID_NUMBER ? ""
                                                                         : guitar?.fretCount! + " frets"}</td>
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
                        <td>{guitar?.scaleLengthMm! <= GUITAR.INVALID_NUMBER ? ""
                                                                             : guitar?.scaleLengthMm! + " mm"}</td>
                    </tr>
                    <tr>
                        <th>Weight:&emsp;</th>
                        <td>{guitar?.weight! <= GUITAR.INVALID_NUMBER ? ""
                                                                      : guitar?.weight! + " kg"}</td>
                    </tr>
                </tbody>
            </table>
        </Sdiv>
    );
}
export default DetailRightSide;