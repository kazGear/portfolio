import { GuitarParams } from "../../types/Guitar";
import Input from "../common/Input";

interface ArgProps {
    guitarParams: GuitarParams;
}

const SelectorPage = ({guitarParams}: ArgProps) => {
    const gParams = guitarParams;

    const changePageHandler = (e: React.FocusEvent<HTMLInputElement>) => {
        gParams.setPage(Number(e.currentTarget.value));
    }

    return (
        <Input inputType="number"
               onBlur={changePageHandler}
               placeholder="（1 ~ 50）"
               min="1"
               max="50"/>
    );
}
export default SelectorPage;