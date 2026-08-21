export interface JobsResponse {
    TotalCount: number;
    Page:       number;
    PageSize:   number;
    TotalPages: number;
    HasPrev:    boolean;
    HasNext:    boolean;
    Jobs:       Job[];
}

export interface Job {
    Id:               number;
    Url:              string;
    Title:            string;
    Location:         string;
    MinSalaryAtMonth: number;
    MaxSalaryAtMonth: number;
    EmploymentType:   string;
    WorkPlace:        string;
    SourceSite:       string;
    CreatedAt:        string;
    UpdatedAt:        string;
    FeatureNamesCSV:  string;
    OptionsCSV:       string;
    FeatureNames:     string[];
    Options:          string[];
}

export type JobParams = {
    title:        string;
    location:     string;
    workPlace:    string;
    minSalaryAtMonthSpecifiedMin?: number;
    minSalaryAtMonthSpecifiedMax?: number;
    maxSalaryAtMonthSpecifiedMin?: number;
    maxSalaryAtMonthSpecifiedMax?: number;
    sourceSite:   string;
    featureNames: string[];
    options:      string[];
    page:         number;
    pageSize:     number;
    isHideOldJob: boolean;

    setTitle:        React.Dispatch<React.SetStateAction<string>>;
    setLocation:     React.Dispatch<React.SetStateAction<string>>;
    setWorkPlace:    React.Dispatch<React.SetStateAction<string>>;
    setMinSalaryAtMonthSpecifiedMin: React.Dispatch<React.SetStateAction<number | undefined>>;
    setMinSalaryAtMonthSpecifiedMax: React.Dispatch<React.SetStateAction<number | undefined>>;
    setMaxSalaryAtMonthSpecifiedMin: React.Dispatch<React.SetStateAction<number | undefined>>;
    setMaxSalaryAtMonthSpecifiedMax: React.Dispatch<React.SetStateAction<number | undefined>>;
    setSourceSite:   React.Dispatch<React.SetStateAction<string>>;
    setFeatureNames: React.Dispatch<React.SetStateAction<string[]>>;
    setOptions:      React.Dispatch<React.SetStateAction<string[]>>;
    setPage:         React.Dispatch<React.SetStateAction<number>>;
    setPageSize:     React.Dispatch<React.SetStateAction<number>>;
    setIsHideOldJob: React.Dispatch<React.SetStateAction<boolean>>;
};

export type ProjectUsage = {
    FeatureName:  string;
    FeatureCount: number;
    Ratio:        number;
}

export type WorkPlace = {
    Location:   string;
    FullRemote: number;
    Hybrid:     number;
    OnSite:     number;
}

export type SalaryRange = {
    FeatureName:  string;
    SalaryLower:  number;
    SalaryMedian: number;
    SalaryHigher: number;
}