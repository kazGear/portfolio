import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import { MonsterDTO } from "../../types/MonsterBattle";

const SdivMonsterSkillsFrame = styled.div`
    border: solid 4px ${COLORS.BORDER_COLOR};
    border-radius: 10px;
    padding: 5px;
    background: rgba(255, 255, 255, 0.8);
`;
const Slist = styled.li`
    list-style: none;
`;
const Sol = styled.ol`
    padding: 0 5px 0 5px;
    margin: 0;
`;

interface ArgProps {
    monster: MonsterDTO;
 }

const MonsterSkillsBlock = ({monster}: ArgProps) => {
    return (
        <SdivMonsterSkillsFrame>
            <Sol>
                {
                    monster.Skills !== undefined ? (
                    monster.Skills.map((skill, index) => (
                        <Slist key={index}>{skill.SkillName}</Slist>
                    ))
                    ) : (
                        <p>Loading ... </p>
                    )
                }
            </Sol>
        </SdivMonsterSkillsFrame>
    );
}

export default MonsterSkillsBlock;