import { useEffect } from "react";
import { CodeDTO } from "../../../types/Common";
import { EditMonsterDTO } from "../../../types/Edit";
import BorderTd from "../../common/CommonBorderTd";
import Select from "../../common/CommonSelect";
import { URLS } from "../../../lib/Constants";
import { api } from "../../../lib/apiClient";

interface ArgProps {
    weekDropDown: CodeDTO[];
    monster: EditMonsterDTO;
    setWeekDropDown: React.Dispatch<React.SetStateAction<CodeDTO[]>>;
}

const EditMonsterWeekBlock = (
    {weekDropDown, monster, setWeekDropDown}: ArgProps
) => {
    /**
     * 弱点ドロップダウン
     */
    useEffect(() => {
        const fetchWeekDropDown = async () => {
            const dropDown = await api.GET<CodeDTO[]>(URLS.FETCH_ELEMENT_CODE);
            setWeekDropDown(dropDown!);
        }
        fetchWeekDropDown();
    }, []);

    return (
        <>
            <BorderTd>{monster.WeekName}</BorderTd>
            <BorderTd>
                <Select styleObj={{width: "50px"}}
                        onChange={(e: React.ChangeEvent<HTMLSelectElement> | undefined) => {
                            monster.AfterWeek = Number.parseInt(e!.target.value);
                            monster.IsChanged = true;
                        }}>
                    <option value="0"></option>
                    {
                        weekDropDown.map((opt, index) => (
                            <option key={index} value={opt.Value}>{opt.Name}</option>
                        ))
                    }

                </Select>
            </BorderTd>
        </>
    );
}

export default EditMonsterWeekBlock;