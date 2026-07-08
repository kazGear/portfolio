import { Link } from "react-router-dom";
import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import MenuTitle from "../common/CommonMenuTitle";
import CommonFrame from "../common/CommonOutSideFrame";

const Slink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT_COLOR};
`;
const SpDescription = styled.p`
    margin: 10px;
`;

interface ArgProps {
    validToken: boolean;
    classOfAnime: string;
    titleStyle: {}
}

const ToShopPageBlock = ({validToken, classOfAnime, titleStyle}: ArgProps) => {
    return (
        <div>
            <Slink to={validToken ? "/ShopPage" : ""} >
                <MenuTitle title={"🏠ショップ"}
                        className={validToken ? classOfAnime : ""}
                        styleObj={validToken ? {} : titleStyle}/>
            </Slink>

            <CommonFrame>
                <SpDescription>
                    闘技場でのモンスターの使用権を購入できます。<br/>
                    その他、戦闘用背景なども追加予定です。
                </SpDescription>
            </CommonFrame>
        </div>
    );
}

export default ToShopPageBlock;