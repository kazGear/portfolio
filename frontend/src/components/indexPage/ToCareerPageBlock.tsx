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

const ToCareerPageBlock = () => {
    return (
        <div>
            <Slink to={"/CareerPage"}>
                <MenuTitle title={"📖職務経歴書"} className={classOfAnime} />
            </Slink>

            <OutSideFrame>
                <SpDescription>
                    HTML版の職務経歴書です。<br/><br/>
                    Microsoft Word のようなシンプルな表示に切り替え可能です。
                </SpDescription>
            </OutSideFrame>
        </div>
    );
};

export default ToCareerPageBlock;