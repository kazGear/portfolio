import CommonInput from "../common/CommonInput";
import { JobParams } from "../../types/Job";

interface ArgProps {
    jobParams: JobParams;
    styleObj?: React.CSSProperties;
}

const SearchMinSalarySpecifiedMax = ({jobParams, styleObj}: ArgProps) => {
    const changeMaxPriceHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        if (e.currentTarget.value === "") {
            jobParams.setMinSalaryAtMonthSpecifiedMax(undefined);
        } else {
            jobParams.setMinSalaryAtMonthSpecifiedMax(Number(e.currentTarget.value));
        }
    }

    return (
        <CommonInput inputType="number"
                     onBlur={changeMaxPriceHandler}
                     min="-3"
                     placeholder="（上限を入力）"
                     styleObj={styleObj}/>
    );
}
export default SearchMinSalarySpecifiedMax;