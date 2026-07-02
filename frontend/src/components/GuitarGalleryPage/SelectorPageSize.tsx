import { GuitarParams } from "../../types/Guitar";
import Input from "../common/Input";

interface ArgProps {
    guitarParams: GuitarParams;
}

const SelectorPageSize = ({guitarParams}: ArgProps) => {
    const gParams = guitarParams;

    const changePageSizeHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setPageSize(Number(e.currentTarget.value));
    }

    return (
        <Input inputType="number"
               onBlur={changePageSizeHandler}
               placeholder=" (10 ~ 100) default 25"
               min="10"
               max="100"/>
    );
}
export default SelectorPageSize;