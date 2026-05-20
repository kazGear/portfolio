import styled from "styled-components";
import { useState } from "react";
import { BattleReportDTO, MonsterReportDTO } from "../types/BattleReport";
import MonsterReport from "../components/battleResultPage/MonsterReportBlock";
import { useCheckToken } from "../hooks/useHooksOfCommon";
import BattleReportControllerBlock from "../components/battleResultPage/BattleReportControllerBlock";
import BattleReportBlock from "../components/battleResultPage/BattleReportBlock";
import MonsterReportControllerBlock from "../components/battleResultPage/MonsterReportControllerBlock";
import OutSideFrame from "../components/common/OutSideFrame";

const SdivOutsideFrame = styled.div`
    height: 100%;
`;
const SdivOptionFrame = styled.div`
    display: flex;
    justify-content: center;
    height: 25%;
`;
const SdivReportFrame = styled.div`
    display: flex;
    justify-content: center;
    height: 65%;
`;

const BattleResultPage = () => {
    const [monsterReport, setMonsterReport] = useState<MonsterReportDTO[]>([]);
    const [battleReport, setBattleReport] = useState<BattleReportDTO[]>([]);
    const [isNowLoadingMonsterReport, setIsNowLoadingMonsterReport] = useState(false);
    const [isNowLoadingBattleReport, setIsNowLoadingBattleReport] = useState(false);
    const [sortType, setSortType] = useState("0");

    useCheckToken();

    return (
        <SdivOutsideFrame>
            <SdivOptionFrame>
                {/* 検索条件部 */}

                <OutSideFrame styleObj={{width: "55%", marginBottom: 0}}>
                    <BattleReportControllerBlock setMonsterReport={setMonsterReport}
                                                 sortType={sortType}
                                                 setIsNowLoadingMonsterReport={setIsNowLoadingMonsterReport}/>
                </OutSideFrame>

                {/* 検索条件部 */}
                <OutSideFrame styleObj={{width: "35%", marginBottom: 0}}>
                    <MonsterReportControllerBlock setBattleReport={setBattleReport}
                                                  setIsNowLoadingBattleReport={setIsNowLoadingBattleReport}/>
                </OutSideFrame>
            </SdivOptionFrame>

            <SdivReportFrame>
                {/* レポート部 */}
                <OutSideFrame styleObj={{width: "55%"}}>
                    <MonsterReport monsterReport={monsterReport}
                                   setSortType={setSortType}
                                   isNowLoadingMonsterReport={isNowLoadingMonsterReport}/>
                 </OutSideFrame>
                {/* レポート部 */}
                <OutSideFrame styleObj={{width: "35%"}}>
                    <BattleReportBlock battleReport={battleReport}
                                       isNowLoadingBattleReport={isNowLoadingBattleReport} />
                </OutSideFrame>
            </SdivReportFrame>
        </SdivOutsideFrame>
    );
}

export default BattleResultPage;