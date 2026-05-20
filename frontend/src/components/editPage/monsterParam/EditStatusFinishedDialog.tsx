import { useServerWithQuery } from "../../../hooks/useHooksOfCommon";
import { useRefreshMonsterStatus } from "../../../hooks/useHooksOfEdit";
import { COLORS } from "../../../lib/Constants";
import { EditMonsterDTO } from "../../../types/Edit";
import Button from "../../common/Button";
import DialogFrame from "../../common/DialogFrame";

interface ArgProps {
    isShow: boolean;
    setShowDialog: React.Dispatch<React.SetStateAction<boolean>>;
    setEditMonsters: React.Dispatch<React.SetStateAction<EditMonsterDTO[]>>;
}

const EditStatusFinishedDialog = ({
    isShow, setShowDialog, setEditMonsters}: ArgProps
) => {
    /**
     * 更新後のステータスを反映
     */
    const goToServer = useServerWithQuery()
    const refreshMonsterStatus = useRefreshMonsterStatus();

    return (
        <DialogFrame showDialog={isShow}>
            <h2 style={{color: COLORS.MAIN_FONT_COLOR}}>
                ステータス更新が完了しました。
            </h2>
            <div style={{textAlign: "end"}}>
                <Button text="閉じる"
                        onClick={() => {
                            refreshMonsterStatus(goToServer, setEditMonsters);
                            setShowDialog(false);
                        }}
                        />
            </div>
        </DialogFrame>
    );
}

export default EditStatusFinishedDialog;