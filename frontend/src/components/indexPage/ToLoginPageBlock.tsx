import MenuTitle from "../common/MenuTitle";
import styled from "styled-components";
import { Link } from "react-router-dom";
import { COLORS } from "../../lib/Constants";
import OutSideFrame from "../common/OutSideFrame";

const Slink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT_COLOR};
`;
const SpDescription = styled.p`
    margin: 10px
`;

const classOfAnime: string = "noneAnimation";

const ToLoginPageBlock = () => {
    return (
        <div>
            <Slink to={"/LoginPage"}>
                <MenuTitle title={"ログイン"} className={classOfAnime} />
            </Slink>

            <OutSideFrame>
                <SpDescription>
                    ログインすると、各種機能をご利用いただけます。
                    数日程は、ログイン状態が維持されます。ユーザ登録も可能です。
                    <br />
                    <br />
                    ※登録済ユーザー&emsp; ログインID：guest、パスワード：guest
                    <br />
                    ※テストユーザー&emsp; ログインID：test、パスワード：test
                    <br />
                    ※管理ユーザー&emsp; ログインID：admin、パスワード：admin
                </SpDescription>
            </OutSideFrame>
        </div>
    );
};

export default ToLoginPageBlock;
