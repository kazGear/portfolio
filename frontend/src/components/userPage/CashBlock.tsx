import styled from "styled-components";
import { UserDTO } from "../../types/UserManage";
import CommonButton from "../common/CommonButton";
import CommonStrong from "../common/CommonStrong";
import { KEYS, URLS } from "../../lib/Constants";
import { useCallback, useEffect, useState } from "react";
import { api } from "../../lib/apiClient";

const CashFrame = styled.div`
    height: 100px;
    margin-left: 25px;
`;
const SpanDanger = styled.span`
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
    useEffect(() => {
        setCash(user?.Cash ?? null);
        setBankruptcyCnt(user?.BankruptcyCnt ?? null);
    }, [user]);
    /**
     * 自己破産（所持金初期化）
     */
    const restartAsPlayer = useCallback(() => {
        const restart = async () => {
            const user: UserDTO | null = await api.PUT<UserDTO>(`${URLS.RESTART_AS_PLAYER}`, loginId);
            setCash(user?.Cash ?? null);
            setBankruptcyCnt(user?.BankruptcyCnt ?? null);
        }
        restart();
    }, [loginId]);

    return (
        <CashFrame>
            <p style={{margin:"20px 0 0 0"}}><CommonStrong>
                所持金</CommonStrong> : {cash != null ? cash.toLocaleString() : ""} Gil
            </p>
            <p style={{margin:0}}>
                <CommonStrong>自己破産</CommonStrong>（所持金初期化）
                <CommonButton text="自己破産 実行" width={120} onClick={restartAsPlayer} />
            </p>
            <p style={{margin:0}}>
                <CommonStrong>自己破産回数</CommonStrong>&nbsp;:&nbsp;
                <SpanDanger>{bankruptcyCnt != null ? bankruptcyCnt : ""} 回</SpanDanger>
            </p>
        </CashFrame>
    );
}

export default CashBlock;