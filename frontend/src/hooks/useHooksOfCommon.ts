import {useCallback } from "react";

/**
 * サーバーと通信する（リクエストのボディ使用）
 */
export const useServerWithJson = () => {
    // モンスターたちの行動 ターン進行
    const useServerWithJson = useCallback(async (paramsJson: any, url: string) => {
        // json形式で大量のパラメータ送信
        const option: {} = {
            method: "POST",
            mode: "cors",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(paramsJson),
        };
        // モンスターたちの一連の行動
        try {
            const res = await fetch(url, option);
            const result = await res.json();
            return result;
        } catch (err) {
            console.error("サーバー通信に失敗しました。");
            console.error(err);
        }
    }, []);

    return useServerWithJson;
}

/**
 * サーバーと通信を行う（クエリパラメータを使用）
 */
export const useServerWithQuery = () => {
     const useServerWithQuery = useCallback(async (urlWithQuery: string) => {
        try {
            const option: {} = {
                method: "POST",
                mode: "cors",
                headers: { "Content-Type": "application/json" },
            };
            // urlパラメータで送信
            const response = await fetch(
                urlWithQuery, option
            );
            const result = await response.json();
            return result;
        } catch (err) {
            console.error("サーバー通信に失敗しました。");
            console.error(err);
        }
    }, []);

    return useServerWithQuery;
};

/**
 * サーバーと通信を行う
 */
export const useServer = () => {
    const useServerWithQuery = useCallback(async (url: string) => {
       try {
           const option: {} = {
               method: "GET",
               mode: "cors",
               headers: { "Content-Type": "application/json" },
           };
           // urlパラメータで送信
           const response = await fetch(
               url, option
           );
           const result = await response.json();
           return result;
       } catch (err) {
           console.error("サーバー通信に失敗しました。");
           console.error(err);
       }
   }, []);

   return useServerWithQuery;
};