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

const WinsBlock = ({user}: ArgProps) => {
    return (
        <Frame>
            <p style={{margin: 0}}><CommonStrong>的中数</CommonStrong>
                &nbsp;:&nbsp;{user != null ? user!.Wins : ""} 回
            </p>
            <p style={{margin: 0}}><CommonStrong>配当金</CommonStrong>
                &nbsp;:&nbsp;{user != null ? user!.WinsGetCash : ""} Gil
            </p>
        </Frame>
    );
}

export default WinsBlock;