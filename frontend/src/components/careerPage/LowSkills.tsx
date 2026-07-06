import BorderTr from "../common/BorderTr";

interface ArgProps {
    lowSkills: string[];
}

const LowSkills = ({lowSkills}: ArgProps) => {
    return (
        <BorderTr>
            <th>スキル<br/>(習熟度・低)</th>
            <td>
                {
                    lowSkills.map(skill =>
                        <p key={skill}>{skill}</p>
                    )
                }
            </td>
        </BorderTr>
    );
}
export default LowSkills;