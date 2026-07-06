import BorderTr from "../common/BorderTr";

interface ArgProps {
    age: number;
}

const Age = ({age}: ArgProps) => {
    return (
        <BorderTr styleObj={{height: "40px"}}>
            <th>年齢</th>
            <td>
                <span>{age}&nbsp;歳</span>
            </td>
        </BorderTr>
    );
}
export default Age;