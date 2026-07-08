import styled from "styled-components";
import { UserDTO } from "../../types/UserManage";
import CommonStrong from "../common/CommonStrong";

const Frame = styled.div`
    height: 50%;
    margin: 20px;
`;

interface ArgProps {
    user: UserDTO | null;
}

const LossesBlock = ({user}: ArgProps) => {
    return (
        <Frame>
            <p style={{margin: 0}}><CommonStrong>敗北数</CommonStrong>
                &nbsp;:&nbsp;{user != null ? user!.Losses : ""} 回
            </p>
            <p style={{margin: 0}}><CommonStrong>損失</CommonStrong>
                &nbsp;:&nbsp;{user != null ? user!.LossesLostCash : ""} Gil
            </p>
        </Frame>
    );
}

export default LossesBlock;