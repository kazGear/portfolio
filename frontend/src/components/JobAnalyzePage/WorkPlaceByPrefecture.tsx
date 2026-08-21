import { ComputedDatum, ResponsiveBar } from '@nivo/bar'
import { useEffect, useState } from 'react';
import { URLS } from '../../lib/Constants';
import { api } from '../../lib/apiClient';
import { WorkPlace } from '../../types/Job';
import useApiErrorHandler from '../../hooks/useApiErrorHandler';

const customLabel = (data: ComputedDatum<WorkPlace>): string => {
    return `${Number(data.value).toLocaleString()} 件`;
};

const WorkPlaceByPrefecture = () => {
    const [data, setData] = useState<WorkPlace[]>([]);

    const errorHandler = useApiErrorHandler();

    useEffect(() => {
        api.GET<WorkPlace[]>(`${URLS.FETCH_WORK_PLACE_BY_PREFECTURE}`)
           .then(data => setData(data ?? []))
           .catch(errorHandler);
    }, []);

    return (
        <div style={{ height: data.length * 60 }}>
            <ResponsiveBar
                data={data}
                keys={["FullRemote", "Hybrid", "OnSite"]}
                indexBy="Location"
                colors={{"scheme": "pastel1"}}
                layout="horizontal"
                groupMode="grouped"
                label={customLabel}
                labelPosition="end"
                labelOffset={10}
                labelSkipWidth={0}
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
                axisBottom={{ legend: '案件数', legendOffset: 32 }}
                axisLeft={{ legend: '', legendOffset: 0 }}
                margin={{ top: 30, right: 120, bottom: 50, left: 70 }}
            />
        </div>
    )
};

export default WorkPlaceByPrefecture;