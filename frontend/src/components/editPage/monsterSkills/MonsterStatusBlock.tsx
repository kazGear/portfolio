import styled from "styled-components";
import monsterImages from "../../../lib/MonsterImages";
import { EditSkillsDTO } from "../../../types/Edit";
import CommonStrong from "../../common/CommonStrong";

const H4 = styled.h4`
    margin: 20px;
`;
const Img = styled.img`
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
            <H4>{monster.MonsterId}</H4>
            <Img src={monsterImages(monster.MonsterId)} alt={monster.MonsterName} />
            <H4><CommonStrong>{monster.MonsterName}</CommonStrong></H4>
            <H4>HP：{monster.Hp}</H4>
            <H4>攻撃力：{monster.MonsterAttack}</H4>
            <H4>速さ：{monster.Speed}</H4>
            <H4>弱点：{monster.WeekName}</H4>
        </div>
    );
}

export default MonsterStatusBlock;