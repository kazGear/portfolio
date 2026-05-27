import { KEYS } from "./Constants";

interface FetchOptions {
    method?: "GET" | "POST" | "PUT" | "DELETE";
    headers?: { [key: string]: string };
    body?: any;
}

/** APIへのアクセス基盤 */
export async function apiClient<T>(
    endpoint: string,
    options:  FetchOptions = {},
): Promise<T | null> {

    // props破壊的変更の回避
    let fetchOptions: FetchOptions = {};
    if (options.body instanceof FormData) {
        fetchOptions = {
            method:  options.method,
            headers: structuredClone(options.headers),
            body:    options.body,
        }
    } else {
        fetchOptions = structuredClone(options);
    }
    fetchOptions.headers ??= {}

    // オプション構築
    if (fetchOptions.body instanceof FormData) {
        fetchOptions.headers = {
            ...fetchOptions.headers
        };
    } else {
        fetchOptions.body = JSON.stringify(fetchOptions.body);
        fetchOptions.headers = {
            "Content-Type": "application/json",
            ...fetchOptions.headers
        };
    }

    // 認証トークンチェック
    const token = localStorage.getItem(KEYS.TOKEN);
    if (token && fetchOptions.headers !== undefined)
        fetchOptions.headers["Authorization"] = `${token}`;

    // apiアクセス
    const res = await fetch(endpoint, { ...fetchOptions, });

    // エラー処理
    if (!res.ok) {
        let message = `API Error ${res.status}`;
        try {
            const data = await res.json();
            if (data?.message) message += ` ${data.message}`;
        } catch { /* JSON じゃない場合は無視 */ }

        throw new Error(message);
    }

    // apiのレスポンス次第でjsonParseが失敗する
    try {
        return await res.json();
    } catch {
        return null;
    }
}

/** httpMethodヘルパー */
export const api = {
    GET: <T>(url: string) => apiClient<T>(url, { method: "GET" }),

    POST: <T>(url: string, body?: any) => apiClient<T>(url, { method: "POST", body: body, }),

    PUT: <T>(url: string, body?: any) => apiClient<T>(url, { method: "PUT", body: body, }),

    DELETE: <T>(url: string, body?: any) => apiClient<T>(url, { method: "DELETE", body: body, }),
};
