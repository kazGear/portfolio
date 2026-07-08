import styled from "styled-components";
import { MonsterReportDTO } from "../../types/BattleReport";
import { COLORS } from "../../lib/Constants";
import monsterImages from "../../lib/MonsterImages";
import CommonBorderTd from "../common/CommonBorderTd";
import CommonNowLoading from "../common/CommonNowLoading";

const Table = styled.table`
    width: 100%;
    border-collapse: collapse;
    position: relative;
`;
const Thead = styled.thead`
    height: 35px;
    max-height: 35px;
    color: ${COLORS.CAPTION_FONT_COLOR};
    border-top: ${COLORS.BORDER_COLOR} 1px solid;
    border-bottom: ${COLORS.BORDER_COLOR} 1px solid;
    position: sticky;
    top: 0;
    transform: translateY(-1px); // 上にスクロールしたものが見えてしまうので蓋をする
    font-weight: bold;
    background-color: white;
`;
const Img = styled.img`
    width: 30px;
    height: 30px;
    vertical-align: middle;
`;
const Radio = styled.input`
    margin-left: 6px;
`;

interface ArgProps {
    monsterReport: MonsterReportDTO[];
    setSortType:  React.Dispatch<React.SetStateAction<string>>
    isNowLoadingMonsterReport: boolean;
}

const MonsterReportBlock = (
    {monsterReport, setSortType, isNowLoadingMonsterReport}: ArgProps
) => {
    // ソート項目
    const sortHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSortType(e.target.value);
    }

    /**
     * ローディング
     */
    if (isNowLoadingMonsterReport) {
        return (
            <div style={{margin: "100px"}}>
                <CommonNowLoading alt="ローディング"/>
            </div>
        );
    }

   return (
        <div>
            <Table>
                <Thead>
                    <tr>
                        <CommonBorderTd>
                            <label style={{marginLeft: "40px"}}>
                                モンスター名<Radio type="radio" name="sortType" value="1" onChange={sortHandler}/>
                            </label>
                        </CommonBorderTd>
                        <CommonBorderTd> </CommonBorderTd>
                        <CommonBorderTd><label>勝利数<Radio type="radio" name="sortType" value="2" onChange={sortHandler}/></label></CommonBorderTd>
                        <CommonBorderTd><label>対戦数<Radio type="radio" name="sortType" value="3" onChange={sortHandler}/></label></CommonBorderTd>
                        <CommonBorderTd><label>勝率<Radio type="radio" name="sortType" value="4" onChange={sortHandler}/></label></CommonBorderTd>
                    </tr>
                </Thead>
                <tbody>
                {
                    monsterReport.map((report) => {
                        return (
                            <tr key={report.MonsterId}>
                                <CommonBorderTd><span style={{marginLeft: "40px"}}>{report.MonsterName}</span></CommonBorderTd>
                                <CommonBorderTd>
                                    <Img src={monsterImages(report.MonsterId)}
                                          alt=""
                                          style={{height: "40px", width: "40px"}}/>
                                </CommonBorderTd>
                                <CommonBorderTd>{report.Wins} 勝</CommonBorderTd>
                                <CommonBorderTd>{report.BattleCount} 戦</CommonBorderTd>
                                <CommonBorderTd>{report.WinRate}</CommonBorderTd>
                            </tr>
                        )
                    })
                }
                </tbody>
            </Table>
        </div>
    );
}

export default MonsterReportBlock;