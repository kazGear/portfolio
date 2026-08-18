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

const ToJobAnalyzeBlock = () => {
    return (
        <div>
            <Slink to={"/JobAnalyzePage"}>
                <MenuTitle title={"案件分析"} className={classOfAnime} />
            </Slink>

            <CommonFrame>
                <Description>
                    収集したITエンジニア案件を横断的に分析できる機能です。<br/><br/>
                    言語・フレームワーク・勤務地・リモート可否・給与などの特徴から、案件の傾向を可視化します。<br/><br/>
                    蓄積した大量の求人データを活用し市場情報を提供します。
                </Description>
            </CommonFrame>
        </div>
    );
};

export default ToJobAnalyzeBlock;