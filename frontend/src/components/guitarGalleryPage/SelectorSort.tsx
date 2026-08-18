import { GuitarParams } from "../../types/Guitar";
import CommonSelect from "../common/CommonSelect";
import { ChangeEvent } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
}

const SelectorSort = ({guitarParams}: ArgProps) => {
    const gParams = guitarParams;

    const changeSortHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setSort(e.target.value);
    }

    return (
        <CommonSelect onChange={changeSortHandler} >
            <option value="name">ギター名</option>
            <option value="price">価格</option>
            <option value="maker">メーカー</option>
        </CommonSelect>
    );
}
export default SelectorSort;