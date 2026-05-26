import { useCallback, useEffect, useState } from "react";
import { KEYS, URLS } from "../../../lib/Constants";
import { EditSkillsDTO } from "../../../types/Edit";
import EditSkillBlock from "./EditSkillBlock";
import HeaderOfBody from "../../common/HeaderOfBody";
import SkillUpdateDialog from "./SkillUpdateDialog"
import { api } from "../../../lib/apiClient";

interface ArgProps {
    editMonsterSkills: EditSkillsDTO[];
    setEditMonsterSkills: React.Dispatch<React.SetStateAction<EditSkillsDTO[]>>;
}

const EditMonsterSkillsBlock = ({editMonsterSkills, setEditMonsterSkills}: ArgProps) => {
    const [showUpdateDialog, setShowUpdateDialog] = useState(false);
    const [isNowLoading, setIsNowLoading] = useState<boolean>(true);
    /**
     * 編集賞モンスタースキル
     */
    useEffect(() => {
        const fetchEditSkills = async () => {
            const monsterSkills = await api.POST<EditSkillsDTO[]>(
                URLS.FETCH_EDIT_SKILLS, localStorage.getItem(KEYS.USER_ID)
            );
            setEditMonsterSkills(monsterSkills!);
            setIsNowLoading(false);
        }
        fetchEditSkills();
    }, []);
    /**
     * スキル更新
     */

    const updateSkillsHandler = useCallback(async () => {
         await api.PUT(URLS.CHANGE_MONSTER_SKILLS, editMonsterSkills);
    }, [editMonsterSkills]);

    return (
        <div>
            <HeaderOfBody message="変更したいスキルを選択してください。"
                          buttonText="スキル変更"
                          buttonWidth={120}
                          callback={() => {
                              setShowUpdateDialog(true);
                              updateSkillsHandler();
                          }}
                          />
            <EditSkillBlock editMonsterSkills={editMonsterSkills}
                            isNowLoading={isNowLoading}/>
            <SkillUpdateDialog showDialog={showUpdateDialog}
                               setShowUpdateDialog={setShowUpdateDialog}/>
        </div>
    );
}

export default EditMonsterSkillsBlock;