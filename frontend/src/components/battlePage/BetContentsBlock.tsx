import { MonsterDTO } from "../../types/MonsterBattle";
import CommonDialogFrame from "../common/CommonDialogFrame";
import GameBetContentsBlock from "./GameBetBlock";


interface ArgProps {
    monsters: MonsterDTO[];
    setBetMonster: React.Dispatch<React.SetStateAction<MonsterDTO | null>>;
    setBetGil: React.Dispatch<React.SetStateAction<number>>;
    showBetDialog: boolean;
    setShowBetDialog: React.Dispatch<React.SetStateAction<boolean>>;
}

const GameBetBlock = ({
    monsters,
    setBetMonster,
    setBetGil,
    showBetDialog,
    setShowBetDialog}: ArgProps
) => {
    return (
        <CommonDialogFrame showDialog={showBetDialog}
                     showFilter={true}>
            <GameBetContentsBlock
                    monsters={monsters}
                    setBetMonster={setBetMonster}
                    setBetGil={setBetGil}
                    setShowDialog={setShowBetDialog} />
        </CommonDialogFrame>
    );
}

export default GameBetBlock;