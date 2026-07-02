import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import Select from "../common/Select";
import { ChangeEvent } from "react";

interface ArgProps {
    guitarParams:        GuitarParams;
    series:              Code[] | null;
}

const SearchSeries = ({guitarParams, series}: ArgProps) => {
    const gParams = guitarParams;

    const changeSeriesHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setSeries(e.target.value);
    }

    return (
        <Select onChange={changeSeriesHandler} >
            <option>未選択</option>
            {
                series?.map(series =>
                        <option key={series.name}
                                value={series.name}>
                            {series.name}
                        </option>
                        )
            }
        </Select>
    );
}
export default SearchSeries;