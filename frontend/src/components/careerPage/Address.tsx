import CommonBorderTr from "../common/CommonBorderTr";

interface ArgProps {
    address: string;
}

const Address = ({address}: ArgProps) => {
    return (
        <CommonBorderTr styleObj={{height: "40px"}}>
            <th>住所</th>
            <td>
                <span>{address}</span>
            </td>
        </CommonBorderTr>
    );
}
export default Address;