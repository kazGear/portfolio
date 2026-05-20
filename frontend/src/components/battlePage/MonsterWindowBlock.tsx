import styled from "styled-components";
import { MetaDataDTO, MonsterDTO } from "../../types/MonsterBattle";
import MonsterInfoBlock from "./MonsterInfoBlock";
import MonsterImgBlock from "./MonsterImgBlock";
import MonsterSkillsBlock from "./MonsterSkillsBlock";

const SdivMonsterWindowFrame = styled.div`
    text-align: center;
    margin: 0 20px 0 20px;
    min-width: 180px;
`;

interface ArgProps {
    monster: MonsterDTO;
    shortLog: MetaDataDTO[];
}

const MonsterWindowBlock = ({monster, shortLog}: ArgProps) => {
    return (
        <SdivMonsterWindowFrame>
            <MonsterInfoBlock monster={monster} shortLog={shortLog}/>
            <MonsterImgBlock monster={monster} shortLog={shortLog}/>
            <MonsterSkillsBlock monster={monster} />
        </SdivMonsterWindowFrame>
    );
}

export default MonsterWindowBlock;