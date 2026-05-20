import { useCallback } from "react";
import { URLS } from "../lib/Constants";
import { MetaDataDTO, MonsterDTO } from "../types/MonsterBattle";
import { isEmpty } from "../lib/CommonLogic";

interface ArgPropsRegistResult {
    monsters: MonsterDTO[];
    lastLog: MetaDataDTO | undefined;
    setResultLog: React.Dispatch<React.SetStateAction<MetaDataDTO | null>>;
    setShowResultDialog: React.Dispatch<React.SetStateAction<boolean>>;
    insertResult: (params: any, urls: string) => Promise<any>;
}
/**
 * 戦闘結果を記録する
 */
export const useRegistResult = () => {
    const registBattleResult = useCallback(({
        monsters,
        lastLog,
        setResultLog,
        setShowResultDialog,
        insertResult
    }: ArgPropsRegistResult) => {
        if (isEmpty(lastLog)) return;
        setResultLog(lastLog!);

        if (lastLog!.ExistWinner || lastLog!.AllLoser) {
            setShowResultDialog(true);
            // 戦績の記録
            insertResult(monsters, URLS.RECORD_BATTLE_RESULT);
        }
    }, []);

    return registBattleResult;
}
