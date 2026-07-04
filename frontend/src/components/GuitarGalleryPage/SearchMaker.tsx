import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import Select from "../common/Select";
import { ChangeEvent, useEffect } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
    makers:       Code[] | null;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SearchMaker = ({guitarParams, makers, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeMakerHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setMakerCd(Number(e.target.value));
    }

    // メーカーを選択した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.makerCd])

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