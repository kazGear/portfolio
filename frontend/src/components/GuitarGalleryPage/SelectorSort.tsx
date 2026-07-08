import { GuitarParams } from "../../types/Guitar";
import CommonSelect from "../common/CommonSelect";
import { ChangeEvent, useEffect } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SelectorSort = ({guitarParams, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeSortHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setSort(e.target.value);
    }

    // ソートを設定した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.sort])

    return (
        <CommonSelect onChange={changeSortHandler} >
            <option value="price">価格</option>
            <option value="maker">メーカー</option>
            <option value="name">ギター名</option>
        </CommonSelect>
    );
}
export default SelectorSort;