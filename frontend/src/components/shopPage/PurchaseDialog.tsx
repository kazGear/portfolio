import styled from "styled-components";
import Accent from "../common/Accent";
import Button from "../common/Button";
import DialogFrame from "../common/DialogFrame";
import { DECO } from "../../lib/Constants";
import Strong from "../common/Strong";

const SbuttonFrame = styled.div`
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
        <DialogFrame showDialog={showDialog}>
            <Strong>{DECO.BLOCK_LINE}</Strong><br/>
            <Strong>{DECO.BLOCK_LINE_R}</Strong>
                <h2 style={{margin: "5px 0 5px 0"}}>購入完了</h2>
            <Strong>{DECO.BLOCK_LINE}</Strong><br/>
            <Strong>{DECO.BLOCK_LINE_R}</Strong>

            <h2 style={{marginBottom: 0}}><Accent>{purchaseItem}</Accent>を獲得しました！</h2>

            <SbuttonFrame>
                <Button text="閉じる"
                        onClick={() => setShowPurchaseDialog(false)}
                        />
            </SbuttonFrame>
        </DialogFrame>
    );
}

export default PurchaseDialog;