import { useEffect } from "react";
import { GuitarParams } from "../../types/Guitar";
import CommonInput from "../common/CommonInput";

interface ArgProps {
    guitarParams: GuitarParams;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SearchName = ({guitarParams, callback}: ArgProps) => {
    const gParams = guitarParams;

    const changeNameHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setName(e.currentTarget.value);
    }

    // カラーを選択した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.name])

    return (
        <CommonInput inputType="text"
                     onBlur={changeNameHandler}
                     placeholder="（部分一致検索）"
                     styleObj={{marginTop: "2px"}}/>
    );
}
export default SearchName;