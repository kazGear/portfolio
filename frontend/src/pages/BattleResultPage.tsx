import styled from "styled-components";
import { useState } from "react";
import { BattleReportDTO, MonsterReportDTO } from "../types/BattleReport";
import MonsterReport from "../components/battleResultPage/MonsterReportBlock";
import BattleReportControllerBlock from "../components/battleResultPage/BattleReportControllerBlock";
import BattleReportBlock from "../components/battleResultPage/BattleReportBlock";
import MonsterReportControllerBlock from "../components/battleResultPage/MonsterReportControllerBlock";
import CommonFrame from "../components/common/CommonOutSideFrame";
import { useCheckToken } from "../hooks/useHooksOfCommon";

const OutsideFrame = styled.div`
    height: 100%;
`;
const OptionFrame = styled.div`
    display: flex;
    justify-content: center;
    height: 25%;
`;
const ReportFrame = styled.div`
    display: flex;
    justify-content: center;
    height: 65%;
`;

const BattleResultPage = () => {
    const [monsterReport, setMonsterReport] = useState<MonsterReportDTO[]>([]);
    const [battleReport, setBattleReport] = useState<BattleReportDTO[]>([]);
    const [isNowLoadingMonsterReport, setIsNowLoadingMonsterReport] = useState(false);
    const [isNowLoadingBattleReport, setIsNowLoadingBattleReport] = useState(false);
    const [sortType, setSortType] = useState("1");

    useCheckToken();

    return (
        <OutsideFrame>
            <OptionFrame>
                {/* 検索条件部 */}
                <CommonFrame styleObj={{width: "55%", margin: "20px 5px 0px 0px"}}>
                    <BattleReportControllerBlock setMonsterReport={setMonsterReport}
                                                 sortType={sortType}
                                                 setIsNowLoadingMonsterReport={setIsNowLoadingMonsterReport}/>
                </CommonFrame>

                {/* 検索条件部 */}
                <CommonFrame styleObj={{width: "35%", margin: "20px 0px 0px 5px"}}>
                    <MonsterReportControllerBlock setBattleReport={setBattleReport}
                                                  setIsNowLoadingBattleReport={setIsNowLoadingBattleReport}/>
                </CommonFrame>
            </OptionFrame>

            <ReportFrame>
                {/* レポート部 */}
                <CommonFrame styleObj={{width: "55%", height: "60vh", margin: "20px 5px 0px 0px"}}>
                    <MonsterReport monsterReport={monsterReport}
                                   setSortType={setSortType}
                                   isNowLoadingMonsterReport={isNowLoadingMonsterReport}/>
                 </CommonFrame>
                {/* レポート部 */}
                <CommonFrame styleObj={{width: "35%", height: "60vh", margin: "20px 0px 0px 5px"}}>
                    <BattleReportBlock battleReport={battleReport}
                                       isNowLoadingBattleReport={isNowLoadingBattleReport} />
                </CommonFrame>
            </ReportFrame>
        </OutsideFrame>
    );
}

export default BattleResultPage;