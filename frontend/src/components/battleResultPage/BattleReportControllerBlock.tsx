import styled from "styled-components";
import { COLORS, KEYS, URLS } from "../../lib/Constants";
import CommonSelect from "../common/CommonSelect";
import CommonButton from "../common/CommonButton";
import { useCallback, useState } from "react";
import { MonsterReportDTO } from "../../types/BattleReport";
import MonsterTypesListBlock from "./MonsterTypesListBlock";
import { api } from "../../lib/apiClient";

const Title = styled.h1`
    font-size: 16px;
    color: ${COLORS.CAPTION_FONT};
    margin-top: 5px;
`;

interface ArgProps {
    setMonsterReport: React.Dispatch<React.SetStateAction<MonsterReportDTO[]>>;
    sortType: string;
    setIsNowLoadingMonsterReport: React.Dispatch<React.SetStateAction<boolean>>;
}

const BattleReportControllerBlock = (
    {setMonsterReport, sortType, setIsNowLoadingMonsterReport}: ArgProps
) => {
    const [monsterTypeId, setMonsterTypeId] = useState("0");
    const [isAscOrder, setIsAscOrder] = useState(true);

    /**
     * ソート制御
     */
    const sortHandler = (e: React.ChangeEvent<HTMLSelectElement>) => {
        if (e.target.value === KEYS.ORDER_BY_ASC) {
            setIsAscOrder(true);
        } else {
            setIsAscOrder(false);
        }
    }
    /**
     * モンスター毎のレポートを取得
     */
    const fetchMonsterReportHandler = useCallback(async () => {
        setIsNowLoadingMonsterReport(true);

        const monsterReport = await api.POST<MonsterReportDTO[]>(URLS.MONSTER_REPORTS, {
            monsterTypeId: monsterTypeId,
            sortType:      sortType,
            isAscOrder:    isAscOrder,
        });

        setMonsterReport(monsterReport!);
        setIsNowLoadingMonsterReport(false);
    }, [monsterTypeId, sortType, isAscOrder]);

    return (
        <div style={{margin: "0 0 0 20px"}}>
            <Title>モンスター戦績</Title>
            <MonsterTypesListBlock setMonsterTypeId={setMonsterTypeId} />
            <CommonSelect title="ソート順" onChange={sortHandler}>
                <option value={KEYS.ORDER_BY_ASC}>昇順</option>
                <option value={KEYS.ORDER_BY_DESC}>降順</option>
            </CommonSelect>
            <div style={{textAlign: "end"}}>
                <CommonButton
                    text="検索"
                    onClick={fetchMonsterReportHandler}
                    styleObj={{margin: "0 15px 15px 0"}}
                />
            </div>
        </div>
    );
};

export default BattleReportControllerBlock;