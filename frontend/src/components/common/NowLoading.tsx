import styled from "styled-components";
import nowLoading from "../../images/background/nowLoading2.gif";

interface StyleProps {
    size?: string;
}

const Simg = styled.img<StyleProps>`
    width: ${(props) => props.size ?? "100px"};
    height: ${(props) => props.size ?? "100px"};
    border-radius: 100%;
`;

interface ArgProps {
    alt: string;
    size?: string;
    styleObj?: React.CSSProperties;
}

const NowLoading = ({alt, size, styleObj}: ArgProps) => {
    return <Simg src={nowLoading}
                 alt={alt}
                 size={size}
                 style={styleObj}/>;
}

export default NowLoading;