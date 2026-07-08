import styled from "styled-components";
import CommonAccent from "../common/CommonAccent";
import CommonButton from "../common/CommonButton";
import CommonDialogFrame from "../common/CommonDialogFrame";
import { DECO } from "../../lib/Constants";
import CommonStrong from "../common/CommonStrong";

const ButtonFrame = styled.div`
    height: 30%;
    text-align: end;
    align-content: end;
`;

interface ArgProps {
    showDialog: boolean;
    purchaseItem: string;
    setShowPurchaseDialog: React.Dispatch<React.SetStateAction<boolean>>;
}

const PurchaseDialog = ({showDialog, purchaseItem, setShowPurchaseDialog}: ArgProps

) => {
    return (
        <CommonDialogFrame showDialog={showDialog}>
            <CommonStrong>{DECO.BLOCK_LINE}</CommonStrong><br/>
            <CommonStrong>{DECO.BLOCK_LINE_R}</CommonStrong>
                <h2 style={{margin: "5px 0 5px 0"}}>購入完了</h2>
            <CommonStrong>{DECO.BLOCK_LINE}</CommonStrong><br/>
            <CommonStrong>{DECO.BLOCK_LINE_R}</CommonStrong>

            <h2 style={{marginBottom: 0}}>
                <CommonAccent>{purchaseItem}</CommonAccent>を獲得しました！
            </h2>

            <ButtonFrame>
                <CommonButton text="閉じる"
                        onClick={() => setShowPurchaseDialog(false)}
                        />
            </ButtonFrame>
        </CommonDialogFrame>
    );
}

export default PurchaseDialog;