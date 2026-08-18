import { GuitarParams } from "../../types/Guitar";
import CommonSelect from "../common/CommonSelect";
import { ChangeEvent } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
}

const SelectorOrder = ({guitarParams}: ArgProps) => {
    const gParams = guitarParams;

    const changeOrderHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setOrder(e.target.value);
    }

    return (
        <CommonSelect onChange={changeOrderHandler} >
            <option value="ASC">昇順</option>
            <option value="DESC">降順</option>
        </CommonSelect>
    );
}
export default SelectorOrder;