
import BorderTr from "../common/BorderTr";

interface ArgProps {
    skills: string[];
}

const Skills = ({skills}: ArgProps) => {
    return (
        <BorderTr>
            <th>スキル</th>
            <td>
                {
                    skills.map(skill =>
                        <p key={skill}>{skill}</p>
                    )
                }
            </td>
        </BorderTr>
    );
}
export default Skills;