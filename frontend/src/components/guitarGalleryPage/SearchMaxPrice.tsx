import { GuitarParams } from "../../types/Guitar";
import CommonInput from "../common/CommonInput";

interface ArgProps {
    guitarParams: GuitarParams;
    styleObj?:    React.CSSProperties;
}

const SearchMaxPrice = ({guitarParams, styleObj}: ArgProps) => {
    const gParams = guitarParams;

    const changeMaxPriceHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        if (e.currentTarget.value === "") {
            // 多額で更新して全件ヒットさせる
            gParams.setMaxPrice(100000000);
        } else {
            gParams.setMaxPrice(Number(e.currentTarget.value));
        }
    }

    return (
        <CommonInput inputType="number"
                     onBlur={changeMaxPriceHandler}
                     min="-3"
                     placeholder="（金額を入力）"
                     styleObj={styleObj}/>
    );
}
export default SearchMaxPrice;