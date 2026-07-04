import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import Select from "../common/Select";
import { ChangeEvent, useEffect } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
    materials:    Code[] | null;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SearchMaterialBack = ({guitarParams, materials, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeMaterialBackHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setBodyMaterialBackCd(Number(e.target.value));
    }

    // カラーを選択した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.bodyMaterialBackCd])

    return (
        <Select onChange={changeMaterialBackHandler} >
            <option value="-1">未選択</option>
            {
                materials?.map(material =>
                        <option key={material.code}
                                value={material.code}>
                            {material.name}
                        </option>
                        )
            }
        </Select>
    );
}
export default SearchMaterialBack;