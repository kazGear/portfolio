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

/** 年齢を算出、調整 (date: yyyy/mm/dd) */
export const calcAge = (date: string): number => {
    const birthDay = new Date(date);
    const today = new Date();

    const thisYearBirthDay = new Date(
        today.getFullYear(),
        birthDay.getMonth(),
        birthDay.getDate()
    );

    let age = today.getFullYear() - birthDay.getFullYear();
    if (today < thisYearBirthDay) age--;

    return age;
}