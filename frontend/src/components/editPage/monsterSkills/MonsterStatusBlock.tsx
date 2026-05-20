import styled from "styled-components";
import monsterImages from "../../../lib/MonsterImages";
import { EditSkillsDTO } from "../../../types/Edit";
import Strong from "../../common/Strong";

const Sh4 = styled.h4`
    margin: 20px;
`;
const Simg = styled.img`
    width: 50px;
    height: 50px;
    margin: 10px;
`;

interface ArgProps {
    monster: EditSkillsDTO;
}

const MonsterStatusBlock = ({monster}: ArgProps) => {
    return (
        <div style={{display: "flex", alignItems: "center", height: "60px"}}>
            <Sh4>{monster.MonsterId}</Sh4>
            <Simg src={monsterImages(monster.MonsterId)} alt="モンスター" />
            <Sh4><Strong>{monster.MonsterName}</Strong></Sh4>
            <Sh4>HP：{monster.Hp}</Sh4>
            <Sh4>攻撃力：{monster.MonsterAttack}</Sh4>
            <Sh4>速さ：{monster.Speed}</Sh4>
            <Sh4>弱点：{monster.WeekName}</Sh4>
        </div>
    );
}

export default MonsterStatusBlock;