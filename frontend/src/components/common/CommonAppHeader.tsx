import styled from "styled-components";
import CommonButton from "./CommonButton";
import { useNavigate } from "react-router-dom";
import { COLORS, KEYS, PREFIX, SIZE, URLS, USER_ROLE } from "../../lib/Constants";
import { useEffect, useState } from "react";
import { isEmpty } from "../../lib/CommonLogic";
import { useCheckLogin } from "../../hooks/useHooksOfIndex";
import { UserDTO } from "../../types/UserManage";
import { api } from "../../lib/apiClient";

const Img = styled.img`
    width: 40px;
    height: 40px;
    margin-right: 10px;
    border-radius: 100%;
`;
const Header = styled.header`
    width: 100%;
    height: ${SIZE.HEADER_HEIGHT};
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
    position: fixed;
`;
const CommonButtonFrame = styled.div`
    display: flex;
    margin-right: auto;
    margin-left: 0px;
`;
const H1 = styled.h1`
    margin: 20px;
    color: ${COLORS.ACCENT_FONT_PINK};
`;

const Span = styled.span`
    transform: translateY(3px);
    margin-right: 10px;
`;

interface ArgProps { title: string; }

const CommonAppHeader = ({title}: ArgProps) => {
    // ユーザー関係
    const [loginId, setLoginId] = useState<string | null>("");
    const [loginUser, setLoginUser] = useState<UserDTO | null>();
    const [isAdmin, setIsAdmin] = useState(false);
    const authorizedPerson: number[] = [USER_ROLE.ADMIN, USER_ROLE.SUPER_ADMIN];
    const [userImage, setUserImage] = useState<string>("");
    const validToken = useCheckLogin();
    const navigate = useNavigate();
    const currentUrl: string = globalThis.location.href;
    const isRootPage: boolean = currentUrl.endsWith("/"); // 最初のページ
    const isRootPage2: boolean = currentUrl.endsWith("/IndexPage");

    // ユーザー情報
    useEffect(() => {
        const id: string | null = localStorage.getItem(KEYS.USER_ID);
        const role: string | null = localStorage.getItem(KEYS.USER_ROLE);

        setLoginId(id);
        setIsAdmin(authorizedPerson.includes(Number.parseInt(role!)));
    }, [validToken]);

    // 表示名取得
    useEffect(() => {
        if (!loginId) return;

        const selectName = async () => {
            const user = await api.POST<UserDTO>(URLS.SELECT_LOGIN_USER, loginId);
            setLoginUser(user);

            if (user) setUserImage(PREFIX.BASE64 + user.UserImage);
        }
        selectName();
    }, [loginId]);

    return (
        <Header>
            <H1>{title}</H1>
            <CommonButtonFrame style={{
                display: isRootPage || isRootPage2 ? "none" : ""
                }}>
                <CommonButton text="モンスター闘技場"
                        width={125}
                        onClick={() => navigate("/BattlePage")}
                        disabled={!validToken}/>
                <CommonButton text="Guitar gallery"
                        width={120}
                        onClick={() => navigate("/GuitarGalleryPage")}
                        disabled={false}/>
                <CommonButton text="闘技場戦績"
                        width={90}
                        onClick={() => navigate("/BattleResultPage")}
                        disabled={!validToken}/>
                <CommonButton text="経歴書"
                        width={60}
                        onClick={() => navigate("/CareerPage")}
                        disabled={false}/>
                <CommonButton text="ユーザーページ"
                        width={120}
                        onClick={() => navigate("/UserPage")}
                        disabled={!validToken}/>
                <CommonButton text="ショップ"
                        width={80}
                        onClick={() => navigate("/ShopPage")}
                        disabled={!validToken}/>
                {/* <CommonButton text="設定"
                        width={60}
                        onClick={() => navigate("/EditPage")}
                        disabled={!validToken || !isAdmin}/> */}
            </CommonButtonFrame>

            <div style={{display: "flex", alignItems: "center"}}>
                {
                    !isEmpty(loginUser) && loginUser!.UserImage.length > 50 ? <Img src={userImage} alt="" /> : ""
                }
                {
                    !isEmpty(loginUser) ? <Span>ようこそ{loginUser?.DispName}さん</Span> : ""
                }
                <CommonButton text="メニューへ" onClick={() => navigate("/")} styleObj={{
                    marginRight: "20px",
                    position: "relative",
                    zIndex: 5000
                }}
                />
            </div>
        </Header>
    );
};

export default CommonAppHeader;
