import { Link } from "react-router-dom";
import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import MenuTitle from "../common/CommonMenuTitle";
import CommonFrame from "../common/CommonOutSideFrame";

const Slink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT_COLOR};
`;
const Description = styled.p`
    margin: 10px;
`;

interface ArgProps {
    isValid: boolean;
    classOfAnime: string;
    titleStyle: {}
}

const ToEditPageBlock = ({isValid, classOfAnime, titleStyle}: ArgProps) => {
    return (
        <div>
            <Slink to={isValid ? "/EditPage" : ""} >
                <MenuTitle title={"⚙各種設定"}
                           className={isValid ? classOfAnime : ""}
                           styleObj={isValid ? {} : titleStyle}/>
            </Slink>

            <CommonFrame>
                <Description>
                    モンスターステータス編集、モンスタースキル編集、使用モンスターの制限などが可能です。<br/>
                    ※管理ユーザのみ使用可能。
                </Description>
            </CommonFrame>
        </div>
    );
}

export default ToEditPageBlock;