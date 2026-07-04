import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import { GUITAR } from "../../lib/Constants";

const Sdiv = styled.div`
    width: 50%;
    height: 85%;
    margin: 20px 40px 40px 40px;
    overflow-y: auto;
`;

const Str = styled.tr`
    border-top: gray 1px solid;
    border-bottom: gray 1px solid;
`;

const Sth = styled.th`
    text-align: left;
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
                    <Str>
                        <Sth>Maker:&emsp;</Sth>
                        <td>{guitar?.makerName}</td>
                    </Str>
                    <Str>
                        <Sth>Name:&emsp;</Sth>
                        <td>{guitar?.name}</td>
                    </Str>
                    <Str>
                        <Sth>Color:&emsp;</Sth>
                        <td>{guitar?.color}</td>
                    </Str>
                    <Str>
                        <Sth>Series:&emsp;</Sth>
                        <td>{guitar?.series}</td>
                    </Str>
                    <Str>
                        <Sth>BodyMaterial:&emsp;</Sth>
                        <td>{guitar?.bodyMaterial}</td>
                    </Str>
                    <Str>
                        <Sth>BodyFinish:&emsp;</Sth>
                        <td>{guitar?.bodyFinish}</td>
                    </Str>
                    <Str>
                        <Sth>NeckMaterial:&emsp;</Sth>
                        <td>{guitar?.neckMaterial === GUITAR.UNkNOWN ? ""
                                                                     : guitar?.neckMaterialName}
                        </td>
                    </Str>
                    <Str>
                        <Sth>Fingerboard:&emsp;</Sth>
                        <td>{guitar?.fingerboard === GUITAR.UNkNOWN ? ""
                                                                    : guitar?.fingerboardName}
                        </td>
                    </Str>
                    <Str>
                        <Sth>FretCount:&emsp;</Sth>
                        <td>{guitar?.fretCount! <= GUITAR.INVALID_NUMBER ? ""
                                                                         : guitar?.fretCount! + " frets"}
                        </td>
                    </Str>
                    <Str>
                        <Sth>Pickups:&emsp;</Sth>
                        <td>{guitar?.pickups}</td>
                    </Str>
                    <Str>
                        <Sth>Bridge:&emsp;</Sth>
                        <td>{guitar?.bridge}</td>
                    </Str>
                    <Str>
                        <Sth>Controls:&emsp;</Sth>
                        <td>{guitar?.controls}</td>
                    </Str>
                    <Str>
                        <Sth>Inlays:&emsp;</Sth>
                        <td>{guitar?.inlays}</td>
                    </Str>
                    <Str>
                        <Sth>Joint:&emsp;</Sth>
                        <td>{guitar?.joint}</td>
                    </Str>
                    <Str>
                        <Sth>ScaleLength:&emsp;</Sth>
                        <td>{guitar?.scaleLengthMm! <= GUITAR.INVALID_NUMBER ? ""
                                                                             : guitar?.scaleLengthMm! + " mm"}
                        </td>
                    </Str>
                    <Str>
                        <Sth>Weight:&emsp;</Sth>
                        <td>{guitar?.weight! <= GUITAR.INVALID_NUMBER ? ""
                                                                      : guitar?.weight! + " kg"}
                        </td>
                    </Str>
                </tbody>
            </table>
        </Sdiv>
    );
}
export default DetailRightSide;