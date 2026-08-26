import styled from "styled-components";
import ErrorImage from "../../images/background/error2.jpg";
import { SIZE } from "../../lib/Constants";

const Frame = styled.div`
    width: 100%;
    height: calc(100vh - ${SIZE.HEADER_HEIGHT});
    overflow-y: hidden;
`;
const Div = styled.div`
    background: url(${ErrorImage});
    background-size: 100% 100%;
    height: 100%;
    overflow-y: hidden;
`;
const H1 = styled.h1`
    margin: 40px 0px 0px 40px;
`;

const CommonErrorView = () => {
    return (
        <Frame>
            <Div>
                <H1>想定外のエラーが発生しました。</H1>
                <h3 style={{margin: "10px 0px 0px 40px"}}>そんな日もありますよね(;^ω^)</h3>
            </Div>
        </Frame>
    );
}

export default CommonErrorView;