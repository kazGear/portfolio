import { useCallback, useState } from "react";
import CommonButton from "./CommonButton";
import { KEYS, URLS } from "../../lib/Constants";
import Input from "./CommonInput";
import { api, apiClient } from "../../lib/apiClient";

const prevImgStyle = {
    width: "90px",
    height: "90px",
    borderRadius: "100%"
};

interface ArgProps {
    styleObj?: React.CSSProperties;
}

const CommonImgUpload = ({styleObj}: ArgProps) => {
    const [selectImage, setSelectImage] = useState<string | null>(null);
    const [imageFile, setImagefile] = useState<File | null>(null);

    const loginId: string | null = localStorage.getItem(KEYS.USER_ID);
    /**
     * 自身のサムネ画像を選択
     */
    const selectImageHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
         if (e.target.files && e.target.files[0]) {
             setSelectImage(URL.createObjectURL(e.target.files[0]));
             setImagefile(e.target.files[0]);
         }
    };
    /**
     * 選択した画像を自身のサムネとして登録
     */
    const uploadImageHandler = useCallback(() => {
        const uploadImage = async () => {
            if (imageFile) {
                const formData = new FormData();
                formData.append("image", imageFile);
                formData.append("loginId", loginId ?? "");

                try {
                    await apiClient(URLS.UPLOAD_IMAGE, { body: formData, method: "PUT" });
                    globalThis.location.href = "/UserPage"; // アップロード画像即時反映
                } catch (err) {
                    alert('Error image upload:' + err);
                }
            }
        };
        uploadImage();
    }, [imageFile, loginId]);

    return (
        <div>
            <Input labelTitle="画像登録："
                   inputType="file"
                   accept="image/*"
                   styleObj={{marginBottom: "6px", width: "108px"}}
                   onChange={selectImageHandler}/>
            {
                selectImage && (
                    <div style={{textAlign: "center"}}>
                        <img src={selectImage} alt="prevImage" style={prevImgStyle} />
                        <div style={{textAlign: "end"}}>
                            <CommonButton text="upload"
                                    onClick={uploadImageHandler}
                                    styleObj={{margin: "10px 10px 0 0"}} />
                        </div>
                    </div>
                )
            }
        </div>
    );
}

export default CommonImgUpload;