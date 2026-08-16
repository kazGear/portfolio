import { JobParams } from "../../types/Job";
import { ChangeEvent, useEffect } from "react";
import CommonSelect from "../common/CommonSelect";
import { sourceSites } from "../../data/job";

interface ArgProps {
    jobParams:      JobParams;
    searchHandler: (jParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;
}

const SearchSourceSite = ({jobParams, searchHandler, styleObj}: ArgProps) => {

    const changeSourceSiteHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        jobParams.setSourceSite(e.currentTarget.value);
    }

    // 所在地を選択した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
        jobParams.setPage(1)
    }, [jobParams.sourceSite])

    return (
        <CommonSelect onChange={changeSourceSiteHandler} >
            <option value="">未選択</option>
            {
                sourceSites.map(elem =>
                        <option key={elem}
                                value={elem}>
                            {elem}
                        </option>
                        )
            }
        </CommonSelect>
    );
}
export default SearchSourceSite;