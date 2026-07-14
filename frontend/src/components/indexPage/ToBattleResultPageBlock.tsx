import { Link } from "react-router-dom";
import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import CommonMenuTitle from "../common/CommonMenuTitle";
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

const ToBattlePageBlock = ({validToken, classOfAnime, titleStyle}: ArgProps) => {
    return (
        <div>
            <SLink to={validToken ? "/BattleResultPage" : "/"} >
                <CommonMenuTitle title={"📜戦闘戦績レポート"}
                        className={validToken ? classOfAnime : ""}
                        styleObj={validToken ? {} : titleStyle}/>
            </SLink>

            <CommonFrame>
                <Description>
                    これまでの戦闘の結果を記録していますので、そのレポートを閲覧できます。<br/>
                    モンスター毎、戦闘毎のレポートがあります。
                </Description>
            </CommonFrame>
        </div>
    );
}

export default ToBattlePageBlock;