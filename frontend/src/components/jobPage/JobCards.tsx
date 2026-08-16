import { JobsResponse } from "../../types/Job";
import JobCard from "./JobCard";

interface ArgProps {
    jobsRes: JobsResponse | null;
}

const JobCards = ({jobsRes: res}: ArgProps) => {
    return (
        <>
            <p style={{marginLeft: "15px", fontWeight: "bolder"}}>
                検索結果 {res?.TotalCount} 件<br/>
                ページ {res?.Page} / {res?.TotalPages} @{res?.PageSize}件
            </p>
            <div style={{margin: "15px"}}>
                {
                    res?.Jobs.map(job => (
                        <JobCard job={job}
                                 key={job.Url}
                                 />
                    )
                )}
            </div>
        </>
    );
}

export default JobCards;