import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import CommonSelect from "../common/CommonSelect";
import { ChangeEvent, useEffect } from "react";

interface ArgProps {
    guitarParams:   GuitarParams;
    series:         Code[] | null;
    searchHandler: (gParams: GuitarParams) => Promise<void>;
}

const SearchSeries = ({guitarParams, series, searchHandler}: ArgProps) => {
    const gParams = guitarParams;

    const changeSeriesHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setSeries(e.currentTarget.value);
    }

    // シリーズを選択した時点で検索実行
    useEffect(() => {
        searchHandler(gParams)
        gParams.setPage(1)
    }, [gParams.series])

    return (
        <CommonSelect onChange={changeSeriesHandler} >
            <option value="">未選択</option>
            {
                series?.map(series =>
                        <option key={series.name}
                                value={series.name}>
                            {series.name}
                        </option>
                        )
            }
        </CommonSelect>
    );
}
export default SearchSeries;