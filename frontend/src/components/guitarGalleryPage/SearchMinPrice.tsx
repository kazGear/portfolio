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
            gParams.setMinPrice(0);
        } else {
            gParams.setMinPrice(Number(e.currentTarget.value));
        }
    }

    return (
        <CommonInput inputType="number"
                     onBlur={changeMinPriceHandler}
                     min="-3"
                     placeholder="（金額を入力）"
                     defaultValue={gParams.minPrice === 0 ? "" : gParams.minPrice}
                     styleObj={styleObj}
                     />
    );
}
export default SearchMinPrice;