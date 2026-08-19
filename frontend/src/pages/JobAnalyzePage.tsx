import { ResponsiveBar } from '@nivo/bar'
import type { ComputedDatum } from "@nivo/bar";
import useApiErrorHandler from "../hooks/useApiErrorHandler";
import CommonNowLoading from "../components/common/CommonNowLoading";
import CommonFrame from "../components/common/CommonFrame";
import { useEffect, useState } from 'react';
import { ProjectUsageByLanguage } from '../types/Job';
import { api } from '../lib/apiClient';
import { URLS } from '../lib/Constants';

type LanguageData = {
    jobCount: number;
    ratio: number;
};

const customLabel = (data: ComputedDatum<ProjectUsageByLanguage>): string => {
    return (
        `${data.data.JobCount?.toLocaleString()}件 ${data.data.Ratio?.toFixed(2)}%`
    )
};

const JobAnalyzePage = () => {
    const [data, setData] = useState<ProjectUsageByLanguage[]>([])
    const errorHandler    = useApiErrorHandler();

    useEffect(() => {
        api.GET<ProjectUsageByLanguage[]>(URLS.FETCH_PROJECT_USAGE_BY_LANGUAGE)
           .then(data => setData(data ?? []))
           .catch(errorHandler);
    }, []);

    return (
        <div>
            <h1 style={{background: "white", paddingLeft: "40px"}}>👷実装中👷</h1>
            <CommonFrame styleObj={{margin: "0px 20px", height: "75vh"}}>
                <ResponsiveBar
                    data={data}
                    keys={["JobCount"]}
                    indexBy="FeatureName"
                    colors={{"scheme": "dark2"}}
                    layout="horizontal"
                    label={customLabel}
                    labelPosition="end"
                    labelOffset={10}
                    labelSkipWidth={12}
                    labelSkipHeight={12}
                    legends={[
                        {
                            dataFrom: 'keys',
                            anchor: 'bottom-right',
                            direction: 'column',
                            translateX: 120,
                            itemsSpacing: 3,
                            itemWidth: 100,
                            itemHeight: 16
                        }
                    ]}
                    axisBottom={{ legend: '利用状況', legendOffset: 32 }}
                    axisLeft={{ legend: '', legendOffset: 0 }}
                    margin={{ top: 50, right: 150, bottom: 50, left: 80 }}
                />
            </CommonFrame>
        </div>
    )
};

export default JobAnalyzePage;
