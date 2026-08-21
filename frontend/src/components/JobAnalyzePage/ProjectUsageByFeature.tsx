import { ResponsiveBar } from '@nivo/bar'
import type { ComputedDatum } from "@nivo/bar";
import { useEffect, useState } from 'react';
import { URLS } from '../../lib/Constants';
import { api } from '../../lib/apiClient';
import { ProjectUsage } from '../../types/Job';
import useApiErrorHandler from '../../hooks/useApiErrorHandler';

interface ArgProps {
    category: string
}

const customLabel = (data: ComputedDatum<ProjectUsage>): string => {
    return (
        `${data.data.FeatureCount?.toLocaleString()}件 ${data.data.Ratio?.toFixed(2)}%`
    )
};

const ProjectUsageByFeature = ({ category }: ArgProps) => {
    const [data, setData]                     = useState<ProjectUsage[]>([]);
    const [marginLeftBase, setMarginLeftBase] = useState<number>(0);

    const errorHandler = useApiErrorHandler();

    useEffect(() => {
        api.GET<ProjectUsage[]>(`${URLS.FETCH_PROJECT_USAGE_BY_FEATURE}?category=${category}`)
           .then(data => setData(data ?? []))
           .catch(errorHandler);

        const base = data.reduce((max, item) => Math.max(max, item.FeatureName.length), 0);
        setMarginLeftBase(base);
    }, []);

    return (
        <div style={{ height: data.length * 30 }}>
            <ResponsiveBar
                data={data}
                keys={["FeatureCount"]}
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
                margin={{ top: 30, right: 140, bottom: 50, left: marginLeftBase + 200 }}
            />
        </div>
    )
};

export default ProjectUsageByFeature;