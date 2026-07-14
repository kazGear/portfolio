import styled from "styled-components";
import ErrorImage from "../../images/background/error.png";

const Body = styled.body`
    width: 100%;
    height: 100%;
`;
const H1 = styled.h1`
    margin: 40px;
`;



const CommonErrorView = () => {
    return (
        <Body style={{background: `url(${ErrorImage})`, backgroundSize: "contain"}}>
            <H1>想定外のエラーが発生しました。</H1>
            <H1>そんな日もありますよね。</H1>
        </Body>
    );
}

export default CommonErrorView;