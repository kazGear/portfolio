import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const Sspan = styled.span`
    font-weight: 900;
    color: ${COLORS.CAPTION_FONT_COLOR};
`;

interface ArgProps {
    children: React.ReactNode;
    strongPing?: boolean;
}

const Strong = ({children}: ArgProps) => {
    return (
        <Sspan>{children}</Sspan>
    );
}

export default Strong;