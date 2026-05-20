import DialogFrame from "../common/DialogFrame";
import BattleStartContentsBlock from "./BattleStartContentsBlock";


interface ArgProps {
    battleStartHandler: (e: any) => Promise<void>;
    selectMonstersCountHandler: (e: any) => void;
    showResultDialog : boolean
}

const GameStartBlock = ({
    battleStartHandler, selectMonstersCountHandler, showResultDialog}: ArgProps
) => {
    return (
        <DialogFrame showDialog={showResultDialog}>
            <BattleStartContentsBlock battleStartHandler={battleStartHandler}
                                      selectMonstersCountHandler={selectMonstersCountHandler}/>
        </DialogFrame>
    );
}

export default GameStartBlock;