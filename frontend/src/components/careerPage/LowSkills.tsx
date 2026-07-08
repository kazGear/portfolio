import CommonBorderTr from "../common/CommonBorderTr";

interface ArgProps {
    lowSkills: string[];
}

const LowSkills = ({lowSkills}: ArgProps) => {
    return (
        <CommonBorderTr>
            <th>スキル<br/>(習熟度・低)</th>
            <td>
                {
                    lowSkills.map(skill =>
                        <p key={skill}>{skill}</p>
                    )
                }
            </td>
        </CommonBorderTr>
    );
}
export default LowSkills;