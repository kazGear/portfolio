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
                    案件検索ページ
                </Description>
            </CommonFrame>
        </div>
    );
};

export default ToGuitarGalleryBlock;