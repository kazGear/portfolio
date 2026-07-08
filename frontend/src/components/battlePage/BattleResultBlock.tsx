
import { MetaDataDTO, MonsterDTO } from "../../types/MonsterBattle";
import CommonDialogFrame from "../common/CommonDialogFrame";
import BattleResultContentsBlock from "./BattleResultContentsBlock";
import { ShopDTO } from "../../types/Shop";

interface ArgProps {
    log:  MetaDataDTO | null;
    betMonster: MonsterDTO | null;
    betGil: number;
    showResultDialog: boolean;
    newShops: ShopDTO[];
}

const BattleResultBlock = ({
    log, betMonster, betGil, showResultDialog, newShops }: ArgProps) => {

    return (
        <CommonDialogFrame showDialog={showResultDialog}
                           showFilter={true}>
            <BattleResultContentsBlock
                    log={log}
                    betMonster={betMonster}
                    betGil={betGil}
                    newShops={newShops} />
        </CommonDialogFrame>
    );
}

export default BattleResultBlock;