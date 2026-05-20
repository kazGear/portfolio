import styled from "styled-components";
import { UserDTO } from "../../types/UserManage";
import Button from "../common/Button";
import Strong from "../common/Strong";
import { KEYS, URLS } from "../../lib/Constants";
import { useServerWithQuery } from "../../hooks/useHooksOfCommon";
import { useCallback, useLayoutEffect, useState } from "react";

const SdivCashFrame = styled.div`
    height: 100px;
    margin-left: 25px;
`;
const SspanDanger = styled.span`
    color: red;
    font-weight: bold;
`;

interface ArgProps {
    user: UserDTO | null;
}

const CashBlock = ({user}: ArgProps) => {
    const [cash, setCash] = useState<number | null>(user?.Cash ?? null);
    const [bankruptcyCnt, setBankruptcyCnt] = useState<number | null>(user?.BankruptcyCnt ?? null);

    const loginId: string | null = localStorage.getItem(KEYS.USER_ID);
    /**
     * 値更新
     */
    useLayoutEffect(() => {
        setCash(user?.Cash ?? null);
        setBankruptcyCnt(user?.BankruptcyCnt ?? null);
    }, [user]);
    /**
     * 自己破産（所持金初期化）
     */
    const goToServer = useServerWithQuery();
    const restartAsPlayer = useCallback(() => {
        const restart = async () => {
            const result: UserDTO | null = await goToServer(`${URLS.RESTART_AS_PLAYER}?loginId=${loginId}`);
            setCash(result?.Cash ?? null);
            setBankruptcyCnt(result?.BankruptcyCnt ?? null);
        }
        restart();
    }, [loginId]);

    return (
        <SdivCashFrame>
            <p style={{margin:"20px 0 0 0"}}><Strong>所持金</Strong> : {cash != null ? cash.toLocaleString() : ""} Gil</p>
            <p style={{margin:0}}><Strong>自己破産</Strong>（所持金初期化）<Button text="自己破産 実行" width={120} onClick={restartAsPlayer} /></p>
            <p style={{margin:0}}><Strong>自己破産回数</Strong> : <SspanDanger>{bankruptcyCnt != null ? bankruptcyCnt : ""} 回</SspanDanger></p>
        </SdivCashFrame>
    );
}

export default CashBlock;