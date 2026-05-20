import { useState } from "react";
import Button from "./Button";
import { KEYS, URLS } from "../../lib/Constants";
import Input from "./Input";

const prevImgStyle = {
    width: "90px",
    height: "90px",
    borderRadius: "100%"
};

interface ArgProps {
    styleObj?: React.CSSProperties;
}

const ImgUpload = ({styleObj}: ArgProps) => {
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
    const uploadImageHandler = async () => {
        if (imageFile) {
            const formData = new FormData();
            formData.append("image", imageFile);

            try {
                const res = await fetch(`${URLS.UPLOAD_IMAGE}?loginId=${loginId}`, {method: "POST", body: formData});

                if (res.ok) {
                    console.log('Image uploaded successfully');
                    window.location.href = "/UserPage"; // アップロード画像即時反映
                } else {
                    const errorText = await res.text();
                    console.error('Image upload failed', errorText);
                }
            } catch (err) {
                console.error('Error uploading image:', err);
            }
        }
    };

    return (
        <div>
            <Input labelTitle="画像登録"
                   inputType="file"
                   accept="image/*"
                   styleObj={{marginBottom: "6px", width: "108px"}}
                   onChange={selectImageHandler}/>
            {
                selectImage && (
                    <div style={{textAlign: "center"}}>
                        <img src={selectImage} alt="prevImage" style={prevImgStyle} />
                        <div style={{textAlign: "end"}}>
                            <Button text="upload" onClick={uploadImageHandler} styleObj={{margin: "10px 10px 0 0"}}/>
                        </div>
                    </div>
                )
            }
        </div>
    );
}

export default ImgUpload;