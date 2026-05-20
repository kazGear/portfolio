import { useCallback, useState } from "react";
import Button from "../../common/Button";
import { useServerWithQuery } from "../../../hooks/useHooksOfCommon";
import { URLS } from "../../../lib/Constants";
import DialogFrame from "../../common/DialogFrame";
import { useRefreshMonsterStatus } from "../../../hooks/useHooksOfEdit";
import { EditMonsterDTO } from "../../../types/Edit";

interface ArgProps {
    setEditMonsters: React.Dispatch<React.SetStateAction<EditMonsterDTO[]>>;
    selectEditType: number;
}

const ClearAllStatusBlock = ({setEditMonsters, selectEditType}: ArgProps) => {
    const [showInitConfirm, setShowInitConfirm] = useState(false);
    const [showInitComplete, setShowInitComplete] = useState(false);
    /**
     * 全モンスターステータス初期化
     */
    const goToServer = useServerWithQuery();
    const clearAllMonstersStatus = useCallback( async () => {
        await goToServer(URLS.INIT_ALL_MONSTERS_STATUS);
    }, [])
    /**
     * 更新後のステータスを反映
     */
    const refreshMonsterStatus = useRefreshMonsterStatus();

    return (
        <>
            <Button text="全ステータス初期化"
                    onClick={() => setShowInitConfirm(true)}
                    width={150}
                    display={selectEditType === 1 ? "inline" : "none"}
                    styleObj={{marginRight: "20px"}}
                    />
            {/* 初期化確認ダイアログ */}
            <DialogFrame showDialog={showInitConfirm}>
                <h3 style={{margin: 0}}>全モンスターのステータスを初期状態に戻します。</h3>
                <h3 style={{margin: 0}}>よろしいですか？</h3>
                <div style={{textAlign: "end"}}>
                    <Button text="いいえ" onClick={() => setShowInitConfirm(false)}/>
                    <Button text="はい" onClick={() => {
                        clearAllMonstersStatus();
                        refreshMonsterStatus(goToServer, setEditMonsters);
                        setShowInitComplete(true);
                        window.location.href = "/EditPage"
                    }
                        }/>
                </div>
            </DialogFrame>
        </>

    );
}

export default ClearAllStatusBlock;