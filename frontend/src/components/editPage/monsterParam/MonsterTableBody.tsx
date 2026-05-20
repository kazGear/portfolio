import styled from "styled-components";
import monsterImages from "../../../lib/MonsterImages";
import BorderTd from "../../common/BorderTd";
import { CodeDTO } from "../../../types/Common";
import { EditMonsterDTO } from "../../../types/Edit";
import EditMonsterName from "./EditMonsterNameBlock";
import EditMonsterHp from "./EditMonsterHpBlock";
import EditMonsterAttack from "./EditMonsterAttackBlock";
import EditMonsterSpeed from "./EditMonsterSpeedBlock";
import EditMonsterWeek from "./EditMonsterWeekBlock";
import { useState } from "react";
import NowLoading from "../../common/NowLoading";

const Simg = styled.img`
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
                <NowLoading alt="ローディング" />
            </div>
        );
    }

    return (
        <tbody>
        {
            editMonsters.map((monster, index) => (
                <tr key={index}>
                    {/* ID */}
                    <BorderTd>{monster.MonsterId}</BorderTd>
                    {/* イメージ */}
                    <BorderTd>
                        <Simg src={monsterImages(monster.MonsterId)}
                             alt="モンスター" />
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
    );
}

export default MonsterTableBody;