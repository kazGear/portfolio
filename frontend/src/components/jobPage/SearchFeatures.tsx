import { JobParams } from "../../types/Job";
import { ChangeEvent, useEffect } from "react";
import CommonBreadcrumbsList from "../common/CommonBreadcrumbsList";
import { elemAddOrRemove } from "../../lib/CommonLogic";

interface ArgProps {
    jobParams:      JobParams;
    features:       string[] | null;
    searchHandler: (jParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;
}

const SearchFeatures = ({jobParams, features, searchHandler, styleObj}: ArgProps) => {

    const clickLanguageHandler = (e: ChangeEvent<HTMLInputElement>) => {
        const clicked = e.currentTarget.value;
        let languages = jobParams.featureNames;

        languages = elemAddOrRemove(languages, clicked);

        jobParams.setFeatureNames([...languages]);
    }

    // 所在地を選択した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
        jobParams.setPage(1)
    }, [jobParams.featureNames])

    return (
        features?.map(elem =>
                <CommonBreadcrumbsList key={elem} value={elem} onChange={clickLanguageHandler} >
                    {elem}
                </CommonBreadcrumbsList>
            )
    );
}
export default SearchFeatures;