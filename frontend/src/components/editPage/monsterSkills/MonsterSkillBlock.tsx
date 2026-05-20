import { AllSkillDTO, EditSkillsDTO } from "../../../types/Edit";
import SkillSelectorBlock from "./SkillSelectorBlock";


interface ArgProps {
    monster: EditSkillsDTO;
    allSkills: AllSkillDTO[];
}

const MonsterSkillBlock = ({monster, allSkills}: ArgProps) => {
    return (
        <div style={{height: "40px", alignContent: "center", display: "flex"}}>
           {
                monster.SkillIds.map((skillId, index) => (
                    <SkillSelectorBlock allSkills={allSkills}
                                        mySkill={skillId}
                                        key={index}
                                        index={index}
                                        monster={monster}
                                        />
                ))
            }
        </div>
    );
}
export default MonsterSkillBlock;