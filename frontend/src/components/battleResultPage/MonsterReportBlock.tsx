import styled from "styled-components";
import { MonsterReportDTO } from "../../types/BattleReport";
import { COLORS } from "../../lib/Constants";
import monsterImages from "../../lib/MonsterImages";
import BorderTd from "../common/BorderTd";
import NowLoading from "../common/NowLoading";

const Stable = styled.table`
    width: 100%;
    border-collapse: collapse;
    position: relative;
`;
const StHead = styled.thead`
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
const Simg = styled.img`
    width: 30px;
    height: 30px;
    vertical-align: middle;
`;
const Sradio = styled.input`
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
                <NowLoading alt="ローディング"/>
            </div>
        );
    }

   return (
        <div>
            <Stable>
                <StHead>
                    <tr>
                        <BorderTd>
                            <label style={{marginLeft: "40px"}}>
                                モンスター名<Sradio type="radio" name="sortType" value="1" onChange={sortHandler}/>
                            </label>
                        </BorderTd>
                        <BorderTd> </BorderTd>
                        <BorderTd><label>勝利数<Sradio type="radio" name="sortType" value="2" onChange={sortHandler}/></label></BorderTd>
                        <BorderTd><label>対戦数<Sradio type="radio" name="sortType" value="3" onChange={sortHandler}/></label></BorderTd>
                        <BorderTd><label>勝率<Sradio type="radio" name="sortType" value="4" onChange={sortHandler}/></label></BorderTd>
                    </tr>
                </StHead>
                <tbody>
                {
                    monsterReport.map((report) => {
                        return (
                            <tr key={report.MonsterId}>
                                <BorderTd><span style={{marginLeft: "40px"}}>{report.MonsterName}</span></BorderTd>
                                <BorderTd>
                                    <Simg src={monsterImages(report.MonsterId)}
                                          alt=""
                                          style={{height: "40px", width: "40px"}}/>
                                </BorderTd>
                                <BorderTd>{report.Wins} 勝</BorderTd>
                                <BorderTd>{report.BattleCount} 戦</BorderTd>
                                <BorderTd>{report.WinRate}</BorderTd>
                            </tr>
                        )
                    })
                }
                </tbody>
            </Stable>
        </div>
    );
}

export default MonsterReportBlock;