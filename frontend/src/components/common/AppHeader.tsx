import styled from "styled-components";
import Button from "./Button";
import { useNavigate } from "react-router-dom";
import { COLORS, KEYS, PREFIX, URLS, USER_ROLE } from "../../lib/Constants";
import { useEffect, useLayoutEffect, useState } from "react";
import { isEmpty } from "../../lib/CommonLogic";
import { useServerWithQuery } from "../../hooks/useHooksOfCommon";
import { useCheckLogin } from "../../hooks/useHooksOfIndex";
import { UserDTO } from "../../types/UserManage";

const Simg = styled.img`
    width: 40px;
    height: 40px;
    margin-right: 10px;
    border-radius: 100%;
`;
const Sheader = styled.header`
    width: 100%;
    height: 60px;
    margin-bottom: 20px;
    box-shadow: ${COLORS.SHADOW};
    text-align: left;
    display: flex;
    align-items: center;
    justify-content: space-between;
    position: relative;
    top: 0x;
    background-color: white;
    z-index: 2000;
`;
const SdivButtonFrame = styled.div`
    display: flex;
    margin-right: auto;
    margin-left: 10px;
`;
const Sh1 = styled.h1`
    margin: 20px;
    color: ${COLORS.ACCENT_FONT_PINK};
`;

const Sspan = styled.span`
    transform: translateY(3px);
    margin-right: 10px;
`;

interface ArgProps { title: string; }

const AppHeader = ({title}: ArgProps) => {
    // ユーザー関係
    const [loginId, setLoginId] = useState<string | null>("");
    const [loginUser, setLoginUser] = useState<UserDTO | null>();
    const [isAdmin, setIsAdmin] = useState(false);
    const authorizedPerson: number[] = [USER_ROLE.ADMIN, USER_ROLE.SUPER_ADMIN];
    const [userImage, setUserImage] = useState<string>("");
    const [validToken, setValidToken] = useState(false);

    const navigate = useNavigate();
    const currentUrl: string = window.location.href;
    const isRootPage: boolean = currentUrl.endsWith("/"); // 最初のページ
    const isRootPage2: boolean = currentUrl.endsWith("/IndexPage");

    useCheckLogin(setValidToken);

    // ユーザー情報
    useLayoutEffect(() => {
        const id: string | null = localStorage.getItem(KEYS.USER_ID);
        const role: string | null = localStorage.getItem(KEYS.USER_ROLE);
        setLoginId(id);
        setIsAdmin(authorizedPerson.includes(parseInt(role!)));
    }, [loginId]);

    // 表示名取得
    const goToServer = useServerWithQuery();
    useEffect(() => {
        const selectName = async () => {
            const user: UserDTO | null = await goToServer(`${URLS.SELECT_LOGIN_USER}?loginId=${loginId}`);
            setLoginUser(user);
            if (user) setUserImage(PREFIX.BASE64 + user.UserImage);
        }
        selectName();
    }, [loginId, loginUser]);

    return (
        <Sheader>
            <Sh1>{title}</Sh1>
            <SdivButtonFrame style={{
                display: isRootPage || isRootPage2 ? "none" : ""
                }}>
                <Button text="モンスター闘技場"
                        width={125}
                        onClick={() => navigate("/BattlePage")}
                        disabled={!validToken}/>
                <Button text="闘技場戦績"
                        width={90}
                        onClick={() => navigate("/BattleResultPage")}
                        disabled={!validToken}/>
                <Button text="ユーザーページ"
                        width={120}
                        onClick={() => navigate("/UserPage")}
                        disabled={!validToken}/>
                <Button text="ショップ"
                        width={80}
                        onClick={() => navigate("/ShopPage")}
                        disabled={!validToken}/>
                <Button text="設定"
                        width={60}
                        onClick={() => navigate("/EditPage")}
                        disabled={!validToken || !isAdmin}/>
                <Button text="ログイン"
                        width={80}
                        onClick={() => navigate("/LoginPage")}
                        disabled={!validToken}/>
            </SdivButtonFrame>

            <div style={{display: "flex", alignItems: "center"}}>
                {
                    !isEmpty(loginUser) && loginUser!.UserImage!.length > 50 ? <Simg src={userImage} alt="" /> : ""
                }
                {
                    !isEmpty(loginUser) ? <Sspan>ようこそ{loginUser?.DispName}さん</Sspan> : ""
                }
                <Button text="メニューへ" onClick={() => navigate("/")} styleObj={{
                    marginRight: "20px",
                    position: "relative",
                    zIndex: 5000
                }}
                />
            </div>
        </Sheader>
    );
};

export default AppHeader;
