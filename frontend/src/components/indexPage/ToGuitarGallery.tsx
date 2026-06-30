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

const ToGuitarGalleryBlock = () => {
    return (
        <div>
            <Slink to={"/GuitarGalleryPage"}>
                <MenuTitle title={"🎸Guitar gallery"} className={classOfAnime} />
            </Slink>

            <OutSideFrame>
                <SpDescription>
                    様々なギターを眺めて楽しめます。<br/><br/>
                    ギター毎に、詳細スペックもご覧いただけます。<br/><br/>
                    情報を検索可能で、メーカー、シリーズ、カラー、ボディ材、価格帯などの条件で絞り込み検索が可能です。<br/>
                    ページネーションやソートにも対応。
                </SpDescription>
            </OutSideFrame>
        </div>
    );
};

export default ToGuitarGalleryBlock;