import styled from "styled-components";
import { BattleReportDTO } from "../../types/BattleReport";
import { COLORS } from "../../lib/Constants";
import monsterImages from "../../lib/MonsterImages";
import React from "react";
import CommonBorderTd from "../common/CommonBorderTd";
import CommonNowLoading from "../common/CommonNowLoading";

const Table = styled.table`
    width: 100%;
    border-collapse: collapse;
`;
const TdHeader = styled.td`
    border-top: ${COLORS.BORDER_COLOR} 5px double;
    border-left: ${COLORS.BORDER_COLOR} 1px solid;
    border-right: ${COLORS.BORDER_COLOR} 1px solid;
    padding-left: 10px;
    color: ${COLORS.CAPTION_FONT_COLOR};
    height: 25px;
    align-content: start;
    font-weight: bold;
`;
const Img = styled.img`
    width: 50px;
    height: 50px;
    vertical-align: middle;
`;

interface ArgProps {
    battleReport:             BattleReportDTO[];
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
                <CommonNowLoading alt="ローディング"/>
            </div>
        );
    }

    return (
        <Table>
            <tbody>
            {
                battleReport.map((report, index) => {
                    return (
                        <React.Fragment key={index + report.BattleId}>
                        { report.Serial === 1 ? // ヘッダーは１試合に１つ
                            <tr>
                                <TdHeader colSpan={4} >
                                    No. {report.BattleId}
                                    &emsp;&emsp;{report.BattleEndDateStr}
                                    &emsp;{report.BattleEndTimeStr}
                                </TdHeader>
                            </tr>
                            : ""
                        }
                            <tr>
                                <CommonBorderTd><span style={{marginLeft: "20px"}}>{report.Serial}</span></CommonBorderTd>
                                <CommonBorderTd>{report.MonsterName}</CommonBorderTd>
                                <CommonBorderTd>
                                    <Img src={monsterImages(report.MonsterId)} alt=""/>
                                </CommonBorderTd>
                                <CommonBorderTd>
                                    <span style={{color: COLORS.ACCENT_FONT_PINK}}>
                                        {report.IsWin ? "Winner !!" : ""}
                                    </span>
                                </CommonBorderTd>
                            </tr>
                        </React.Fragment>
                    )
                })
            }
            </tbody>
        </Table>
    );
}

export default BattleReportBlock;