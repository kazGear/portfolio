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
import OutSideFrame from "../components/common/OutSideFrame";
import ImgUpload from "../components/common/ImgUpload";
import { api } from "../lib/apiClient";
import { useCheckToken } from "../hooks/useHooksOfCommon";

const SdivPageFrame = styled.div`
    display: flex;
    height: 100%;
    margin-top: 60px;
`;
const SdivPageL = styled.div`
    width: 50%;
    height: 100%;
    min-width: 400px;
`;
const SdivPageR = styled.div`
    width: 50%;
    height: 85%;
    overflow: hidden;
`;
const cashStyle: React.CSSProperties = {
    margin: "20px",
    height: "110px",
    minWidth: "100px",
    overflow: "hidden"
}
const monstersStyle: React.CSSProperties = {
    margin: "20px",
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
        <SdivPageFrame>
            <SdivPageL>
                <OutSideFrame styleObj={iconStyle}>
                    <UserIconBlock user={user}  />
                    <UserIdsBlock user={user} />
                    <ImgUpload styleObj={{width: "110px"}} />
                </OutSideFrame>
                <OutSideFrame styleObj={cashStyle}>
                    <CashBlock user={user} />
                </OutSideFrame>
                <OutSideFrame styleObj={winAndLoseStyle}>
                    <WinsBlock user={user} />
                </OutSideFrame>
                <OutSideFrame styleObj={winAndLoseStyle}>
                    <LossesBlock user={user}/>
                </OutSideFrame>
            </SdivPageL>
            <SdivPageR>
                <OutSideFrame styleObj={monstersStyle}>
                    <MonstersBlock monsters={monsters} loginId={loginId} />
                </OutSideFrame>
            </SdivPageR>
        </SdivPageFrame>

    );
}

export default UserPage;