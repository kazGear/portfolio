import { Link } from "react-router-dom";
import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import MenuTitle from "../common/MenuTitle";
import OutSideFrame from "../common/OutSideFrame";

const Slink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT_COLOR};
`;
const SpDescription = styled.p`
    margin: 10px 10px 10px 10px;
`;

interface ArgProps {
    validToken: boolean;
    classOfAnime: string;
    titleStyle: {}
}

const ToUserPageBlock = ({validToken, classOfAnime, titleStyle}: ArgProps) => {
    return (
        <div>
            <Slink to={validToken ? "/UserPage" : ""}>
                <MenuTitle
                    title={"ユーザーページ"}
                    className={validToken ? classOfAnime : ""}
                    styleObj={validToken ? {} : titleStyle}
                />
            </Slink>

            <OutSideFrame>
                <SpDescription>
                    ユーザーの所持金、所持物などのユーザー情報を確認できます。
                    <br />
                    所持金が尽きた際には、自己破産も可能です（所持金のリセット）。
                </SpDescription>
            </OutSideFrame>
        </div>
    );
}

export default ToUserPageBlock;