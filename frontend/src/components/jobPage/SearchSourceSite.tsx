import { JobParams } from "../../types/Job";
import { ChangeEvent } from "react";
import CommonSelect from "../common/CommonSelect";
import { sourceSites } from "../../data/job";

interface ArgProps {
    jobParams: JobParams;
    styleObj?: React.CSSProperties;
}

const SearchSourceSite = ({jobParams, styleObj}: ArgProps) => {

    const changeSourceSiteHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        jobParams.setSourceSite(e.currentTarget.value);
    }

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