import { useCallback, useLayoutEffect, useState } from "react";
import { useServerWithJson, useServerWithQuery } from "../../../hooks/useHooksOfCommon";
import { KEYS, URLS } from "../../../lib/Constants";
import { EditSkillsDTO } from "../../../types/Edit";
import EditSkillBlock from "./EditSkillBlock";
import HeaderOfBody from "../../common/HeaderOfBody";
import SkillUpdateDialog from "./SkillUpdateDialog";
import { isNotEmittedStatement } from "typescript";

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
    const goToServerWithQuery = useServerWithQuery();
    useLayoutEffect(() => {
        const fetchEditSkills = async () => {
            const monsterSkills: EditSkillsDTO[] = await goToServerWithQuery(
                URLS.FETCH_EDIT_SKILLS + `?loginId=${localStorage.getItem(KEYS.USER_ID)}`
            );
            setEditMonsterSkills(monsterSkills);
            setIsNowLoading(false);
        }
        fetchEditSkills();
    }, []);
    /**
     * スキル更新
     */
    const goToServerWithJson = useServerWithJson();
    const updateSkillsHandler = useCallback(async () => {
         await goToServerWithJson(
            editMonsterSkills, URLS.CHANGE_MONSTER_SKILLS
        );
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