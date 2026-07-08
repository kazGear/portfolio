import styled from "styled-components";
import monsterImages from "../../../lib/MonsterImages";
import BorderTd from "../../common/CommonBorderTd";
import { CodeDTO } from "../../../types/Common";
import { EditMonsterDTO } from "../../../types/Edit";
import EditMonsterName from "./EditMonsterNameBlock";
import EditMonsterHp from "./EditMonsterHpBlock";
import EditMonsterAttack from "./EditMonsterAttackBlock";
import EditMonsterSpeed from "./EditMonsterSpeedBlock";
import EditMonsterWeek from "./EditMonsterWeekBlock";
import { useState } from "react";
import CommonNowLoading from "../../common/CommonNowLoading";
import CommonStrong from "../../common/CommonStrong";

const Img = styled.img`
    vertical-align: middle;
    width: 50px;
    height: 50px;
    margin: 5px;
`;

interface ArgProps {
    editMonsters: EditMonsterDTO[];
    isNowLoading: boolean;
}

const MonsterTableBody = ({editMonsters, isNowLoading}: ArgProps) => {
    const [weekDropDown, setWeekDropDown] = useState<CodeDTO[]>([]);

    /**
     * ローディング
     */
    if (isNowLoading) {
        return (
            <div style={{margin: "100px"}}>
                <CommonNowLoading alt="ローディング" />
            </div>
        );
    }

    return (
        <table style={{width: "100%"}}>
            <thead>
                <tr style={{textAlign: "center"}}>
                    <th><CommonStrong>ID</CommonStrong></th>
                    <td><CommonStrong>イメージ</CommonStrong></td>
                    <td><CommonStrong>モンスター名</CommonStrong></td>
                    <td><CommonStrong>HP</CommonStrong></td>
                    <td><CommonStrong>攻撃力</CommonStrong></td>
                    <td><CommonStrong>速さ</CommonStrong></td>
                    <td><CommonStrong>弱点</CommonStrong></td>
                    <td>⇒変更後</td>
                </tr>
            </thead>
            <tbody>
            {
                editMonsters.map((monster, index) => (
                    <tr key={index + monster.MonsterName} style={{textAlign: "center"}}>
                        {/* ID */}
                        <BorderTd>{monster.MonsterId}</BorderTd>
                        {/* イメージ */}
                        <BorderTd>
                            <Img src={monsterImages(monster.MonsterId)}
                                 alt={monster.MonsterName} />
                        </BorderTd>
                        {/* 名称 */}
                        <EditMonsterName monster={monster}/>
                        {/* HP */}
                        <EditMonsterHp monster={monster}/>
                        {/* 攻撃力 */}
                        <EditMonsterAttack monster={monster}/>
                        {/* 速さ */}
                        <EditMonsterSpeed monster={monster}/>
                        {/* 弱点 */}
                        <EditMonsterWeek weekDropDown={weekDropDown} monster={monster} setWeekDropDown={setWeekDropDown}/>
                    </tr>
                ))
            }
            </tbody>
        </table>
    );
}

export default MonsterTableBody;