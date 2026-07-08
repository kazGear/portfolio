import { UserDTO } from "../../types/UserManage";
import CommonStrong from "../common/CommonStrong";
import styled from "styled-components";

const SdivIDFrame = styled.div`
    height: 100px;
    min-width: 100px;
    margin-bottom: 30px;
`;

interface ArgProps {
    user: UserDTO | null;
}

const UserImageBlock = ({user}: ArgProps) => {
    return (
        <SdivIDFrame>
            <p>
                <CommonStrong>ログインID</CommonStrong>
                <br/>
                {user != null ? user.LoginId : ""}
            </p>
            <p>
                <CommonStrong>ロール</CommonStrong>
                <br/>
                {user != null ? user.RoleName : ""}
            </p>
        </SdivIDFrame>
    );
}

export default UserImageBlock;