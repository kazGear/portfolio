import styled from "styled-components";
import { useEffect, useState } from "react";
import { UserDTO } from "../../types/UserManage";
import { PREFIX } from "../../lib/Constants";
import nowLoading from "../../images/background/nowLoading.gif";
import { isEmpty } from "../../lib/CommonLogic";
import CommonNowLoading from "../common/CommonNowLoading";

const ImageFrame = styled.div`
    height: 150px;
    min-width: 120px;
    margin: 20px;
    align-content: center;
`;
const Img = styled.img`
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
        <ImageFrame>
            { userImage.length < 50 ? ( // base64 prefixは30文字無い程度
                    <CommonNowLoading alt="ユーザーイメージ" size="80pxs"/>
                ) : (
                    <Img src={userImage} alt="ユーザーイメージ"/>
                )
            }
        </ImageFrame>
    );
}

export default UserIconBlock;