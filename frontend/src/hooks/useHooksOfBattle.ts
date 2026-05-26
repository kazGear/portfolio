import { useCallback } from "react";
import { URLS } from "../lib/Constants";
import { MetaDataDTO, MonsterDTO } from "../types/MonsterBattle";
import { isEmpty } from "../lib/CommonLogic";
import { api } from "../lib/apiClient";

interface ArgPropsRegistResult {
    monsters: MonsterDTO[];
    lastLog: MetaDataDTO | undefined;
    setResultLog: React.Dispatch<React.SetStateAction<MetaDataDTO | null>>;
    setShowResultDialog: React.Dispatch<React.SetStateAction<boolean>>;
}
/**
 * 戦闘結果を記録する
 */
export const useRegistResult = () => {
    const insertBattleResult = useCallback(({
        monsters,
        lastLog,
        setResultLog,
        setShowResultDialog,
    }: ArgPropsRegistResult) => {

        if (isEmpty(lastLog)) return;
        setResultLog(lastLog!);

        if (lastLog!.ExistWinner || lastLog!.AllLoser) {
            setShowResultDialog(true);
            // 戦績の記録
            api.POST<MonsterDTO[]>(URLS.RECORD_BATTLE_RESULT, monsters);
        }
    }, []);
    return insertBattleResult;
}
