import { JobParams } from "../../types/Job";
import { ChangeEvent, useEffect } from "react";
import CommonSelect from "../common/CommonSelect";
import { workPlaces } from "../../data/job";

interface ArgProps {
    jobParams:      JobParams;
    searchHandler: (jParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;
}

const SearchWorkPlace = ({jobParams, searchHandler, styleObj}: ArgProps) => {

    const changeWorkPlace = (e: ChangeEvent<HTMLSelectElement>) => {
        jobParams.setWorkPlace(e.currentTarget.value);
    }

    // 勤務地を選択した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
        jobParams.setPage(1)
    }, [jobParams.workPlace])

    return (
        <CommonSelect onChange={changeWorkPlace} >
            <option value="">未選択</option>
            {
                workPlaces.map(elem =>
                        <option key={elem}
                                value={elem}>
                            {elem}
                        </option>
                        )
            }
        </CommonSelect>
    );
}
export default SearchWorkPlace;