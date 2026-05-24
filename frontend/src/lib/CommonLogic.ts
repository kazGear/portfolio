import { api } from "./apiClient";
import { URLS } from "./Constants";

/** 要素が空であるか判定 */
export const isEmpty = (arg: any): boolean => {
    if (arg === null || arg === undefined) return true;
    if (Array.isArray(arg) && arg.length <= 0) return true;
    if (typeof arg === "string" && arg === "") return true;
    if (typeof arg === "string" && arg === "null") return true;
    if (Object.keys(arg).length === 0) return true;

    return false;
}

/** 読み取り不可なJSONを読み取り可能にする */
export const jsonClone = (arg: any) => {
    return JSON.parse(JSON.stringify(arg));
}

/**
 *  トークンが有効か確認
 */
export const checkToken = () => {
    const checkToken = async () => {
        try {
            await api.POST<Response>(URLS.CHECK_LOGIN_TOKEN);
        } catch (err) {
            // 期限切れ, ログイン失敗
            window.location.href = "/LoginPage";
        }
    }
    checkToken();
}