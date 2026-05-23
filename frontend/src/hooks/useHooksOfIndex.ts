import { useLayoutEffect } from "react";
import { KEYS, URLS } from "../lib/Constants";
import { api } from "../lib/apiClient";

export const useCheckLogin = (
    setValidToken: React.Dispatch<React.SetStateAction<boolean>>
) => {
    useLayoutEffect(() => {
        const checkToken = async () => {
            try {
                await api.POST(URLS.CHECK_LOGIN_TOKEN);
                setValidToken(true);
            } catch (err) {
                localStorage.removeItem(KEYS.TOKEN);
                localStorage.removeItem(KEYS.USER_ID);
                localStorage.removeItem(KEYS.USER_ROLE);
                setValidToken(false);
            }
        }
        checkToken();
    }, []);
}