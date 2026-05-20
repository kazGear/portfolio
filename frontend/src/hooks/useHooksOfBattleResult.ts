import { useCallback } from "react";
import { isEmpty } from "../lib/CommonLogic";

interface ArgPropsCheckFromTo {
    from: string;
    to: string;
    setInvalid: React.Dispatch<React.SetStateAction<boolean>>;
    setDisable: React.Dispatch<React.SetStateAction<boolean>>;
}
/**
 * from to の入力で矛盾がないか検証する
 */
export const useCheckFromTo = () => {
    const checkFromTo = useCallback(({from, to, setInvalid, setDisable}: ArgPropsCheckFromTo) => {
        // from or to のみの入力は有効
        if (isEmpty(from)) {
            setInvalid(false);
            setDisable(false);
            return;
        }
        if (isEmpty(to)) {
            setInvalid(false);
            setDisable(false);
            return;
        }
        // 過去、未来の日付が逆転していないか
        if (from > to) {
            setInvalid(true);
            setDisable(true);
        } else {
            setInvalid(false);
            setDisable(false);
        }
    }, []);

    return checkFromTo;
}