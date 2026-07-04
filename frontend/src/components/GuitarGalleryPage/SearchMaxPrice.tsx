import { GuitarParams } from "../../types/Guitar";
import Input from "../common/Input";

interface ArgProps {
    guitarParams: GuitarParams;
}

const SearchMaxPrice = ({guitarParams}: ArgProps) => {
    const gParams = guitarParams;

    const changeMaxPriceHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setMaxPrice(Number(e.currentTarget.value));
    }

    return (
        <Input inputType="number"
               onBlur={changeMaxPriceHandler}
               min="-3"
               placeholder="（金額を入力）"/>
    );
}
export default SearchMaxPrice;