import { useEffect } from "react";
import CommonInput from "../common/CommonInput";
import { JobParams } from "../../types/Job";

interface ArgProps {
    jobParams:      JobParams;
    searchHandler: (jParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;
}

const SearchMaxSalarySpecifiedMax = ({jobParams, searchHandler, styleObj}: ArgProps) => {
    const changeMaxPriceHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        if (e.currentTarget.value === "") {
            jobParams.setMaxSalaryAtMonthSpecifiedMax(undefined);
        } else {
            jobParams.setMaxSalaryAtMonthSpecifiedMax(Number(e.currentTarget.value));
        }
    }

    // 価格を設定した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
        jobParams.setPage(1)
    }, [jobParams.maxSalaryAtMonthSpecifiedMax])

    return (
        <CommonInput inputType="number"
                     onBlur={changeMaxPriceHandler}
                     min="-3"
                     placeholder="（上限を入力）"
                     styleObj={styleObj}/>
    );
}
export default SearchMaxSalarySpecifiedMax;