import { useEffect } from "react";
import { GUITAR } from "../../lib/Constants";
import { GuitarParams } from "../../types/Guitar";
import CommonInput from "../common/CommonInput";

interface ArgProps {
    guitarParams: GuitarParams;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SearchMinPrice = ({guitarParams, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeMinPriceHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        if (e.currentTarget.value === "") {
            // 0で更新すると全件検索されなくなる
            gParams.setMinPrice(GUITAR.PARSE_ERROR_PRICE);
        } else {
            gParams.setMinPrice(Number(e.currentTarget.value));
        }
    }

    // 価格を設定した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.minPrice])

    return (
        <CommonInput inputType="number"
                     onBlur={changeMinPriceHandler}
                     min="-3"
                     placeholder="（金額を入力）"/>
    );
}
export default SearchMinPrice;