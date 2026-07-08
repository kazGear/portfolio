import CommonDialogFrame from "../common/CommonDialogFrame";
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
        <CommonDialogFrame showDialog={showResultDialog}>
            <BattleStartContentsBlock battleStartHandler={battleStartHandler}
                                      selectMonstersCountHandler={selectMonstersCountHandler}/>
        </CommonDialogFrame>
    );
}

export default GameStartBlock;