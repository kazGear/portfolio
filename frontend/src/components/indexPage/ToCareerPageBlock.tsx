import CommonMenuTitle from "../common/CommonMenuTitle";
import styled from "styled-components";
import { Link } from "react-router-dom";
import { COLORS } from "../../lib/Constants";
import CommonFrame from "../common/CommonFrame";

const SLink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT};
`;
const Description = styled.p`
    margin: 10px
`;

const classOfAnime: string = "noneAnimation";

const ToCareerPageBlock = () => {
    return (
        <div>
            <SLink to={"/CareerPage"}>
                <CommonMenuTitle title={"📖職務経歴書"} className={classOfAnime} />
            </SLink>

            <CommonFrame>
                <Description>
                    HTML版の職務経歴書です。<br/><br/>
                    Microsoft Word のようなシンプルな表示に切り替え可能です。
                </Description>
            </CommonFrame>
        </div>
    );
};

export default ToCareerPageBlock;