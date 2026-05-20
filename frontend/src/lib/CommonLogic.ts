/** 要素が空であるか判定 */
export const isEmpty = (arg: any): boolean => {
    if (arg === null || arg === undefined) return true;
    if (Array.isArray(arg) && arg.length <= 0) return true;
    if (typeof arg === "string" && arg === "") return true;
    if (typeof arg === "string" && arg === "null") return true;
    return false;
}

/** 読み取り不可なJSONを読み取り可能にする */
export const jsonClone = (arg: any) => {
    return JSON.parse(JSON.stringify(arg));
}