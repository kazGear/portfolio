
import { GuitarParams } from "../../types/Guitar";
import CommonInput from "../common/CommonInput";

interface ArgProps {
    guitarParams: GuitarParams;
    styleObj?:    React.CSSProperties;
}

const SearchName = ({guitarParams, styleObj}: ArgProps) => {
    const gParams = guitarParams;

    const changeNameHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setName(e.currentTarget.value);
    }

    return (
        <CommonInput inputType="text"
                     onBlur={changeNameHandler}
                     placeholder="（部分一致検索）"
                     styleObj={styleObj}/>
    );
}
export default SearchName;