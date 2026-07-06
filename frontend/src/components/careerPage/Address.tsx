import BorderTr from "../common/BorderTr";

interface ArgProps {
    address: string;
}

const Address = ({address}: ArgProps) => {
    return (
        <BorderTr styleObj={{height: "40px"}}>
            <th>住所</th>
            <td>
                <span>{address}</span>
            </td>
        </BorderTr>
    );
}
export default Address;