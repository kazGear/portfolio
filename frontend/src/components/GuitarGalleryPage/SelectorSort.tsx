import { GuitarParams } from "../../types/Guitar";
import Select from "../common/Select";
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
        <Select onChange={changeSortHandler} >
            <option value="price">価格</option>
            <option value="maker">メーカー</option>
            <option value="name">ギター名</option>
        </Select>
    );
}
export default SelectorSort;