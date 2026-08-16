import { useEffect } from "react";
import { JobParams } from "../../types/Job";
import CommonInput from "../common/CommonInput";

interface ArgProps {
    jobParams:      JobParams;
    searchHandler: (jobParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;

}

const SelectorPageSize = ({jobParams, searchHandler, styleObj}: ArgProps) => {
    const changePageSizeHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        jobParams.setPageSize(Number(e.currentTarget.value));
    }

    // ページサイズを設定した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
        jobParams.setPage(1)
    }, [jobParams.pageSize])

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