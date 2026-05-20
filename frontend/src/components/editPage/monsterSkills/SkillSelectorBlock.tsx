import { AllSkillDTO, EditSkillsDTO } from "../../../types/Edit";
import Select from "../../common/Select";

interface ArgProps {
    allSkills: AllSkillDTO[];
    mySkill: string;
    monster: EditSkillsDTO;
    index: number;
}

const SkillSelectorBlock = ({allSkills, mySkill, monster, index}: ArgProps) => {
    return (
        <Select defaultValue={mySkill}
                styleObj={{marginRight: 0}}
                onChange={(e: React.ChangeEvent<HTMLSelectElement>) => {
                    monster.IsChanged = true;
                    monster.SkillIds[index] = e.target.value;
                }}>
        {
            allSkills.map((skill, index) => (
                <option key={index} value={skill.SkillId}>{skill.SkillName}</option>
            ))
        }
        </Select>
    );
}

export default SkillSelectorBlock;