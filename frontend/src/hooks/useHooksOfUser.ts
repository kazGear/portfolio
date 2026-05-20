import { useCallback } from "react";
import { useServerWithQuery } from "./useHooksOfCommon";
import { KEYS, URLS } from "../lib/Constants";
import { isEmpty } from "../lib/CommonLogic";
import { useLayoutEffect } from "react";
import { UserDTO } from "../types/UserManage";

interface ArgPropsLogin {
    inputLoginId: string;
    inputPassword: string;
    setToken: React.Dispatch<React.SetStateAction<string | null>>;
    setShowAlert: React.Dispatch<React.SetStateAction<boolean>>;
}
/**
 * ユーザーのログイン処理
 */
export const useLogin = () => {
    const goToServer = useServerWithQuery();
    const callback = useCallback( async ({
        inputLoginId, inputPassword, setToken, setShowAlert}: ArgPropsLogin
    ) => {
        try {
            const user: UserDTO = await goToServer(
                `${URLS.LOGIN_USER}?loginId=${inputLoginId}&password=${inputPassword}`
            );
            // トークン有 >>> ログイン成功
            if (user.Token != null) {
                localStorage.setItem(KEYS.TOKEN, user.Token);
                localStorage.setItem(KEYS.USER_ID, inputLoginId);
                localStorage.setItem(KEYS.USER_ROLE, user.Role.toString());
                setToken(user.Token);
                setShowAlert(false);
                window.location.href = "/IndexPage";
            } else if (isEmpty(user.Token)) {
                localStorage.removeItem(KEYS.TOKEN);
                localStorage.removeItem(KEYS.USER_ID);
                localStorage.removeItem(KEYS.USER_ROLE);
                setShowAlert(true);
                setTimeout(() => window.location.href = "/LoginPage", 2000);
            }
        } catch (err) {
            localStorage.removeItem(KEYS.TOKEN);
            localStorage.removeItem(KEYS.USER_ID);
            localStorage.removeItem(KEYS.USER_ROLE);
            setShowAlert(true);
            setTimeout(() => window.location.href = "/LoginPage", 2000);
        }
    }, []);
    return callback;
}


interface ArgPropsForCreateUsedList {
    users: UserDTO[] | null;
    setUsedLoginIdList: React.Dispatch<React.SetStateAction<string[] | null>>;
    setUsedDispNameList: React.Dispatch<React.SetStateAction<string[] | null>>;
    setUsedDispShortNameList: React.Dispatch<React.SetStateAction<string[] | null>>;
}
/**
 * 既に登録済のログインID、表示名を取得する（検証で使用）
 */
export const useCreateUsedList = (
    {users, setUsedLoginIdList, setUsedDispNameList, setUsedDispShortNameList}: ArgPropsForCreateUsedList
) => {
    useLayoutEffect(() => {
        if (users !== null) {
            const usedLoginId: string[] = [];
            const usedDispName: string[] = [];
            const usedDispShortName: string[] = [];

            // すでに登録されているデータをまとめて保存
            users.forEach((user) => {
                usedLoginId.push(user.LoginId);
                usedDispName.push(user.DispName);
                usedDispShortName.push(user.DispShortName);
            });
            setUsedLoginIdList(usedLoginId);
            setUsedDispNameList(usedDispName);
            setUsedDispShortNameList(usedDispShortName);
        }
    }, [users]);
}