import { useEffect } from "react";
import { api } from "../lib/apiClient";
import { URLS } from "../lib/Constants";

/**
 *  トークンが有効か確認
 */
export const useCheckToken = () => {
    useEffect(() => {
        const checkToken = async () => {
            try {
                await api.POST<Response>(URLS.CHECK_LOGIN_TOKEN);
            } catch (err) {
                // 期限切れ, ログイン失敗
                globalThis.location.href = "/LoginPage";
            }
        }
        checkToken();
    }, []);
}