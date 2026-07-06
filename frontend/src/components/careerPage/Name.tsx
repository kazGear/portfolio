import BorderTr from "../common/BorderTr";

interface ArgProps {
    name: string;
}

const Name = ({name}: ArgProps) => {
    return (
        <BorderTr styleObj={{height: "40px"}}>
            <th>氏名</th>
            <td>
                <span>{name}</span>
            </td>
        </BorderTr>
    );
}
export default Name;