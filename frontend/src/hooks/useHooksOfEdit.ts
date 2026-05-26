import { useCallback } from "react";
import { KEYS, URLS } from "../lib/Constants";
import { EditMonsterDTO } from "../types/Edit";
import { api } from "../lib/apiClient";

/**
 * 更新後のモンスターステータスを反映
 */
export const useRefreshMonsterStatus = () => {
    const refreshMonsterStatus = useCallback( async (
        setEditMonsters: React.Dispatch<React.SetStateAction<EditMonsterDTO[]>>
    ) => {
        const loginId: string | null = localStorage.getItem(KEYS.USER_ID);
        const monsters = await api.POST<EditMonsterDTO[]>(URLS.FETCH_EDIT_MONSTERS, loginId);

        setEditMonsters([...monsters!]);
    }, []);
    return refreshMonsterStatus;
}
