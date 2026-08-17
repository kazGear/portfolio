import { useState } from "react";

import { JobParams } from "../types/Job";

export const useJobParams = () => {
    const [title, setTitle]               = useState<string>("");
    const [location, setLocation]         = useState<string>("");
    const [workPlace, setWorkPlace]       = useState<string>("");
    const [minSalaryAtMonthSpecifiedMin, setMinSalaryAtMonthSpecifiedMin] = useState<number | undefined>();
    const [minSalaryAtMonthSpecifiedMax, setMinSalaryAtMonthSpecifiedMax] = useState<number | undefined>();
    const [maxSalaryAtMonthSpecifiedMin, setMaxSalaryAtMonthSpecifiedMin] = useState<number | undefined>();
    const [maxSalaryAtMonthSpecifiedMax, setMaxSalaryAtMonthSpecifiedMax] = useState<number | undefined>();
    const [sourceSite, setSourceSite]     = useState<string>("");
    const [featureNames, setFeatureNames] = useState<string[]>([]);
    const [options, setOptions]           = useState<string[]>([]);
    const [page, setPage]                 = useState<number>(1);
    const [pageSize, setPageSize]         = useState<number>(50);
    const [isHideOldJob, setIsHideOldJob] = useState<boolean>(true); // チェックボックスもONにしておく

    const params: JobParams = {
        title:        title,
        location:     location,
        workPlace:    workPlace,
        minSalaryAtMonthSpecifiedMin: minSalaryAtMonthSpecifiedMin,
        minSalaryAtMonthSpecifiedMax: minSalaryAtMonthSpecifiedMax,
        maxSalaryAtMonthSpecifiedMin: maxSalaryAtMonthSpecifiedMin,
        maxSalaryAtMonthSpecifiedMax: maxSalaryAtMonthSpecifiedMax,
        sourceSite:   sourceSite,
        featureNames: featureNames,
        options:      options,
        page:         page,
        pageSize:     pageSize,
        isHideOldJob: isHideOldJob,

        setTitle:        setTitle,
        setLocation:     setLocation,
        setWorkPlace:    setWorkPlace,
        setMinSalaryAtMonthSpecifiedMin: setMinSalaryAtMonthSpecifiedMin,
        setMinSalaryAtMonthSpecifiedMax: setMinSalaryAtMonthSpecifiedMax,
        setMaxSalaryAtMonthSpecifiedMin: setMaxSalaryAtMonthSpecifiedMin,
        setMaxSalaryAtMonthSpecifiedMax: setMaxSalaryAtMonthSpecifiedMax,
        setSourceSite:   setSourceSite,
        setFeatureNames: setFeatureNames,
        setOptions:      setOptions,
        setPage:         setPage,
        setPageSize:     setPageSize,
        setIsHideOldJob: setIsHideOldJob,
    };
    return params;
}