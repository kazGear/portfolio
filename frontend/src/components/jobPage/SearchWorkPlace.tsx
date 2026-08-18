import { JobParams } from "../../types/Job";
import { ChangeEvent } from "react";
import CommonSelect from "../common/CommonSelect";
import { workPlaces } from "../../data/job";

interface ArgProps {
    jobParams: JobParams;
    styleObj?: React.CSSProperties;
}

const SearchWorkPlace = ({jobParams, styleObj}: ArgProps) => {

    const changeWorkPlace = (e: ChangeEvent<HTMLSelectElement>) => {
        jobParams.setWorkPlace(e.currentTarget.value);
    }

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