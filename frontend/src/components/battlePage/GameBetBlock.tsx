import styled from "styled-components";
import Button from "../common/Button";
import React, { useLayoutEffect, useState } from "react";
import { MonsterDTO } from "../../types/MonsterBattle";
import { COLORS, KEYS, URLS } from "../../lib/Constants";
import { useServerWithQuery } from "../../hooks/useHooksOfCommon";
import { UserDTO } from "../../types/UserManage";
import MonsterSelectorBlock from "./MonsterSelectorBlock";


const SdivInputFrame = styled.div`
    display: flex;
    justify-content: space-between;
`;
const Sh1 = styled.h1`
    margin: 0 0 5px 0;
    color: ${COLORS.CAPTION_FONT_COLOR};
`;

interface ArgProps {
    monsters: MonsterDTO[];
    setBetMonster: React.Dispatch<React.SetStateAction<MonsterDTO | null>>;
    setBetGil: React.Dispatch<React.SetStateAction<number>>;
    setShowDialog: React.Dispatch<React.SetStateAction<boolean>>;
}

const GameBetContentsBlock = (
    {monsters, setBetMonster, setBetGil, setShowDialog}: ArgProps
) => {
    const [user, setUser] = useState<UserDTO | null>(null);
    const [cash, setCash] = useState("0");
    const [cashLimit, setCashLimit] = useState(0);

    // 選択したモンスターに賭ける
    const rowClickHandler = (row: any) => {
        setBetMonster(row);
    }

    // ユーザ情報
    const goToServer = useServerWithQuery();
    useLayoutEffect(() => {
        const selectUser = async () => {
            const user: UserDTO = await goToServer(`${URLS.USER_INFO}?loginId=${localStorage.getItem(KEYS.USER_ID)}`);
            setUser(user);
            setCash(user.Cash.toLocaleString());
            setCashLimit(user.Cash);
        }
        selectUser();
    }, []);

    // 掛け金検証
    const [betError, setBetError] = useState(false);
    const validBet = (e: any) => {
        if (e.target.value < 0 || cashLimit < e.target.value) {
            setBetError(true);
        } else if (!String(e.target.value).match(/^[1-9]\d*$/g)) {
            setBetError(true);
        } else {
            setBetError(false);
        }
    }
    // 賭けモンスター検証
    const [selectError, setSelectError] = useState(true);
    const validSelect = (e: any) => {
        if (e.target.value !== 0) {
             setSelectError(false);
        } else {
            setSelectError(true);
        }
    }

    return (
        <>
            <Sh1>どのモンスターに賭けますか？</Sh1>

            <MonsterSelectorBlock
                monsters={monsters}
                rowClickHandler={rowClickHandler}
                validSelect={validSelect}
                selectError={selectError}
            />

            <p style={{margin: "10px 0 0 0"}}>賭け金を入力してください</p>
            <span>所持金： {cash}&nbsp;Gil</span>
            <SdivInputFrame>
                <input
                    type="number"
                    style={{width: "100px", border: "1px solid black"}}
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                        validBet(e);
                        setBetGil(parseInt(e.target.value));
                    }}/ >
                <Button
                    text={"スタート"}
                    width={120}
                    onClick={() => setShowDialog(false)}
                    disabled={ betError || selectError} />
            </SdivInputFrame>

            {
                betError ? <span style={{color: "red"}}>
                                ※掛け金は1～{cash}Gilの範囲にしてください。
                           </span> : ""
            }
        </>
    );
}

export default GameBetContentsBlock;