import styled from "styled-components";
import { useLayoutEffect, useState } from "react";
import { COLORS, KEYS, USER_ROLE } from "../lib/Constants";
import { useCheckLogin } from "../hooks/useHooksOfIndex";
import ToLoginPageBlock from "../components/indexPage/ToLoginPageBlock";
import ToUserPageBlock from "../components/indexPage/ToUserPageBlock";
import ToBattlePageBlock from "../components/indexPage/ToBattlePageBlock";
import ToBattleResultPageBlock from "../components/indexPage/ToBattleResultPageBlock";
import ToShopPageBlock from "../components/indexPage/ToShopPageBlock";
import ToEditPageBlock from "../components/indexPage/ToEditPageBlock";

const SdivLinkFrame = styled.div`
    width: 90%;
    margin: 0 auto;
`;
const SdivContentsFrame = styled.div`
    width: 50%;
    margin: 0 20px 0 20px;
`;

const fontColor: string = COLORS.MAIN_FONT_COLOR;
const backColor: string = COLORS.MENU_DISABLED;
const classOfAnime: string = "noneAnimation";
const titleStyle: {} = {
    color: fontColor,
    background: backColor,
}

const IndexPage = () => {
    const [validToken, setValidToken] = useState(false);
    const [usableSettings, setUsableSettings] = useState(false);
    const authorizedPerson: number[] = [USER_ROLE.ADMIN, USER_ROLE.SUPER_ADMIN];

    useCheckLogin(setValidToken);

    /**
     * 設定メニューを使用できるか
     */
    useLayoutEffect(() => {
        const role = localStorage.getItem(KEYS.USER_ROLE);
        setUsableSettings(authorizedPerson.includes(parseInt(role!)));
    }, []);

    return (
        <SdivLinkFrame>
            <div style={{display: "flex"}}>
                <SdivContentsFrame>
                    {/* ログインページ */}
                    <ToLoginPageBlock />
                </SdivContentsFrame>

                <SdivContentsFrame>
                    {/* ユーザーページ */}
                    <ToUserPageBlock validToken={validToken}
                                     classOfAnime={classOfAnime}
                                     titleStyle={titleStyle} />
                </SdivContentsFrame>
            </div>

            <div style={{display: "flex"}}>
                <SdivContentsFrame>
                    {/* 戦闘ページ */}
                    <ToBattlePageBlock validToken={validToken}
                                       classOfAnime={classOfAnime}
                                       titleStyle={titleStyle}/>
                </SdivContentsFrame>

                <SdivContentsFrame>
                    {/* 戦闘レポートページ */}
                    <ToBattleResultPageBlock validToken={validToken}
                                             classOfAnime={classOfAnime}
                                             titleStyle={titleStyle}/>
                </SdivContentsFrame>
            </div>

            <div style={{display: "flex"}}>
                <SdivContentsFrame>
                    {/* ショップページ */}
                    <ToShopPageBlock validToken={validToken}
                                     classOfAnime={classOfAnime}
                                     titleStyle={titleStyle}/>
                </SdivContentsFrame>

                <SdivContentsFrame>
                    {/* 設定ページ */}
                    <ToEditPageBlock isValid={validToken && usableSettings}
                                     classOfAnime={classOfAnime}
                                     titleStyle={titleStyle}/>
                </SdivContentsFrame>
            </div>

            <p style={{color: `${COLORS.ACCENT_FONT_PINK}`}}>※スマホ非対応、Chrome, edge推奨。</p>
        </SdivLinkFrame>
    );
};

export default IndexPage;
