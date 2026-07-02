import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import Select from "../common/Select";
import { ChangeEvent } from "react";

interface ArgProps {
    guitarParams:  GuitarParams;
    materials:  Code[] | null;
}

const SearchMaterialBack = ({guitarParams, materials}: ArgProps) => {
    const gParams = guitarParams;

    const changeMaterialBackHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setBodyMaterialBackCd(Number(e.target.value));
    }

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