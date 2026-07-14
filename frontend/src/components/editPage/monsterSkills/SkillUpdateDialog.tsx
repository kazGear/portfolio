import { COLORS } from "../../../lib/Constants";
import CommonButton from "../../common/CommonButton";
import CommonDialogFrame from "../../common/CommonDialogFrame";

interface ArgProps {
    showDialog: boolean;
    setShowUpdateDialog: React.Dispatch<React.SetStateAction<boolean>>;
}

const SkillUpdateDialog = ({
    showDialog, setShowUpdateDialog}: ArgProps
) => {
    return (
        <CommonDialogFrame showDialog={showDialog}>
            <h2 style={{color: COLORS.MAIN_FONT}}>
                スキル更新が完了しました。
            </h2>
            <div style={{textAlign: "end"}}>
                <CommonButton text="閉じる"
                        onClick={() => setShowUpdateDialog(false)}
                        />
            </div>
        </CommonDialogFrame>
    );
}

export default SkillUpdateDialog;