
import CommonBorderTr from "../common/CommonBorderTr";

interface ArgProps {
    skills: string[];
}

const Skills = ({skills}: ArgProps) => {
    return (
        <CommonBorderTr>
            <th>スキル</th>
            <td>
                {
                    skills.map(skill =>
                        <p key={skill}>{skill}</p>
                    )
                }
            </td>
        </CommonBorderTr>
    );
}
export default Skills;