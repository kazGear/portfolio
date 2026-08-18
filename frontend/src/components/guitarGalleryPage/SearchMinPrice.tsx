import { GUITAR } from "../../lib/Constants";
import { GuitarParams } from "../../types/Guitar";
import CommonInput from "../common/CommonInput";

interface ArgProps {
    guitarParams: GuitarParams;
    styleObj?:    React.CSSProperties;
}

const SearchMinPrice = ({guitarParams, styleObj}: ArgProps) => {
    const gParams = guitarParams;

    const changeMinPriceHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        if (e.currentTarget.value === "") {
            // 0で更新すると全件検索されなくなる
            gParams.setMinPrice(GUITAR.PARSE_ERROR_PRICE);
        } else {
            gParams.setMinPrice(Number(e.currentTarget.value));
        }
    }

    return (
        <CommonInput inputType="number"
                     onBlur={changeMinPriceHandler}
                     min="-3"
                     placeholder="（金額を入力）"
                     styleObj={styleObj}
                     />
    );
}
export default SearchMinPrice;