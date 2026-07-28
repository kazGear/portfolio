import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import { GUITAR } from "../../lib/Constants";

const Div = styled.div`
    width: 50%;
    height: 85%;
    margin: 20px 40px 40px 40px;
    overflow-y: auto;
`;

const Tr = styled.tr`
    border-top: gray 1px solid;
    border-bottom: gray 1px solid;
`;

const Th = styled.th`
    text-align: left;
`;

interface ArgProps {
    selectedGuitars: Guitar | null;
}

const DetailRightSide = ({selectedGuitars}: ArgProps) => {
    const guitar = selectedGuitars;

    return (
        <Div>
            <h2>Guitars spec</h2>
            <table>
                <tbody>
                    <Tr>
                        <Th>Maker:&emsp;</Th>
                        <td>{guitar?.makerName}</td>
                    </Tr>
                    <Tr>
                        <Th>Name:&emsp;</Th>
                        <td>{guitar?.name}</td>
                    </Tr>
                    <Tr>
                        <Th>Color:&emsp;</Th>
                        <td>{guitar?.color}</td>
                    </Tr>
                    <Tr>
                        <Th>Series:&emsp;</Th>
                        <td>{guitar?.series}</td>
                    </Tr>
                    <Tr>
                        <Th>BodyMaterial:&emsp;</Th>
                        <td>{guitar?.bodyMaterial}</td>
                    </Tr>
                    <Tr>
                        <Th>BodyFinish:&emsp;</Th>
                        <td>{guitar?.bodyFinish}</td>
                    </Tr>
                    <Tr>
                        <Th>NeckMaterial:&emsp;</Th>
                        <td>{guitar?.neckMaterial === GUITAR.UNkNOWN ? ""
                                                                     : guitar?.neckMaterialName}
                        </td>
                    </Tr>
                    <Tr>
                        <Th>Fingerboard:&emsp;</Th>
                        <td>{guitar?.fingerboard === GUITAR.UNkNOWN ? ""
                                                                    : guitar?.fingerboardName}
                        </td>
                    </Tr>
                    <Tr>
                        <Th>FretCount:&emsp;</Th>
                        <td>{guitar?.fretCount! <= GUITAR.INVALID_NUMBER ? ""
                                                                         : guitar?.fretCount! + " frets"}
                        </td>
                    </Tr>
                    <Tr>
                        <Th>Pickups:&emsp;</Th>
                        <td>{guitar?.pickups}</td>
                    </Tr>
                    <Tr>
                        <Th>Bridge:&emsp;</Th>
                        <td>{guitar?.bridge}</td>
                    </Tr>
                    <Tr>
                        <Th>Controls:&emsp;</Th>
                        <td>{guitar?.controls}</td>
                    </Tr>
                    <Tr>
                        <Th>Inlays:&emsp;</Th>
                        <td>{guitar?.inlays}</td>
                    </Tr>
                    <Tr>
                        <Th>Joint:&emsp;</Th>
                        <td>{guitar?.joint}</td>
                    </Tr>
                    <Tr>
                        <Th>ScaleLength:&emsp;</Th>
                        <td>{guitar?.scaleLengthMm! <= GUITAR.INVALID_NUMBER ? ""
                                                                             : guitar?.scaleLengthMm! + " mm"}
                        </td>
                    </Tr>
                    <Tr>
                        <Th>Weight:&emsp;</Th>
                        <td>{guitar?.weight! <= GUITAR.INVALID_NUMBER ? ""
                                                                      : guitar?.weight! + " kg"}
                        </td>
                    </Tr>
                </tbody>
            </table>
        </Div>
    );
}
export default DetailRightSide;