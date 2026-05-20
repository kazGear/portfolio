import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const Sspan = styled.span`
    font-weight: 900;
    color: ${COLORS.ACCENT_FONT_PINK};
`;

interface ArgProps {
    children: React.ReactNode;
    strongPing?: boolean;
}

const Accent = ({children}: ArgProps) => {
    return (
        <Sspan>{children}</Sspan>
    );
}

export default Accent;