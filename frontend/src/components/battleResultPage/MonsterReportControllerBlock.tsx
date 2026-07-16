import { ChangeEvent, useCallback, useState } from "react";
import { COLORS, URLS } from "../../lib/Constants";
import CommonButton from "../common/CommonButton";
import CommonFromToDate from "../common/CommonFromTo";
import { BattleReportDTO } from "../../types/BattleReport";
import styled from "styled-components";
import BattleScaleListBlock from "./BattleScaleListBlock";
import { api } from "../../lib/apiClient";
import useApiErrorHandler from "../../hooks/useApiErrorHandler";
import { ApiError } from "../../types/ApiError";
import { isEmpty } from "../../lib/CommonLogic";


const Title = styled.h1`
    font-size: 16px;
    color: ${COLORS.CAPTION_FONT};
    margin-top: 5px;
`;

interface ArgProps {
    setBattleReport: React.Dispatch<React.SetStateAction<BattleReportDTO[]>>;
    setIsNowLoadingBattleReport: React.Dispatch<React.SetStateAction<boolean>>;
}

const MonsterReportControllerBlock = (
    {setBattleReport, setIsNowLoadingBattleReport}: ArgProps
) => {
    const [from, setFrom]               = useState("");
    const [to, setTo]                   = useState("");
    const [battleScale, setBattleScale] = useState("0");
    const [disable, setDisable]         = useState(false);
    const errorHandler                  = useApiErrorHandler();

    /**
     * 戦闘規模の選択
     */
    const changeBattleScaleHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        setBattleScale(e.target.value);
    }
    /**
     * 戦闘毎のレポートを取得
     */
    const fetchBattleReportHandler = useCallback(async () => {
        try {
            setIsNowLoadingBattleReport(true);

            const battleReport = await api.POST<BattleReportDTO[]>(URLS.BATTLE_REPORTS, {
                battleScale: battleScale,
                from:        from,
                to:          to,
            });

            if (isEmpty(battleReport)) throw new ApiError(500, "Fetch battleReport failed ...")

            setBattleReport(battleReport!);
            setIsNowLoadingBattleReport(false);
        } catch (e) {
            console.log(e);
            errorHandler(e)
        }
    }, [battleScale, from, to]);

    return (
        <div style={{margin: "0 0 10px 20px"}}>
            <Title>戦闘結果</Title>
            <BattleScaleListBlock
                changeBattleScaleHandler={changeBattleScaleHandler}
            />
            <CommonFromToDate
                labelText="期間"
                setDisable={setDisable}
                from={from}
                setFrom={setFrom}
                to={to}
                setTo={setTo}
            />
            <br/>
            <div style={{textAlign: "end"}}>
                <CommonButton
                    text="検索"
                    onClick={fetchBattleReportHandler}
                    disabled={disable}
                    styleObj={{margin: "15px 15px 0 0"}}
                    />
            </div>
        </div>
    );
}

export default MonsterReportControllerBlock;