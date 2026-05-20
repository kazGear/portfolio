import styled from "styled-components";
import { BattleReportDTO } from "../../types/BattleReport";
import { COLORS } from "../../lib/Constants";
import monsterImages from "../../lib/MonsterImages";
import React from "react";
import BorderTd from "../common/BorderTd";
import NowLoading from "../common/NowLoading";

const Stable = styled.table`
    width: 100%;
    border-collapse: collapse;
`;
const StdHeader = styled.td`
    border-top: ${COLORS.BORDER_COLOR} 5px double;
    border-left: ${COLORS.BORDER_COLOR} 1px solid;
    border-right: ${COLORS.BORDER_COLOR} 1px solid;
    padding-left: 10px;
    color: ${COLORS.CAPTION_FONT_COLOR};
    height: 25px;
    align-content: start;
    font-weight: bold;
`;
const Simg = styled.img`
    width: 50px;
    height: 50px;
    vertical-align: middle;
`;

interface ArgProps {
    battleReport: BattleReportDTO[];
    isNowLoadingBattleReport: boolean;
}

const BattleReportBlock = (
    {battleReport, isNowLoadingBattleReport}: ArgProps
) => {

    /**
     * ローディング
     */
    if (isNowLoadingBattleReport) {
        return (
            <div style={{margin: "100px"}}>
                <NowLoading alt="ローディング"/>
            </div>
        );
    }

    return (
        <Stable>
            <tbody>
            {
                battleReport.map((report, index) => {
                    return (
                        <React.Fragment key={index}>
                        { report.Serial === 1 ? // ヘッダーは１試合に１つ
                            <tr>
                                <StdHeader colSpan={4} >
                                    No. {report.BattleId}
                                    &emsp;&emsp;{report.BattleEndDateStr}
                                    &emsp;{report.BattleEndTimeStr}
                                </StdHeader>
                            </tr>
                            : ""
                        }
                            <tr>
                                <BorderTd><span style={{marginLeft: "20px"}}>{report.Serial}</span></BorderTd>
                                <BorderTd>{report.MonsterName}</BorderTd>
                                <BorderTd>
                                    <Simg src={monsterImages(report.MonsterId)} alt=""/>
                                </BorderTd>
                                <BorderTd>
                                    <span style={{color: COLORS.ACCENT_FONT_PINK}}>
                                        {report.IsWin ? "Winner !!" : ""}
                                    </span>
                                </BorderTd>
                            </tr>
                        </React.Fragment>
                    )
                })
            }
            </tbody>
        </Stable>
    );
}

export default BattleReportBlock;