import { isEmpty } from "../../lib/CommonLogic";
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

    // 設計ミスにより０に意味を持たせてしまっている
    // コード設計を見直し、0は空ける。
    if (gParams.bodyMaterialTopCd > -1)
        uParams.append("bodyMaterialTopCd", gParams.bodyMaterialTopCd.toString());

    // 設計ミスにより０に意味を持たせてしまっている
    // コード設計を見直し、0は空ける。
    if (gParams.bodyMaterialBackCd > -1)
        uParams.append("bodyMaterialBackCd", gParams.bodyMaterialBackCd.toString());

    // 例外対策で、マイナス値に意味を持たせてしまっている
    // -1: price parse error, -2: open price -3: 未定義
    if (gParams.minPrice >= -3)
        uParams.append("minPrice", gParams.minPrice.toString());

    if (gParams.maxPrice > 0)
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