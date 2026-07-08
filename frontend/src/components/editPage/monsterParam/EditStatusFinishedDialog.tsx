import { useRefreshMonsterStatus } from "../../../hooks/useHooksOfEdit";
import { COLORS } from "../../../lib/Constants";
import { EditMonsterDTO } from "../../../types/Edit";
import CommonButton from "../../common/CommonButton";
import CommonDialogFrame from "../../common/CommonDialogFrame";

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
    const refreshMonsterStatus = useRefreshMonsterStatus();

    return (
        <CommonDialogFrame showDialog={isShow}>
            <h2 style={{color: COLORS.MAIN_FONT_COLOR}}>
                ステータス更新が完了しました。
            </h2>
            <div style={{textAlign: "end"}}>
                <CommonButton text="閉じる"
                        onClick={() => {
                            refreshMonsterStatus(setEditMonsters);
                            setShowDialog(false);
                        }}
                        />
            </div>
        </CommonDialogFrame>
    );
}

export default EditStatusFinishedDialog;