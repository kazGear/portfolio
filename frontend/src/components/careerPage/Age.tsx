import CommonBorderTr from "../common/CommonBorderTr";

interface ArgProps {
    age: number;
}

const Age = ({age}: ArgProps) => {
    return (
        <CommonBorderTr styleObj={{height: "40px"}}>
            <th>年齢</th>
            <td>
                <span>{age}&nbsp;歳</span>
            </td>
        </CommonBorderTr>
    );
}
export default Age;