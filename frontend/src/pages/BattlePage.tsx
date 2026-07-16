import styled from "styled-components";
import { BattleResults, MetaDataDTO, MonsterDTO } from "../types/MonsterBattle";
import { useCallback, useEffect, useRef, useState } from "react";
import { KEYS, URLS } from "../lib/Constants";
import { useRegistResult } from "../hooks/useHooksOfBattle";
import { isEmpty } from "../lib/CommonLogic";
import BattleResultBlock from "../components/battlePage/BattleResultBlock";
import GameBetBlock from "../components/battlePage/BetContentsBlock";
import GameStartBlock from "../components/battlePage/GameStartBlock";
import MessageWindowBlock from "../components/battlePage/MessageWindowBlock";
import CommandButtonBlock from "../components/battlePage/CommandButtonBlock";
import MonsterWindowBlock from "../components/battlePage/MonsterWindowBlock";
import { ShopDTO } from "../types/Shop";
import { api } from "../lib/apiClient";
import { useCheckToken } from "../hooks/useHooksOfCommon";
import useApiErrorHandler from "../hooks/useApiErrorHandler";
import { ApiError } from "../types/ApiError";

const OutSideFrame = styled.div`
    position: relative;
    margin: 20px 20px 20px 20px;
    height: 100%;
`;
const MonsterWindowFrame = styled.div`
    display: flex;
    justify-content: center;
`;

/**
 * モンスター１体分のログを作成
 */
const createShortLog = (battleLog: MetaDataDTO[]): [MetaDataDTO[], number] => {
    const shortLog = [];
    let index = 0;

    for (const log of battleLog) {
        shortLog.push(log);
        index++;
        // モンスター１体分としてログを区切る
        if (log.IsStop || shortLog.length <= 0) {
            return [shortLog, index];
        }
    }
    return [shortLog, 0];
}

