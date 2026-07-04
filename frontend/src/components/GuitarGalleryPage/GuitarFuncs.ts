import { isEmpty } from "../../lib/CommonLogic";
import { GUITAR } from "../../lib/Constants";
import { GuitarParams } from "../../types/Guitar";

export const createQueryParams = (gParams: GuitarParams): URLSearchParams => {
    const uParams = new URLSearchParams();

    if (gParams.makerCd !== 0)
        uParams.append("makerCd", gParams.makerCd.toString());

    if (!isEmpty(gParams.name))
        uParams.append("name", gParams.name);

    if (gParams.colorCd !== 0)
        uParams.append("colorCd", gParams.colorCd.toString());

    if (!isEmpty(gParams.series))
        uParams.append("series", gParams.series);

    if (gParams.bodyMaterialTopCd > 0)
        uParams.append("bodyMaterialTopCd", gParams.bodyMaterialTopCd.toString());

    if (gParams.bodyMaterialBackCd > 0)
        uParams.append("bodyMaterialBackCd", gParams.bodyMaterialBackCd.toString());

    // 例外対策で、マイナス値に意味を持たせている
    // -1: open price, -2: 未定義 -3: price parse error
    if (gParams.minPrice >= -3)
        uParams.append("minPrice", gParams.minPrice.toString());

    if (gParams.maxPrice >= -3)
        uParams.append("maxPrice", gParams.maxPrice.toString());

    if (!isEmpty(gParams.order))
        uParams.append("order", gParams.order);

    if (!isEmpty(gParams.sort))
        uParams.append("sort", gParams.sort);

    if (0 < gParams.page && gParams.page <= 50)
        uParams.append("page", gParams.page.toString());

    if (10 <= gParams.pageSize && gParams.pageSize <= 100)
        uParams.append("pageSize", gParams.pageSize.toString());

    return uParams;
}

export const parsePrice = (price: number | undefined): string => {
    if (price === undefined) return String(price);

    let result: string;

    if (price === GUITAR.OPEN_PRICE) {
        result = "OPEN PRICE";
    } else if (price <= GUITAR.UNDEFINED_PRICE) {
        result = "?????? 円";
    } else {
        result = price.toLocaleString("ja-JP") + " 円";
    }
    return result;
}