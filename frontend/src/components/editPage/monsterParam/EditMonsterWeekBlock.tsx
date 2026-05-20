import { useLayoutEffect } from "react";
import { useServerWithQuery } from "../../../hooks/useHooksOfCommon";
import { CodeDTO } from "../../../types/Common";
import { EditMonsterDTO } from "../../../types/Edit";
import BorderTd from "../../common/BorderTd";
import Select from "../../common/Select";
import { URLS } from "../../../lib/Constants";

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
    const goToServer = useServerWithQuery();
    useLayoutEffect(() => {
        const fetchWeekDropDown = async () => {
            const dropDown: CodeDTO[] = await goToServer(
                URLS.FETCH_ELEMENT_CODE
            );
            setWeekDropDown(dropDown);
        }
        fetchWeekDropDown();
    }, []);

    return (
        <>
            <BorderTd>{monster.WeekName}</BorderTd>
            <BorderTd>
                <Select styleObj={{width: "50px"}}
                        onChange={(e: React.ChangeEvent<HTMLSelectElement> | undefined) => {
                            monster.AfterWeek = parseInt(e!.target.value);
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