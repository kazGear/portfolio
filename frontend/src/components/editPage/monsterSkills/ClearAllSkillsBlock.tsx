import { useCallback, useState } from "react";
import { useServer, useServerWithQuery } from "../../../hooks/useHooksOfCommon";
import Button from "../../common/Button";
import { KEYS, URLS } from "../../../lib/Constants";
import DialogFrame from "../../common/DialogFrame";
import { EditSkillsDTO } from "../../../types/Edit";

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
    const goToServer = useServer();
    const initMonsterSkills = useCallback( async () => {
        await goToServer(URLS.INIT_ALL_MONSTERS_SKILLS);
    }, []);

    return (
        <>
            <Button text="全スキル初期化"
                    onClick={() => setShowInitConfirm(true)}
                    width={150}
                    display={selectEditType === 2 ? "inline" : "none"}
                    styleObj={{marginRight: "20px"}}
                    />
            {/* 初期化確認ダイアログ */}
            <DialogFrame showDialog={showInitConfirm}>
                <h3 style={{margin: 0}}>全モンスターのスキルを初期状態に戻します。</h3>
                <h3 style={{margin: 0}}>よろしいですか？</h3>
                <div style={{textAlign: "end"}}>
                    <Button text="いいえ" onClick={() => setShowInitConfirm(false)}/>
                    <Button text="はい" onClick={() => {
                        initMonsterSkills();
                        setShowInitComplete(true);
                        window.location.href = "/EditPage";
                    }
                        }/>
                </div>
            </DialogFrame>
            {/* 初期化完了通知 */}
        </>

    );
}

export default ClearAllSkillsBlock;