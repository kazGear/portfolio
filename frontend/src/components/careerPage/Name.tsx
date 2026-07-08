import CommonBorderTr from "../common/CommonBorderTr";

interface ArgProps {
    name: string;
}

const Name = ({name}: ArgProps) => {
    return (
        <CommonBorderTr styleObj={{height: "40px"}}>
            <th>氏名</th>
            <td>
                <span>{name}</span>
            </td>
        </CommonBorderTr>
    );
}
export default Name;