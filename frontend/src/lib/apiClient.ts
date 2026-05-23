import { DOMAIN, KEYS } from "./Constants";

interface fetchOptions {
    method?: string;
    headers?: { [key: string]: string };
    body?: any;
};

/** APIへのアクセス基盤 */
export async function apiClient<T>(
    endpoint: string,
    options: fetchOptions = {}
): Promise<T | undefined> {
    // オプション構築
    const headers: Record<string, string> = {
        "Content-Type": "application/json",
        ...(options.headers || {}),
    };
    // トークンチェック
    const token = localStorage.getItem(KEYS.TOKEN);
    if (token) headers["Authorization"] = `${token}`;
    // アクセス
    const res = await fetch(endpoint, {
        ...options, headers,
    });
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
        return;
    }
}

/** httpMethodヘルパー */
export const api = {
    GET: <T>(url: string) =>
        apiClient<T>(url, { method: "GET" }),

    POST: <T>(url: string, body?: any) =>
        apiClient<T>(url, {
            method: "POST",
            body: JSON.stringify(body),
      }),

    PUT: <T>(url: string, body?: any) =>
        apiClient<T>(url, {
            method: "PUT",
            body: JSON.stringify(body),
        }),

    DELETE: <T>(url: string) =>
        apiClient<T>(url, { method: "DELETE" }),
};