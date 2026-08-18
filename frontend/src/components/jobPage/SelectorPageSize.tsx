import { JobParams } from "../../types/Job";
import CommonInput from "../common/CommonInput";

interface ArgProps {
    jobParams: JobParams;
    styleObj?: React.CSSProperties;

}

const SelectorPageSize = ({jobParams, styleObj}: ArgProps) => {
    const changePageSizeHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        jobParams.setPageSize(Number(e.currentTarget.value));
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