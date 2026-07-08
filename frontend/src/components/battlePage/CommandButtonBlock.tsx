import styled from "styled-components";
import CommonButton from "../common/CommonButton";
import { useEffect, useState } from "react";

const SdivCommonButtonFrame = styled.div`
    height: 40px;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
`;

const resetHandler = () => {
    globalThis.location.reload();
}

interface ArgProps {
    battleStartHandler: () => Promise<void>;
    nextTurnHandler: () => void;
    monsterCount: number;
    battleStarted: boolean;
}

const CommandCommonButtonBlock = ({
    battleStartHandler,
    nextTurnHandler,
    monsterCount,
    battleStarted
}: ArgProps) => {

    const [battleBtnDisabled, setBattleBtnDisabled] = useState(battleStarted);
    const [nextBtnDisabled, setNextBtnDisabled] = useState(true);

    // ボタン活性制御, リセット
    useEffect(() => {
        if (battleStarted) {
            setBattleBtnDisabled(true);
            setNextBtnDisabled(false);
        }
        if (monsterCount <= 0) {
            setBattleBtnDisabled(false);
            setNextBtnDisabled(true);
        }
    }, [monsterCount, battleStarted]);

    return (
        <SdivCommonButtonFrame>
            <CommonButton
                id={"nextTurn"}
                text={"戦闘開始！"}
                styleObj={{width: "30%", height: "30px"}}
                onClick={battleStartHandler}
                disabled={battleBtnDisabled}
                />
            <CommonButton
                id={"nextMessage>"}
                text={"次のターン！"}
                styleObj={{width: "30%", height: "30px"}}
                onClick={nextTurnHandler}
                disabled={nextBtnDisabled}
                />
            <CommonButton
                id={"nextMessage>"}
                text={"リセット"}
                styleObj={{width: "30%", height: "30px"}}
                onClick={resetHandler}
                />
        </SdivCommonButtonFrame>
    );
}

export default CommandCommonButtonBlock;