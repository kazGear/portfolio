import { JobParams } from "../../types/Job";
import { ChangeEvent, useEffect } from "react";
import CommonSelect from "../common/CommonSelect";
import { prefectures } from "../../data/job";

interface ArgProps {
    jobParams:      JobParams;
    searchHandler: (jParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;
}

const SearchLocation = ({jobParams, searchHandler, styleObj}: ArgProps) => {

    const changeLocationHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        jobParams.setLocation(e.currentTarget.value);
    }

    // 所在地を選択した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
        jobParams.setPage(1)
    }, [jobParams.location])

    return (
        <CommonSelect onChange={changeLocationHandler} >
            <option value="">未選択</option>
            {
                prefectures.map(elem =>
                        <option key={elem}
                                value={elem}>
                            {elem}
                        </option>
                        )
            }
        </CommonSelect>
    );
}
export default SearchLocation;