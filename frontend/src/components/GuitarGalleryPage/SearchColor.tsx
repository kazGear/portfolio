import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import Select from "../common/Select";
import { ChangeEvent } from "react";

interface ArgProps {
    guitarParams:        GuitarParams;
    colors:              Code[] | null;
}

const SearchColor = ({guitarParams, colors}: ArgProps) => {
    const gParams = guitarParams;

    const changeColorHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setColorCd(Number(e.target.value));
    }

    return (
        <Select onChange={changeColorHandler} >
            <option value="0">未選択</option>
            {
                colors?.map(color =>
                        <option key={color.code}
                                value={color.code}>
                            {color.name}
                        </option>
                        )
            }
        </Select>
    );
}
export default SearchColor;