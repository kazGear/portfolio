import { Link } from "react-router-dom";
import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import MenuTitle from "../common/CommonMenuTitle";
import CommonFrame from "../common/CommonFrame";

const SLink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT_COLOR};
`;
const Description = styled.p`
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
            <SLink to={validToken ? "/UserPage" : ""}>
                <MenuTitle
                    title={"👦ユーザーページ"}
                    className={validToken ? classOfAnime : ""}
                    styleObj={validToken ? {} : titleStyle}
                />
            </SLink>

            <CommonFrame>
                <Description>
                    ユーザーの所持金、所持物などのユーザー情報を確認できます。
                    <br />
                    所持金が尽きた際には、自己破産も可能です（所持金のリセット）。
                </Description>
            </CommonFrame>
        </div>
    );
}

export default ToUserPageBlock;