const BattlePage = () => {
    // モンスター関係
    const [monsters, setMonsters] = useState<MonsterDTO[]>([]); // 戦闘モンスター
    const [monsterCount, setMonsterCount] = useState(0);
    const selectMonstersCount = useRef(2);

    // 賭け関係
    const [betMonster, setBetMonster] = useState<MonsterDTO | null>(null);
    const [betGil, setBetGil] = useState(0);
    const [battleStarted, setBattleStarted] = useState(false);

    // サーバから送られるログ
    const [battleLog, setBattleLog] = useState<MetaDataDTO[]>([]); // １ターン・全モンスター文のログ
    const [shortLog, setShortLog] = useState<MetaDataDTO[]>([]); // １ターン・各モンスターのログ
    const [resultLog, setResultLog] = useState<MetaDataDTO | null>(null); // 勝敗結果

    // ダイアログの表示可否
    const [showResultDialog, setShowResultDialog] = useState(false);
    const [showStartDialog, setShowStartDialog]   = useState(true);
    const [showBetDialog, setShowBetDialog]       = useState(false);
    const [showBattleView, setShowBattleView]     = useState(false);

    const [loginId, setLoginId]   = useState<string>("");
    const [newShops, setNewShops] = useState<ShopDTO[]>([]);

    const errorHandler = useApiErrorHandler()

    useEffect(() => {
        setLoginId(localStorage.getItem(KEYS.USER_ID)!);
    }, []);

    useCheckToken();

    /**
     * 戦闘モンスター数選択
     */
    const selectMonstersCountHandler = useCallback((e: any) => {
        selectMonstersCount.current = e.target.value;
    }, []);

    /**
     * ゲーム開始、モンスター初期化
    */
   const gameStartHandler = useCallback(async (e: any) => {
       try {
            const initMonsters = await api.POST<MonsterDTO[]>(URLS.INIT_MONSTERS, {
                selectMonstersCount: selectMonstersCount.current.toString(),
                loginId:             loginId,
            });

            if (isEmpty(initMonsters)) throw new ApiError(500, "Init monsters failed ...");

            setMonsters([...initMonsters!]);
            setMonsterCount(initMonsters!.length);
            // 画面切り替え
            setShowStartDialog(false);
            setShowBetDialog(true);
            setShowBattleView(true);

        } catch (e) {
            console.log(e);
            errorHandler(e);
        }
    }, [loginId])
    /**
     *  全モンスターが行動。戦闘ログを取得
     */
    const battleHandler = useCallback( async () => {
        try {
            const moveResult = await api.POST<BattleResults>(URLS.BATTLE_NEXT_TURN, monsters);

            if (isEmpty(moveResult)) throw new ApiError(500, "Monsters missing move ...");

            setMonsters([...moveResult!.Monsters]);
            setBattleLog([...moveResult!.BattleLog]);
            setBattleStarted(true);
            setMonsterCount([...moveResult!.Monsters].length);
        } catch (e) {
            console.log(e);
            errorHandler(e);
        }
    }, [monsters]);

    const insertBattleResult = useRegistResult();
    /**
     *  各モンスターのターン送り。戦闘ログを元に描画、表現を行う
     */
    const nextTurnHandler = useCallback( async () => {
        const [shortLog, index] = createShortLog([...battleLog]); // モンスター１体分のログ
        setShortLog([...shortLog]);

        const afterLog = battleLog.slice(index); // 残りのログ
        setBattleLog([...afterLog]);
        setMonsterCount(monsterCount - 1);
        setBattleStarted(false);

        try {
            // モンスターの戦績を記録
            const lastLog: MetaDataDTO | undefined = shortLog.pop();
            insertBattleResult({
                monsters, lastLog, setResultLog, setShowResultDialog
            });

            // ユーザの成績を記録
            if (!isEmpty(lastLog) && !isEmpty(lastLog!.WinnerMonsterId)) {
                const newShops = await api.PUT<ShopDTO[]>(URLS.RECORD_USER_RESULT, {
                    betMonsterId:     betMonster!.MonsterId,
                    betGil:           betGil,
                    betRate:          betMonster!.BetRate,
                    winningMonsterId: lastLog?.WinnerMonsterId!,
                    loginId:          loginId,
                });
                setNewShops(newShops!);
            }
        } catch (e) {
            console.log(e);
            errorHandler(e);
        }
    }, [monsters, battleLog, monsterCount]);

    return (
        <>
            <OutSideFrame style={{display: showBattleView ? "block" : "none", overflow: "hidden"}} >
                {/* モンスター画面 */}
                <MonsterWindowFrame>
                    {
                        monsters.map((monster, index) => (
                            <MonsterWindowBlock
                                monster={monster}
                                shortLog={shortLog}
                                key={(monster.MonsterId + `_${index}`)}
                                />
                        ))
                    }
                </MonsterWindowFrame>
                {/* 操作部 */}
                <CommandButtonBlock
                    battleStartHandler={battleHandler}
                    nextTurnHandler={nextTurnHandler}
                    monsterCount={monsterCount}
                    battleStarted={battleStarted}
                    />
                {/* メッセージ部 */}
                <MessageWindowBlock
                    shortLog={shortLog}
                    />
            </OutSideFrame>
            {/* 開始ダイアログ */}
            <GameStartBlock
                battleStartHandler={gameStartHandler}
                selectMonstersCountHandler={selectMonstersCountHandler}
                showResultDialog={showStartDialog}
                />
            {/* 賭けダイアログ */}
            <GameBetBlock
                monsters={monsters}
                setBetMonster={setBetMonster}
                setBetGil={setBetGil}
                showBetDialog={showBetDialog}
                setShowBetDialog={setShowBetDialog} />
            {/* 終了ダイアログ */}
            <BattleResultBlock
                log={resultLog}
                betMonster={betMonster}
                betGil={betGil}
                showResultDialog={showResultDialog}
                newShops={newShops}/>
        </>
    );
}

export default BattlePage;