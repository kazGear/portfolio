import styled from "styled-components";
import Button from "../common/Button";
import { useEffect, useState } from "react";

const SdivButtonFrame = styled.div`
    height: 40px;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
`;

const resetHandler = () => {
    window.location.reload();
}

interface ArgProps {
    battleStartHandler: () => Promise<void>;
    nextTurnHandler: () => void;
    monsterCount: number;
    battleStarted: boolean;
}

const CommandButtonBlock = ({
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
        <SdivButtonFrame>
            <Button
                id={"nextTurn"}
                text={"戦闘開始！"}
                styleObj={{width: "30%"}}
                onClick={battleStartHandler}
                disabled={battleBtnDisabled}
                />
            <Button
                id={"nextMessage>"}
                text={"次のターン！"}
                styleObj={{width: "30%"}}
                onClick={nextTurnHandler}
                disabled={nextBtnDisabled}
                />
            <Button
                id={"nextMessage>"}
                text={"リセット"}
                styleObj={{width: "30%"}}
                onClick={resetHandler}
                />
        </SdivButtonFrame>
    );
}

export default CommandButtonBlock;