import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import CommonSelect from "../common/CommonSelect";
import { ChangeEvent, useEffect } from "react";

interface ArgProps {
    guitarParams: GuitarParams;
    colors:       Code[] | null;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SearchColor = ({guitarParams, colors, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeColorHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        gParams.setColorCd(Number(e.target.value));
    }

    // カラーを選択した時点で検索実行
    useEffect(() => {
        callback(gParams)
        gParams.setPage(1)
    }, [gParams.colorCd])

    return (
        <CommonSelect onChange={changeColorHandler} >
            <option value="0">未選択</option>
            {
                colors?.map(color =>
                        <option key={color.code}
                                value={color.code}>
                            {color.name}
                        </option>
                        )
            }
        </CommonSelect>
    );
}
export default SearchColor;