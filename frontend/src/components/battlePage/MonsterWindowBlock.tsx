import styled from "styled-components";
import { MetaDataDTO, MonsterDTO } from "../../types/MonsterBattle";
import MonsterInfoBlock from "./MonsterInfoBlock";
import MonsterImgBlock from "./MonsterImgBlock";
import MonsterSkillsBlock from "./MonsterSkillsBlock";

const MonsterWindowFrame = styled.div`
    text-align: center;
    margin: 0 5px 0 5px;
    min-width: 155px;
`;

interface ArgProps {
    monster:  MonsterDTO;
    shortLog: MetaDataDTO[];
}

const MonsterWindowBlock = ({monster, shortLog}: ArgProps) => {
    return (
        <MonsterWindowFrame>
            <MonsterInfoBlock monster={monster} shortLog={shortLog}/>
            <MonsterImgBlock monster={monster} shortLog={shortLog}/>
            <MonsterSkillsBlock monster={monster} />
        </MonsterWindowFrame>
    );
}

export default MonsterWindowBlock;