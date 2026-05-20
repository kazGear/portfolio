import { useLayoutEffect, useState } from "react";
import { useServerWithQuery } from "../../../hooks/useHooksOfCommon";
import OutSideFrame from "../../common/OutSideFrame";
import Select from "../../common/Select";
import { URLS } from "../../../lib/Constants";
import { CodeDTO } from "../../../types/Common";

interface ArgProps {
    setSelectEditType: React.Dispatch<React.SetStateAction<number>>
}

const EditSelectorBlock = ({setSelectEditType}: ArgProps) => {
    const [editTypeDropDown, setEditTypeDropDown] = useState<CodeDTO[]>([]);

    /**
     * 設定種類
     */
    const goToServer = useServerWithQuery();
    useLayoutEffect(() => {
        const fetchDropDown = async () => {
            const dropDown: CodeDTO[] = await goToServer(URLS.EDIT_INIT);
            setEditTypeDropDown(dropDown);
        }
        fetchDropDown();
    }, []);

    return (
        <div>
            <Select title="設定種類"
                    onChange={(e: React.ChangeEvent<HTMLSelectElement>) => {
                        setSelectEditType(parseInt(e.target.value))
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