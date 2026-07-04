import { GuitarParams } from "../../types/Guitar";
import Select from "../common/Select";
import { ChangeEvent, useEffect } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SelectorOrder = ({guitarParams, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeOrderHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setOrder(e.target.value);
    }

    // 並び順を設定した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.order])

    return (
        <Select onChange={changeOrderHandler} >
            <option value="ASC">昇順</option>
            <option value="DESC">降順</option>
        </Select>
    );
}
export default SelectorOrder;