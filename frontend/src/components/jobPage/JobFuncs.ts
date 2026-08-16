import { isEmpty } from "../../lib/CommonLogic";
import { JobParams } from "../../types/Job";

export const createQueryParamsJob = (jobParams: JobParams): URLSearchParams => {
    const uParams = new URLSearchParams();

    if (!isEmpty(jobParams.title))
        uParams.append("title", jobParams.title);

    if (!isEmpty(jobParams.location))
        uParams.append("location", jobParams.location);

    if (!isEmpty(jobParams.workPlace))
        uParams.append("workPlace", jobParams.workPlace);

    if (jobParams.minSalaryAtMonthSpecifiedMin !== undefined) {
        if (jobParams.minSalaryAtMonthSpecifiedMin >= -3)
            uParams.append("minSalaryAtMonthSpecifiedMin", jobParams.minSalaryAtMonthSpecifiedMin.toString());
    }

    if (jobParams.minSalaryAtMonthSpecifiedMax !== undefined) {
        if (jobParams.minSalaryAtMonthSpecifiedMax >= -3)
            uParams.append("minSalaryAtMonthSpecifiedMax", jobParams.minSalaryAtMonthSpecifiedMax.toString());
    }

    if (jobParams.maxSalaryAtMonthSpecifiedMin !== undefined) {
        if (jobParams.maxSalaryAtMonthSpecifiedMin >= -3)
            uParams.append("maxSalaryAtMonthSpecifiedMin", jobParams.maxSalaryAtMonthSpecifiedMin.toString());
    }

    if (jobParams.maxSalaryAtMonthSpecifiedMax !== undefined) {
        if (jobParams.maxSalaryAtMonthSpecifiedMax >= -3)
            uParams.append("maxSalaryAtMonthSpecifiedMax", jobParams.maxSalaryAtMonthSpecifiedMax.toString());
    }

    if (!isEmpty(jobParams.sourceSite))
        uParams.append("sourceSite", jobParams.sourceSite);

    if (0 < jobParams.page && jobParams.page <= 50)
        uParams.append("page", jobParams.page.toString());

    if (10 <= jobParams.pageSize && jobParams.pageSize <= 100)
        uParams.append("pageSize", jobParams.pageSize.toString());

    // 複数の同一キーで登録した場合、フレームによっては配列として受け取れる
    jobParams.featureNames.forEach(feature => uParams.append("featureNames", feature));

    return uParams;
}

export const parsePrice = (price: number | undefined): string => {
    // if (price === undefined) return String(price);

    // let result: string;

    // if (price === GUITAR.OPEN_PRICE) {
    //     result = "OPEN PRICE";
    // } else if (price <= GUITAR.UNDEFINED_PRICE) {
    //     result = "???,??? 円";
    // } else {
    //     result = price.toLocaleString("ja-JP") + " 円"; // 3桁区切り
    // }
    // return result;
    return "";
}
