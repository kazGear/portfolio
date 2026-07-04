import { useEffect } from "react";
import { GuitarParams } from "../../types/Guitar";
import Input from "../common/Input";

interface ArgProps {
    guitarParams: GuitarParams;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SelectorPage = ({guitarParams, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changePageHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setPage(Number(e.currentTarget.value));
    }

    // ページを変更した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.page])

    return (
        <Input inputType="number"
               onBlur={changePageHandler}
               placeholder=" (1 ~ 50) default 1"
               min="1"
               max="50"/>
    );
}
export default SelectorPage;