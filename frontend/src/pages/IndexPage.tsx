import styled from "styled-components";
import { useEffect, useState } from "react";
import { COLORS, KEYS, USER_ROLE } from "../lib/Constants";
import { useCheckLogin } from "../hooks/useHooksOfIndex";
import ToLoginPageBlock from "../components/indexPage/ToLoginPageBlock";
import ToUserPageBlock from "../components/indexPage/ToUserPageBlock";
import ToBattlePageBlock from "../components/indexPage/ToBattlePageBlock";
import ToBattleResultPageBlock from "../components/indexPage/ToBattleResultPageBlock";
import ToShopPageBlock from "../components/indexPage/ToShopPageBlock";
import ToEditPageBlock from "../components/indexPage/ToEditPageBlock";
import ToGuitarGalleryBlock from "../components/indexPage/ToGuitarGalleryBlock";
import ToCareerPageBlock from "../components/indexPage/ToCareerPageBlock";

const SdivLinkFrame = styled.div`
    width: 90%;
    margin: 0px auto 0px;
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
    cursor: "default",
}

const IndexPage = () => {
    const [usableSettings, setUsableSettings] = useState(false);
    const validToken = useCheckLogin();
    const authorizedPerson: number[] =
        [USER_ROLE.ADMIN, USER_ROLE.SUPER_ADMIN];

    /**
     * 設定メニューを使用できるか
     */
    useEffect(() => {
        const role = localStorage.getItem(KEYS.USER_ROLE);
        setUsableSettings(authorizedPerson.includes(Number.parseInt(role!)));
    }, []);

    return (
        <SdivLinkFrame>
            <div style={{display: "flex"}}>
                <SdivContentsFrame>
                    {/* ログインページ */}
                    <ToLoginPageBlock />
                </SdivContentsFrame>

                <SdivContentsFrame>
                    {/* ギターギャラリー */}
                    <ToGuitarGalleryBlock />
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
                    {/* 経歴書ページ */}
                    <ToCareerPageBlock/>
                </SdivContentsFrame>
            </div>

            <div style={{display: "flex"}}>
                <SdivContentsFrame>
                    {/* 戦闘レポートページ */}
                    <ToBattleResultPageBlock validToken={validToken}
                                             classOfAnime={classOfAnime}
                                             titleStyle={titleStyle}/>
                </SdivContentsFrame>

                <SdivContentsFrame>
                    {/* ショップページ */}
                    <ToShopPageBlock validToken={validToken}
                                     classOfAnime={classOfAnime}
                                     titleStyle={titleStyle}/>
                </SdivContentsFrame>


            </div>

            <div style={{display: "flex"}}>
                <SdivContentsFrame>
                    {/* ユーザーページ */}
                    <ToUserPageBlock validToken={validToken}
                                     classOfAnime={classOfAnime}
                                     titleStyle={titleStyle} />
                </SdivContentsFrame>

                <SdivContentsFrame>
                    {/* 設定ページ */}
                    <ToEditPageBlock isValid={validToken && usableSettings}
                                     classOfAnime={classOfAnime}
                                     titleStyle={titleStyle}/>
                </SdivContentsFrame>


            </div>

            <h3 style={{
                color: `${COLORS.ACCENT_FONT_PINK}`,
                background: "white",
                height: "30px",
                borderRadius: "15px",
                }}>
                &emsp;※対応環境：PC(Chrome, edge推奨) ~ iPad mini(横画面1024px)
            </h3>
        </SdivLinkFrame>
    );
};

export default IndexPage;
