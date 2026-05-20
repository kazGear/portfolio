import styled from "styled-components";
import { useEffect, useState } from "react";
import { UserDTO } from "../../types/UserManage";
import { PREFIX } from "../../lib/Constants";
import nowLoading from "../../images/background/nowLoading.gif";
import { isEmpty } from "../../lib/CommonLogic";
import NowLoading from "../common/NowLoading";

const SdivImageFrame = styled.div`
    height: 150px;
    min-width: 120px;
    margin: 20px 0px 20px 0px;
    align-content: center;
`;
const Simg = styled.img`
    widht: 120px;
    height: 120px;
    border-radius: 100%;
`;

interface ArgProps {
    user: UserDTO | null;
}

const UserIconBlock = ({user}: ArgProps) => {
    const [userImage, setUserImage] = useState("");
    /**
     * ユーザ情報取得
     */
    useEffect(() => {
        if (!isEmpty(user)) {
            const image: string | undefined = PREFIX.BASE64 + user!.UserImage;
            setUserImage(image);
        }
    }, [user, userImage]);

    return (
        <SdivImageFrame>
            { userImage.length < 50 ? ( // base64 prefixは30文字無い程度
                    <NowLoading alt="ユーザーイメージ" size="80pxs"/>
                ) : (
                    <Simg src={userImage} alt="ユーザーイメージ"/>
                )
            }
        </SdivImageFrame>
    );
}

export default UserIconBlock;