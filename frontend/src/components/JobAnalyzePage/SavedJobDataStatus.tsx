import { useEffect, useState } from 'react';
import { COLORS, URLS } from '../../lib/Constants';
import { api } from '../../lib/apiClient';
import useApiErrorHandler from '../../hooks/useApiErrorHandler';
import { SavedJobData } from '../../types/Job';
import CommonBorderTr from '../common/CommonBorderTr';
import styled from 'styled-components';
import CommonNowLoading from '../common/CommonNowLoading';

const Th = styled.th`
    text-align: left;
    color: ${COLORS.ACCENT_FONT_PINK};
    font-size: 18px;
`;

const Td = styled.th`
    text-align: left;
    font-size: 18px;
`;

const SavedJobDataStatus = () => {
    const [data, setData] = useState<SavedJobData[]>([]);

    const errorHandler = useApiErrorHandler();

    useEffect(() => {
        api.GET<SavedJobData[]>(`${URLS.FETCH_SAVED_JOB_DATA_STATUS}`)
           .then(data => setData(data ?? []))
           .catch(errorHandler);
    }, []);

    return (
        <table style={{margin: "40px auto", width: "90%"}}>
            <thead>
                <CommonBorderTr>
                    <Th>Source site</Th>
                    <Th>Saved pageID min</Th>
                    <Th>Saved pageID max</Th>
                    <Th>Job count</Th>
                    <Th>Exist ratio</Th>
                </CommonBorderTr>
            </thead>
            <tbody>
                {
                    data.length <= 0 ? <CommonNowLoading alt='Now Loading ...'/> :
                    data.map((data, index) => {
                        return (
                            <CommonBorderTr key={data.SourceSite + index}>
                                <Td>{data.SourceSite}</Td>
                                <Td>{data.SavedPageIdMin.toLocaleString()}</Td>
                                <Td>{data.SavedPageIdMax.toLocaleString()}</Td>
                                <Td>{data.JobCount.toLocaleString()} 件</Td>
                                <Td>{data.ExistRatio} %</Td>
                            </CommonBorderTr>
                        )
                    })
                }
            </tbody>
        </table>
    )
};

export default SavedJobDataStatus;