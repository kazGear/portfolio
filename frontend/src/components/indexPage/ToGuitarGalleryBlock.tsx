import MenuTitle from "../common/CommonMenuTitle";
import styled from "styled-components";
import { Link } from "react-router-dom";
import { COLORS } from "../../lib/Constants";
import CommonFrame from "../common/CommonFrame";

const Slink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT_COLOR};
`;
const Description = styled.p`
    margin: 10px
`;

const classOfAnime: string = "noneAnimation";

const ToGuitarGalleryBlock = () => {
    return (
        <div>
            <Slink to={"/GuitarGalleryPage"}>
                <MenuTitle title={"🎸Guitar gallery"} className={classOfAnime} />
            </Slink>

            <CommonFrame>
                <Description>
                    様々なギターを眺めて楽しめます。<br/><br/>
                    ギター毎に、詳細スペックもご覧いただけます。<br/><br/>
                    情報を検索可能で、メーカー、シリーズ、カラー、ボディ材、価格帯などの条件で絞り込み検索が可能です。<br/>
                    ページネーションやソートにも対応。
                </Description>
            </CommonFrame>
        </div>
    );
};

export default ToGuitarGalleryBlock;