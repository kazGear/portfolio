import { KEYS } from "./Constants";

interface fetchOptions {
    method?: "GET" | "POST" | "PUT" | "DELETE";
    headers?: { [key: string]: string };
    body?: any;
}

/** APIへのアクセス基盤 */
export async function apiClient<T>(
    endpoint: string,
    options: fetchOptions = {},
): Promise<T | null> {
    const fetchOptions = structuredClone(options);

    // オプション構築
    let headers: Record<string, string>;
    if (fetchOptions.body instanceof FormData) {
        headers = {
            ...fetchOptions.headers
        };
    } else {
        fetchOptions.body = JSON.stringify(fetchOptions.body)
        headers = {
            "Content-Type": "application/json",
            ...fetchOptions.headers
        };
    }

    // 認証トークンチェック
    const token = localStorage.getItem(KEYS.TOKEN);
    if (token && fetchOptions.headers !== null)
        fetchOptions.headers!["Authorization"] = `${token}`;

    // apiアクセス
    const res = await fetch(endpoint, { ...fetchOptions, headers, });

    // エラー処理
    if (!res.ok) {
        let message = `API Error ${res.status}`;
        try {
            const data = await res.json();
            if (data?.message) message += ` ${data.message}`;
        } catch {
            // JSON じゃない場合は無視
        }
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

    POST: <T>(url: string, body?: any) =>
        apiClient<T>(url, {
            method: "POST",
            body: body,
        }),

    PUT: <T>(url: string, body?: any) =>
        apiClient<T>(url, {
            method: "PUT",
            body: body,
        }),

    DELETE: <T>(url: string) => apiClient<T>(url, { method: "DELETE" }),
};
