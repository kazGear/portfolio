import CommonBorderTr from "../common/CommonBorderTr";

interface ArgProps {
    tools: string[];
}

const Tools = ({tools}: ArgProps) => {
    return (
        <CommonBorderTr>
            <th>ツール</th>
            <td>
                {
                    tools.map(tool =>
                        <p key={tool}>{tool}</p>
                    )
                }
            </td>
        </CommonBorderTr>
    );
}
export default Tools;