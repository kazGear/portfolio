import { Link } from "react-router-dom";
import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import CommonMenuTitle from "../common/CommonMenuTitle";
import CommonOutSideFrame from "../common/CommonFrame";

const SLink = styled(Link)`
    text-decoration: none;
    color: ${COLORS.MAIN_FONT_COLOR};
`;
const Description = styled.p`
    margin: 10px;
`;

interface ArgProps {
    validToken: boolean;
    classOfAnime: string;
    titleStyle: {}
}

const ToBattleResultPageBlock = ({validToken, classOfAnime, titleStyle}: ArgProps) => {
    return (
        <div>
            <SLink to={validToken ? "/BattlePage" : "/"} >
                <CommonMenuTitle title={"🐉モンスター闘技場"}
                           className={validToken ? classOfAnime : ""}
                           styleObj={validToken ? {} : titleStyle}/>
            </SLink>

            <CommonOutSideFrame>
                <Description>
                    某RPGカジノ風のモンスター闘技場です。どのモンスターが勝ち残るか当ててみてください。<br/>
                    ※バッチ処理でも毎晩強制的に戦わされています。
                </Description>
            </CommonOutSideFrame>
        </div>
    );
}

export default ToBattleResultPageBlock;