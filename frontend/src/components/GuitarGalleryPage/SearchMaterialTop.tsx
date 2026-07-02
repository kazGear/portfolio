import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import Select from "../common/Select";
import { ChangeEvent } from "react";

interface ArgProps {
    guitarParams:  GuitarParams;
    materials:  Code[] | null;
}

const SearchMaterialTop = ({guitarParams, materials}: ArgProps) => {
    const gParams = guitarParams;

    const changeMaterialTopHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setBodyMaterialTopCd(Number(e.target.value));
    }

    return (
        <Select onChange={changeMaterialTopHandler} >
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
export default SearchMaterialTop;