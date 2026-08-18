import { JobParams, JobsResponse } from "../../types/Job";
import CommonPagination from "../common/CommonPagination";

interface ArgProps {
    jobsRes:   JobsResponse | null;
    jobParams: JobParams;
    styleObj?: React.CSSProperties;

}

const SelectorPage = ({jobsRes, jobParams, styleObj}: ArgProps) => {
    const changePrevPageHandler = () => {
        jobParams.setPage(jobParams.page - 1);
    }

    const changeNextPageHandler = () => {
        jobParams.setPage(jobParams.page + 1);
    }

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