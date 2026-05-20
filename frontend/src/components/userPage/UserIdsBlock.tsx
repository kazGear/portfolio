import { UserDTO } from "../../types/UserManage";
import ImgUpload from "../common/ImgUpload";
import Strong from "../common/Strong";
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
                <Strong>ログインID</Strong>
                <br/>
                {user != null ? user.LoginId : ""}
            </p>
            <p>
                <Strong>ロール</Strong>
                <br/>
                {user != null ? user.RoleName : ""}
            </p>
        </SdivIDFrame>
    );
}

export default UserImageBlock;