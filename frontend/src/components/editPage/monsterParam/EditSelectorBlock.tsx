import { useEffect, useState } from "react";
import CommonSelect from "../../common/CommonSelect";
import { URLS } from "../../../lib/Constants";
import { CodeDTO } from "../../../types/Common";
import { api } from "../../../lib/apiClient";
import useApiErrorHandler from "../../../hooks/useApiErrorHandler";
import { ApiError } from "../../../types/ApiError";
import { isEmpty } from "../../../lib/CommonLogic";

interface ArgProps {
    setSelectEditType: React.Dispatch<React.SetStateAction<number>>
}

const EditSelectorBlock = ({setSelectEditType}: ArgProps) => {
    const [editTypeDropDown, setEditTypeDropDown] = useState<CodeDTO[]>([]);
    const errorHandler                            = useApiErrorHandler()

    /**
     * 設定種類
     */
    useEffect(() => {
        const fetchDropDown = async () => {
            try {
                const dropDown = await api.GET<CodeDTO[]>(URLS.EDIT_INIT);

                if (isEmpty(dropDown)) throw new ApiError(500, "Fetch dropDown failed ...")

                setEditTypeDropDown(dropDown!);
            } catch (e) {
                console.log(e);
                errorHandler(e);
            }
        }
        fetchDropDown();
    }, []);

    return (
        <div>
            <CommonSelect title="設定種類"
                    onChange={(e: React.ChangeEvent<HTMLSelectElement>) => {
                        setSelectEditType(Number.parseInt(e.target.value))
                    }}>
            {
                editTypeDropDown.map((opt, index) => (
                    <option value={opt.Value} key={index}>{opt.Name}</option>
                ))
            }
            </CommonSelect>
        </div>
    );
}

export default EditSelectorBlock;