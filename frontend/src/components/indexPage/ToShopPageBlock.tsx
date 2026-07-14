import { Link } from "react-router-dom";
import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import MenuTitle from "../common/CommonMenuTitle";
import CommonFrame from "../common/CommonFrame";

const SLink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT};
`;
const Description = styled.p`
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
            <SLink to={validToken ? "/ShopPage" : ""} >
                <MenuTitle title={"🏠ショップ"}
                        className={validToken ? classOfAnime : ""}
                        styleObj={validToken ? {} : titleStyle}/>
            </SLink>

            <CommonFrame>
                <Description>
                    闘技場でのモンスターの使用権を購入できます。<br/>
                    その他、戦闘用背景なども追加予定です。
                </Description>
            </CommonFrame>
        </div>
    );
}

export default ToShopPageBlock;