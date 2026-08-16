import { useEffect } from "react";
import { JobParams, JobsResponse } from "../../types/Job";
import CommonPagination from "../common/CommonPagination";

interface ArgProps {
    jobsRes:         JobsResponse | null;
    jobParams:      JobParams;
    searchHandler: (jobParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;

}

const SelectorPage = ({jobsRes, jobParams, searchHandler, styleObj}: ArgProps) => {
    const changePrevPageHandler = () => {
        jobParams.setPage(jobParams.page - 1);
    }

    const changeNextPageHandler = () => {
        jobParams.setPage(jobParams.page + 1);
    }

    // ページを変更した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
    }, [jobParams.page])

    return (
        <CommonPagination changePrevPageHandler={changePrevPageHandler}
                          changeNextPageHandler={changeNextPageHandler}
                          hasPrev={jobsRes !== null ? jobsRes.HasPrev : false}
                          hasNext={jobsRes !== null ? jobsRes.HasNext : false}
                          styleObj={styleObj}>
            <span> {jobsRes?.Page} / {jobsRes?.TotalPages} </span>
        </CommonPagination>
    );
}
export default SelectorPage;