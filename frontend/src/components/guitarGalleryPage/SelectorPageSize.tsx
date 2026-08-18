import { GuitarParams } from "../../types/Guitar";
import CommonInput from "../common/CommonInput";

interface ArgProps {
    guitarParams: GuitarParams;
    styleObj?:    React.CSSProperties;

}

const SelectorPageSize = ({guitarParams, styleObj}: ArgProps) => {
    const gParams = guitarParams;

    const changePageSizeHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setPageSize(Number(e.currentTarget.value));
    }

    return (
        <CommonInput inputType="number"
                     onBlur={changePageSizeHandler}
                     placeholder=" (10 ~ 100) default 50"
                     min="10"
                     max="100"
                     styleObj={styleObj}/>
    );
}
export default SelectorPageSize;