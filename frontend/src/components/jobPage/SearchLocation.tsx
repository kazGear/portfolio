import { JobParams } from "../../types/Job";
import { ChangeEvent } from "react";
import CommonSelect from "../common/CommonSelect";
import { prefectures } from "../../data/job";

interface ArgProps {
    jobParams: JobParams;
    styleObj?: React.CSSProperties;
}

const SearchLocation = ({jobParams, styleObj}: ArgProps) => {

    const changeLocationHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        jobParams.setLocation(e.currentTarget.value);
    }

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