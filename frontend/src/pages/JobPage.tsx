import styled from "styled-components";
import { URLS } from "../lib/Constants";
import { useEffect, useState } from "react";
import { api } from "../lib/apiClient";
import useApiErrorHandler from "../hooks/useApiErrorHandler";
import { JobsResponse } from "../types/Job";

const JobPage = () => {

    const errorHandler = useApiErrorHandler();

    /**
     * 店舗アイテム表示
     */
    useEffect(() => {
        api.GET<JobsResponse>(URLS.FETCH_JOBS).then(result => console.log(result))
                                              .catch(errorHandler);

    }, []);

    return (
        <div>
            案件ページ
        </div>
    )
};

export default JobPage;
