import { GuitarParams } from "../../types/Guitar";
import Input from "../common/Input";

interface ArgProps {
    guitarParams: GuitarParams;
}

const SearchMinPrice = ({guitarParams}: ArgProps) => {
    const gParams = guitarParams;

    const changeMinPriceHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        if (e.currentTarget.value === "") {
            gParams.setMinPrice(-3); // 0で更新すると全件検索されなくなる
        } else {
            gParams.setMinPrice(Number(e.currentTarget.value));
        }
    }

    return (
        <Input inputType="number"
               onBlur={changeMinPriceHandler}
               min="-3"
               placeholder="（金額を入力）"/>
    );
}
export default SearchMinPrice;