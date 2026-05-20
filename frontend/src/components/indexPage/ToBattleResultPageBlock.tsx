import { Link } from "react-router-dom";
import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import MenuTitle from "../common/MenuTitle";
import OutSideFrame from "../common/OutSideFrame";

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

const ToBattlePageBlock = ({validToken, classOfAnime, titleStyle}: ArgProps) => {
    return (
        <div>
            <Slink to={validToken ? "/BattleResultPage" : "/"} >
                <MenuTitle title={"戦闘戦績レポート"}
                        className={validToken ? classOfAnime : ""}
                        styleObj={validToken ? {} : titleStyle}/>
            </Slink>

            <OutSideFrame>
                <SpDescription>
                    これまでの戦闘の結果を記録していますので、そのレポートを閲覧できます。<br/>
                    モンスター毎、戦闘毎のレポートがあります。
                </SpDescription>
            </OutSideFrame>
        </div>
    );
}

export default ToBattlePageBlock;