import styled from "styled-components";
import { JobParams } from "../../types/Job";
import React, { ChangeEvent, useEffect } from "react";

const Label = styled.label`
    display: inline-block;
    margin: 6px 0px 0px 16px;
    height: 25px;
`;

interface ArgProps {
    jobParams:      JobParams;
    searchHandler: (jParams: JobParams) => Promise<void>;
    styleObj?:      React.CSSProperties;
}

const HideOldJob = ({jobParams, searchHandler, styleObj}: ArgProps) => {

    const toggleHideHandler = (e: ChangeEvent<HTMLInputElement>) => {
        const isHide = jobParams.isHideOldJob;

        // 反対の状態に切り替え
        if (isHide) {
            jobParams.setIsHideOldJob(false);
        } else {
            jobParams.setIsHideOldJob(true);
        }
    }

    // タイトルを入力した時点で検索実行
    useEffect(() => {
        searchHandler(jobParams)
        jobParams.setPage(1)
    }, [jobParams.isHideOldJob])

    return (
        <Label>
            <input type="checkbox"
                   onChange={toggleHideHandler}
                   defaultChecked={true}
                   />
            <span style={{display: "inline-block", transform: "translateY(-2px)"}}>
                古い案件を表示しない
            </span>
        </Label>
    );
}
export default HideOldJob;