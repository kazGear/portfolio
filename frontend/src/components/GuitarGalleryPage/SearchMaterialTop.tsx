import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import CommonSelect from "../common/CommonSelect";
import { ChangeEvent, useEffect } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
    materials:    Code[] | null;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SearchMaterialTop = ({guitarParams, materials, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeMaterialTopHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setBodyMaterialTopCd(Number(e.target.value));
    }

    // 木材を選択した時点で検索実行
    useEffect(() => {
        callback(gParams)
        gParams.setPage(1)
    }, [gParams.bodyMaterialTopCd])

    return (
        <CommonSelect onChange={changeMaterialTopHandler} >
            <option value="-1">未選択</option>
            {
                materials?.map(material =>
                        <option key={material.code}
                                value={material.code}>
                            {material.name}
                        </option>
                        )
            }
        </CommonSelect>
    );
}
export default SearchMaterialTop;