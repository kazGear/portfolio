import MenuTitle from "../common/CommonMenuTitle";
import styled from "styled-components";
import { Link } from "react-router-dom";
import { COLORS } from "../../lib/Constants";
import CommonFrame from "../common/CommonFrame";

const Slink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT};
`;
const Description = styled.p`
    margin: 10px
`;

const classOfAnime: string = "noneAnimation";

const ToGuitarGalleryBlock = () => {
    return (
        <div>
            <Slink to={"/JobPage"}>
                <MenuTitle title={"案件検索"} className={classOfAnime} />
            </Slink>

            <CommonFrame>
                <Description>
                    様々なIT案件を検索できます。<br/><br/>
                    多様な条件から案件を探すことができます。<br/>
                    表示されるのは案件の概要ですが、
                    詳細を知りたい場合は案件カードをクリックすることで
                    案件の掲載元を参照できます。<br/><br/>
                    定期的に案件数は増加していきます。
                </Description>
            </CommonFrame>
        </div>
    );
};

export default ToGuitarGalleryBlock;