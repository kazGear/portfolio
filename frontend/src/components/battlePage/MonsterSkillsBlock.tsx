import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import { MonsterDTO } from "../../types/MonsterBattle";

const MonsterSkillsFrame = styled.div`
    border: solid 4px ${COLORS.BORDER};
    border-radius: 10px;
    padding: 5px;
    background: rgba(255, 255, 255, 0.8);
`;
const List = styled.li`
    list-style: none;

    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
`;
const Ol = styled.ol`
    padding: 0 5px 0 5px;
    margin: 0;
`;

interface ArgProps {
    monster: MonsterDTO;
 }

const MonsterSkillsBlock = ({monster}: ArgProps) => {
    return (
        <MonsterSkillsFrame>
            <Ol>
                {
                    monster.Skills !== undefined ? (
                    monster.Skills.map((skill, index) => (
                        <List key={index}>{skill.SkillName}</List>
                    ))
                    ) : (
                        <p>Loading ... </p>
                    )
                }
            </Ol>
        </MonsterSkillsFrame>
    );
}

export default MonsterSkillsBlock;