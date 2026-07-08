import styled from "styled-components";
import { UserDTO } from "../types/UserManage";
import { KEYS, URLS } from "../lib/Constants";
import { MonsterDTO } from "../types/MonsterBattle";
import CashBlock from "../components/userPage/CashBlock";
import UserIdsBlock from "../components/userPage/UserIdsBlock";
import UserIconBlock from "../components/userPage/UserIconBlock";
import WinsBlock from "../components/userPage/WinsBlock";
import LossesBlock from "../components/userPage/LossesBlock";
import MonstersBlock from "../components/userPage/MonstersBlock";
import React, { useEffect, useState } from "react";
import CommonFrame from "../components/common/CommonOutSideFrame";
import CommonImgUpload from "../components/common/CommonImgUpload";
import { api } from "../lib/apiClient";
import { useCheckToken } from "../hooks/useHooksOfCommon";

const PageFrame = styled.div`
    display: flex;
    height: 100%;
`;
const PageL = styled.div`
    width: 40%;
    height: 100%;
    min-width: 440px;
`;
const PageR = styled.div`
    width: 57%;
    height: 85vh;
`;
const cashStyle: React.CSSProperties = {
    margin: "20px",
    height: "110px",
    minWidth: "100px",
    overflow: "hidden"
}
const monstersStyle: React.CSSProperties = {
    margin: "20px 20px 20px 0px",
    height: "100%",
    minWidth: "100px",
}
const iconStyle: React.CSSProperties = {
    margin: "20px",
    display: "flex",
    alignItems: "center",
    justifyContent: "space-around",
}
const winAndLoseStyle: React.CSSProperties = {
    margin: "20px",
    height: "80px"
}

const UserPage = () => {
    const [user, setUser] = useState<UserDTO | null>(null);
    const [monsters, setMonsters] = useState<MonsterDTO[] | null>([]);

    const loginId = localStorage.getItem(KEYS.USER_ID);

    useCheckToken();

    /**
     * ユーザ情報取得
     */
    useEffect(() => {
        const selectUser = async () => {
            const loginUser = await api.POST<UserDTO>(URLS.USER_INFO, loginId);
            const userMonsters = await api.POST<MonsterDTO[]>(URLS.MONSTERS_INFO, loginId);
            setUser(loginUser);
            setMonsters(userMonsters);
        }
        selectUser();
    }, []);

    return (
        <PageFrame>
            <PageL>
                <CommonFrame styleObj={iconStyle}>
                    <UserIconBlock user={user}  />
                    <UserIdsBlock user={user} />
                    <CommonImgUpload styleObj={{width: "110px"}} />
                </CommonFrame>
                <CommonFrame styleObj={cashStyle}>
                    <CashBlock user={user} />
                </CommonFrame>
                <CommonFrame styleObj={winAndLoseStyle}>
                    <WinsBlock user={user} />
                </CommonFrame>
                <CommonFrame styleObj={winAndLoseStyle}>
                    <LossesBlock user={user}/>
                </CommonFrame>
            </PageL>
            <PageR>
                <CommonFrame styleObj={monstersStyle}>
                    <MonstersBlock monsters={monsters} loginId={loginId} />
                </CommonFrame>
            </PageR>
        </PageFrame>

    );
}

export default UserPage;