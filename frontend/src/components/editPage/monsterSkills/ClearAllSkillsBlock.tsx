import { useCallback, useState } from "react";
import CommonButton from "../../common/CommonButton";
import { URLS } from "../../../lib/Constants";
import CommonDialogFrame from "../../common/CommonDialogFrame";
import { EditSkillsDTO } from "../../../types/Edit";
import { api } from "../../../lib/apiClient";

interface ArgProps {
    selectEditType: number;
    setEditMonsterSkills:  React.Dispatch<React.SetStateAction<EditSkillsDTO[]>>;
}

const ClearAllSkillsBlock = ({selectEditType, setEditMonsterSkills}: ArgProps) => {
    const [showInitConfirm, setShowInitConfirm] = useState(false);
    const [showInitComplete, setShowInitComplete] = useState(false);

    /**
     * モンスタースキル初期化
     */
    const initMonsterSkills = useCallback( async () => {
        await api.PUT(URLS.INIT_ALL_MONSTERS_SKILLS);
    }, []);

    return (
        <>
            <CommonButton text="全スキル初期化"
                    onClick={() => setShowInitConfirm(true)}
                    width={150}
                    display={selectEditType === 2 ? "inline" : "none"}
                    styleObj={{marginRight: "20px"}}
                    />
            {/* 初期化確認ダイアログ */}
            <CommonDialogFrame showDialog={showInitConfirm}>
                <h3 style={{margin: 0}}>全モンスターのスキルを初期状態に戻します。</h3>
                <h3 style={{margin: 0}}>よろしいですか？</h3>
                <div style={{textAlign: "end"}}>
                    <CommonButton text="いいえ" onClick={() => setShowInitConfirm(false)}/>
                    <CommonButton text="はい" onClick={() => {
                        initMonsterSkills();
                        setShowInitComplete(true);
                        globalThis.location.href = "/EditPage";
                    }
                        }/>
                </div>
            </CommonDialogFrame>
            {/* 初期化完了通知 */}
        </>

    );
}

export default ClearAllSkillsBlock;