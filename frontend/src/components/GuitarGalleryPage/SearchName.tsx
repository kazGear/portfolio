import { GuitarParams } from "../../types/Guitar";
import Input from "../common/Input";

interface ArgProps {
    guitarParams: GuitarParams;
}

const SearchName = ({guitarParams}: ArgProps) => {
    const gParams = guitarParams;

    const changeNameHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setName(e.currentTarget.value);
    }

    return (
        <Input inputType="text" onBlur={changeNameHandler} placeholder="（部分一致検索）"/>
    );
}
export default SearchName;