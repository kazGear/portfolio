import { useCallback } from "react";
import { KEYS, URLS } from "../lib/Constants";
import { EditMonsterDTO } from "../types/Edit";

/**
 * 更新後のモンスターステータスを反映
 */
export const useRefreshMonsterStatus = () => {
    const refreshMonsterStatus = useCallback( async (
        goToServer: any,
        setEditMonsters: React.Dispatch<React.SetStateAction<EditMonsterDTO[]>>
    ) => {
        const loginId: string | null = localStorage.getItem(KEYS.USER_ID);
        const monsters: EditMonsterDTO[] = await goToServer(
            URLS.FETCH_EDIT_MONSTERS + `?loginId=${loginId}`
        );
        setEditMonsters([...monsters]);
    }, []);
    return refreshMonsterStatus;
}
