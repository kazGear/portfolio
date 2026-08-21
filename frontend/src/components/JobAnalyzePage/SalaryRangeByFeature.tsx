import { ResponsiveBar } from '@nivo/bar'
import type { ComputedDatum } from "@nivo/bar";
import { useEffect, useState } from 'react';
import { URLS } from '../../lib/Constants';
import { api } from '../../lib/apiClient';
import { SalaryRange } from '../../types/Job';
import useApiErrorHandler from '../../hooks/useApiErrorHandler';

interface ArgProps {
    category: string
}

const customLabel = (data: ComputedDatum<SalaryRange>): string => {
    return `${Number(data.value).toLocaleString()}円`;
};

const SalaryRangeByFeature = ({ category }: ArgProps) => {
    const [data, setData]                     = useState<SalaryRange[]>([]);
    const [marginLeftBase, setMarginLeftBase] = useState<number>(0);

    const errorHandler = useApiErrorHandler();

    useEffect(() => {
        api.GET<SalaryRange[]>(`${URLS.FETCH_SALARY_RANGE_BY_FEATURE}?category=${category}`)
           .then(data => setData(data ?? []))
           .catch(errorHandler);

        const base = data.reduce((max, item) => Math.max(max, item.FeatureName.length), 0);
        setMarginLeftBase(base);
    }, []);

    return (
        <div style={{ height: data.length * 60 }}>
            <ResponsiveBar
                data={data}
                keys={["SalaryHigher", "SalaryMedian", "SalaryLower"]}
                indexBy="FeatureName"
                colors={{"scheme": "pastel1"}}
                layout="horizontal"
                groupMode="grouped"
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
                axisBottom={{ legend: '給与', legendOffset: 32 }}
                axisLeft={{ legend: '', legendOffset: 0 }}
                margin={{ top: 30, right: 140, bottom: 50, left: marginLeftBase + 200 }}
            />
        </div>
    )
};

export default SalaryRangeByFeature;