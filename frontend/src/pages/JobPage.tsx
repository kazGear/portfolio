import { URLS } from "../lib/Constants";
import { useEffect, useState } from "react";
import { api } from "../lib/apiClient";
import useApiErrorHandler from "../hooks/useApiErrorHandler";
import { JobParams, JobsResponse } from "../types/Job";
import CommonFrame from "../components/common/CommonFrame";
import JobCards from "../components/jobPage/JobCards";
import SearchConditionsJob from "../components/jobPage/SearchConditionsJob";
import { useJobParams } from "../hooks/useJobParams";
import { createQueryParamsJob } from "../components/jobPage/JobFuncs";
import CommonNowLoading from "../components/common/CommonNowLoading";

const JobPage = () => {
    const [jobs, setJobs]                         = useState<JobsResponse | null>(null);
    const [languages, setLanguages]               = useState<string[] | null>([]);
    const [frameworkLibrary, setFrameworkLibrary] = useState<string[] | null>([]);
    const [role, setRole]                         = useState<string[] | null>([]);
    const [infrastructure, setInfrastructure]     = useState<string[] | null>([]);
    const [database, setDatabase]                 = useState<string[] | null>([]);
    const [cloud, setCloud]                       = useState<string[] | null>([]);

    const jobParams: JobParams = useJobParams();

    const errorHandler = useApiErrorHandler();

    /**
     * 初期化データ取得
     */
    useEffect(() => {
        api.GET<JobsResponse>(URLS.FETCH_JOBS).then(result => setJobs(result)).catch(errorHandler);

        api.GET<string[]>(URLS.FETCH_FEATURES + "?category=LANGUAGE")
           .then(result => setLanguages(result))
           .catch(errorHandler);
        api.GET<string[]>(URLS.FETCH_FEATURES + "?category=FRAMEWORK_LIBRARY")
           .then(result => setFrameworkLibrary(result))
           .catch(errorHandler);
        api.GET<string[]>(URLS.FETCH_FEATURES + "?category=ROLE")
           .then(result => setRole(result))
           .catch(errorHandler);
        api.GET<string[]>(URLS.FETCH_FEATURES + "?category=INFRASTRUCTURE")
           .then(result => setInfrastructure(result))
           .catch(errorHandler);
        api.GET<string[]>(URLS.FETCH_FEATURES + "?category=DATABASE")
           .then(result => setDatabase(result))
           .catch(errorHandler);
        api.GET<string[]>(URLS.FETCH_FEATURES + "?category=CLOUD")
           .then(result => setCloud(result))
           .catch(errorHandler);
    }, []);

    // 案件データ取得
    const jobSearchHandler = async (jobParams: JobParams) => {
        const queryParams = createQueryParamsJob(jobParams);
        setJobs(null); // 案件表示エリアにローディングアイコンを表示するため

        const resJobs = await api.GET<JobsResponse>(`${URLS.FETCH_JOBS}?${queryParams.toString()}`);

        setJobs(resJobs);
    };

    // 条件を変更した時点で検索実行
    useEffect(() => {
        jobSearchHandler(jobParams);
        jobParams.setPage(1);
    }, [
        jobParams.title,
        jobParams.location,
        jobParams.workPlace,
        jobParams.minSalaryAtMonthSpecifiedMin,
        jobParams.minSalaryAtMonthSpecifiedMax,
        jobParams.maxSalaryAtMonthSpecifiedMin,
        jobParams.maxSalaryAtMonthSpecifiedMax,
        jobParams.sourceSite,
        jobParams.pageSize,
        jobParams.isHideOldJob,
        jobParams.featureNames,
    ]);

    // ページ送り
    useEffect(() => {
        jobSearchHandler(jobParams);
    }, [jobParams.page]);

    return (
        <div style={{ display: "flex" }}>
            <CommonFrame styleObj={{width: "30%", minWidth: "360px", height: "87vh", margin: "20px 0px 0px 20px"}}>
                <SearchConditionsJob jobsRes={jobs}
                                     jobParams={jobParams}
                                     languages={languages}
                                     frameworkLibrary={frameworkLibrary}
                                     role={role}
                                     infrastructure={infrastructure}
                                     database={database}
                                     cloud={cloud} />
            </CommonFrame>
            <CommonFrame styleObj={{width: "70%", minWidth: "360px",height: "87vh", margin: "20px 20px 0px 10px"}}>
                {
                    jobs !== null ? (
                        <JobCards jobsRes={jobs} />
                    ) : (
                        <div style={{textAlign: "center", marginTop: "40%"}}>
                            <CommonNowLoading alt="各案件データ" size="300" />
                        </div>
                    )
                }
            </CommonFrame>
        </div>
    )
};

export default JobPage;
