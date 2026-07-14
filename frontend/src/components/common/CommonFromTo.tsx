import styled from "styled-components";
import { COLORS, SIZE } from "../../lib/Constants";
import { useEffect, useState } from "react";
import { useCheckFromTo } from "../../hooks/useHooksOfBattleResult";

const FromToFrame = styled.div`
    margin: 10px;
    display: inline;
`;
const Input = styled.input`
    font-family: cursive;
    color: ${COLORS.MAIN_FONT};
    height: ${SIZE.INPUT_HEIGHT};
    width: 100px;
    border: 1px solid ${COLORS.BORDER};
    box-shadow: ${COLORS.DIALOG_SHADOW} 1px 1px;
`;
const AlertMessage = styled.p`
    color: ${COLORS.ALERT_MESSAGE};
    display: inline;
    margin-left: 5px;
    font-size: 12px;
`;

interface ArgProps {
    labelText: string;
    setDisable: React.Dispatch<React.SetStateAction<boolean>>;
    from: string;
    to: string;
    setFrom: React.Dispatch<React.SetStateAction<string>>;
    setTo: React.Dispatch<React.SetStateAction<string>>;
}

const CommonFromTo = (
    {labelText, setDisable, from, to, setFrom, setTo}: ArgProps
) => {
    const [invalid, setInvalid] = useState(false);

    const setFromHandler = (e:React.ChangeEvent<HTMLInputElement>) => {
        setFrom(e.target.value);
    }
    const setToHandler = (e:React.ChangeEvent<HTMLInputElement>) => {
        setTo(e.target.value);
    }

    const chechFromTo = useCheckFromTo();
    useEffect(() => {
        chechFromTo({from, to, setInvalid, setDisable});
    }, [from, to]);

    return (
        <FromToFrame>
            <label style={{paddingRight: "10px"}}>{labelText}</label>
            <Input type="date"
                    onChange={setFromHandler}
                    />
            <span> ~ </span>
            <Input type="date"
                    onChange={setToHandler}
                    />
            {invalid ?
                <AlertMessage>
                    日付の設定が無効です。
                </AlertMessage>
                : ""
            }
        </FromToFrame>
    );
}

export default CommonFromTo;