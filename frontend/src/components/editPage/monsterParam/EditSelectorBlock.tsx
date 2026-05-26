import { useEffect, useState } from "react";
import Select from "../../common/Select";
import { URLS } from "../../../lib/Constants";
import { CodeDTO } from "../../../types/Common";
import { api } from "../../../lib/apiClient";

interface ArgProps {
    setSelectEditType: React.Dispatch<React.SetStateAction<number>>
}

const EditSelectorBlock = ({setSelectEditType}: ArgProps) => {
    const [editTypeDropDown, setEditTypeDropDown] = useState<CodeDTO[]>([]);

    /**
     * 設定種類
     */
    useEffect(() => {
        const fetchDropDown = async () => {
            const dropDown = await api.GET<CodeDTO[]>(URLS.EDIT_INIT);
            setEditTypeDropDown(dropDown!);
        }
        fetchDropDown();
    }, []);

    return (
        <div>
            <Select title="設定種類"
                    onChange={(e: React.ChangeEvent<HTMLSelectElement>) => {
                        setSelectEditType(Number.parseInt(e.target.value))
                    }}>
            {
                editTypeDropDown.map((opt, index) => (
                    <option value={opt.Value} key={index}>{opt.Name}</option>
                ))
            }
            </Select>
        </div>
    );
}

export default EditSelectorBlock;