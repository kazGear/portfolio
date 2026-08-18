import CommonInput from "../common/CommonInput";
import { JobParams } from "../../types/Job";

interface ArgProps {
    jobParams: JobParams;
    styleObj?: React.CSSProperties;
}

const SearchTitle = ({jobParams, styleObj}: ArgProps) => {

    const changeNameHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        jobParams.setTitle(e.currentTarget.value);
    }

    return (
        <CommonInput inputType="text"
                     onBlur={changeNameHandler}
                     placeholder="（部分一致検索）"
                     styleObj={styleObj}/>
    );
}
export default SearchTitle;