import styled from "styled-components";
import { COLORS, DAMAGE_VIEW, STATE_NAME } from "../../lib/Constants";
import { useEffect, useState } from "react";
import { MetaDataDTO, MonsterDTO } from "../../types/MonsterBattle";

const SdivMonsterInfoFrame = styled.div`
    border: solid 4px ${COLORS.BORDER_COLOR};
    margin-bottom: 10px;
    border-radius: 10px;
    height: 78px;
    font-size: 15px;
    color: ${props => props.color};
    background: rgba(255, 255, 255, 0.8);
`;
const Stable = styled.table`
    width: 100%;
    height: 100%;
`;
const StdStatus = styled.td`
    white-space: nowrap;
    max-width: 105px;
    text-overflow: ellipsis;
    overflow: hidden;
`;

interface ArgProps {
    monster: MonsterDTO;
    shortLog: MetaDataDTO[];
}

const MonsterInfoBlock = ({monster, shortLog}: ArgProps) => {
    const [hp, setHp] = useState(monster.Hp);
    const [status, setStatus] = useState([""]);

    useEffect(() => {
        for (const log of shortLog) {
            // HP更新
            if (   log.TargetMonsterId === monster.MonsterId
                && log.ImpactPoint !== 0
            ) {
                setTimeout(() => {
                    setHp(log.BeforeHp - log.ImpactPoint)
                }, DAMAGE_VIEW.DAMAGE_END);
            }
            // 状態異常付与
            if (   log.TargetMonsterId === monster.MonsterId
                && log.EnableState
            ) {
                status.push(log.StateName)
                setStatus([...status]);
            }
            // 状態異常解除
            if (   log.TargetMonsterId === monster.MonsterId
                && log.DisableState
            ) {
                const filterd: string[] = status.filter(e => e !== log.StateName);
                setStatus([...filterd]);
            }
        }
        // 戦闘不能
        if (hp <= 0) setStatus([STATE_NAME.LOSER]);
    }, [shortLog]);

    // 状態表示
    const infoColor = hp <= 0 ? STATE_NAME.LOSER : "";

    return (
        <SdivMonsterInfoFrame
            color={infoColor === STATE_NAME.LOSER ? COLORS.LOSER_FONT_COLOR
                                                  : COLORS.MAIN_FONT_COLOR}>
            {
                monster.Hp !== undefined ? (
                <Stable>
                    <thead>
                        <tr>
                            <td colSpan={2}>{monster.MonsterName}</td>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>HP : </td><td>{hp <= 0 ? 0 : hp} / {monster.MaxHp}</td>
                        </tr>
                        <tr>
                            <td>状態 : </td><StdStatus>{status.join(" ")}</StdStatus>
                        </tr>
                    </tbody>
                </Stable>
                ) : (
                    <p>Loading ...</p>
                )
            }
        </SdivMonsterInfoFrame>
    );
}

export default MonsterInfoBlock;