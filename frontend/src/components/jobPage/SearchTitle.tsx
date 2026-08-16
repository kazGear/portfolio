import CommonInput from "../common/CommonInput";
import { JobParams } from "../../types/Job";
import { useEffect } from "react";

interface ArgProps {
    jobParams:      JobParams;
    searchHandler: (jParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;
}

const SearchTitle = ({jobParams, searchHandler, styleObj}: ArgProps) => {

    const changeNameHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        jobParams.setTitle(e.currentTarget.value);
    }

    // タイトルを入力した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
        jobParams.setPage(1)
    }, [jobParams.title])

    return (
        <CommonInput inputType="text"
                     onBlur={changeNameHandler}
                     placeholder="（部分一致検索）"
                     styleObj={styleObj}/>
    );
}
export default SearchTitle;