import { useEffect } from "react";
import { GuitarParams } from "../../types/Guitar";
import Input from "../common/Input";

interface ArgProps {
    guitarParams: GuitarParams;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SearchMaxPrice = ({guitarParams, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeMaxPriceHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        if (e.currentTarget.value === "") {
            // 多額で更新して全件ヒットさせる
            gParams.setMaxPrice(100000000);
        } else {
            gParams.setMaxPrice(Number(e.currentTarget.value));
        }
    }

   // 価格を設定した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.maxPrice])

    return (
        <Input inputType="number"
               onBlur={changeMaxPriceHandler}
               min="-3"
               placeholder="（金額を入力）"/>
    );
}
export default SearchMaxPrice;