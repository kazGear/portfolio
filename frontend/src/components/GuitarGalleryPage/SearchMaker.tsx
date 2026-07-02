import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import Select from "../common/Select";
import { ChangeEvent } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
    makers:       Code[] | null;
}

const SearchMaker = ({guitarParams, makers}: ArgProps) => {
    const gParams = guitarParams;

    const changeMakerHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setMakerCd(Number(e.target.value));
    }

    return (
        <Select onChange={changeMakerHandler} >
            <option value="0">未選択</option>
            {
                makers?.map(maker =>
                        <option key={maker.code}
                                value={maker.code}>
                            {maker.name}
                        </option>
                        )
            }
        </Select>
    );
}
export default SearchMaker;