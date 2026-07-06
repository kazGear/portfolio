import { useEffect } from "react";
import { GuitarParams } from "../../types/Guitar";
import Input from "../common/Input";

interface ArgProps {
    guitarParams: GuitarParams;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SelectorPageSize = ({guitarParams, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changePageSizeHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setPageSize(Number(e.currentTarget.value));
    }

    // ページサイズを設定した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.pageSize])

    return (
        <Input inputType="number"
               onBlur={changePageSizeHandler}
               placeholder=" (10 ~ 100) default 50"
               min="10"
               max="100"
               styleObj={{marginTop: "8px"}}/>
    );
}
export default SelectorPageSize;