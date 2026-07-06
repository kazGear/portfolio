import BorderTr from "../common/BorderTr";

interface ArgProps {
    tools: string[];
}

const Tools = ({tools}: ArgProps) => {
    return (
        <BorderTr>
            <th>ツール</th>
            <td>
                {
                    tools.map(tool =>
                        <p key={tool}>{tool}</p>
                    )
                }
            </td>
        </BorderTr>
    );
}
export default Tools;