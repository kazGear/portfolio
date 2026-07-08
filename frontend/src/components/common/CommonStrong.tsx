import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const Span = styled.span`
    font-weight: 900;
    color: ${COLORS.CAPTION_FONT_COLOR};
`;

interface ArgProps {
    children: React.ReactNode;
}

const CommonStrong = ({children}: ArgProps) => {
    return (
        <Span>{children}</Span>
    );
}

export default CommonStrong;