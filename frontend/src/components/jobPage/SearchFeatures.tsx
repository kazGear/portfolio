import { JobParams } from "../../types/Job";
import { ChangeEvent } from "react";
import CommonBreadcrumbsList from "../common/CommonBreadcrumbsList";
import { elemAddOrRemove } from "../../lib/CommonLogic";

interface ArgProps {
    jobParams: JobParams;
    features:  string[] | null;
    styleObj?: React.CSSProperties;
}

const SearchFeatures = ({jobParams, features, styleObj}: ArgProps) => {

    const clickFeaturesHandler = (e: ChangeEvent<HTMLInputElement>) => {
        const clicked = e.currentTarget.value;
        let features  = jobParams.featureNames;

        features = elemAddOrRemove(features, clicked);

        jobParams.setFeatureNames([...features]);
    }

    return (
        features?.map(elem =>
                <CommonBreadcrumbsList key={elem} value={elem} onChange={clickFeaturesHandler} >
                    {elem}
                </CommonBreadcrumbsList>
            )
    );
}
export default SearchFeatures;