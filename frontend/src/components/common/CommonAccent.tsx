import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const Span = styled.span`
    font-weight: 900;
    color: ${COLORS.ACCENT_FONT_PINK};
`;

interface ArgProps {
    children: React.ReactNode;
    strongPing?: boolean;
}

const CommonAccent = ({children}: ArgProps) => {
    return (
        <Span>{children}</Span>
    );
}

export default CommonAccent;