import styled from "styled-components";
import { BattleResults, MetaDataDTO, MonsterDTO } from "../types/MonsterBattle";
import { useCallback, useEffect, useRef, useState } from "react";
import { useCheckToken, useServerWithQuery, useServerWithJson } from "../hooks/useHooksOfCommon";
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

const SdivOutSideFrame = styled.div`
    position: relative;
    margin: 20px;
    height: 100%;
`;
const SdivMonsterWindowFrame = styled.div`
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
    const [showStartDialog, setShowStartDialog] = useState(true);
    const [showBetDialog, setShowBetDialog] = useState(false);
    const [showBattleView, setShowBattleView] = useState(false);

    const [loginId, setLoginId] = useState<string | null>("");
    const [newShops, setNewShops] = useState<ShopDTO[]>([]);

    useCheckToken();

    useEffect(() => {
        setLoginId(localStorage.getItem(KEYS.USER_ID));
    }, []);

    /**
     * 戦闘モンスター数選択
     */
    const selectMonstersCountHandler = useCallback((e: any) => {
        selectMonstersCount.current = e.target.value;
    }, []);

    /**
     * ゲーム開始、モンスター初期化
    */
   const goToServer = useServerWithQuery();
   const gameStartHandler = useCallback(async (e: any) => {
        const initMonsters: MonsterDTO[] = await goToServer(
            `${URLS.INIT_MONSTERS}?selectMonstersCount=${selectMonstersCount.current}&loginId=${loginId}`);
        setMonsters([...initMonsters]);
        setMonsterCount(initMonsters.length);
        // 画面切り替え
        setShowStartDialog(false);
        setShowBetDialog(true);
        setShowBattleView(true);
    }, [loginId])
    /**
     *  全モンスターが行動。戦闘ログを取得
     */
    const goToServerWithJson = useServerWithJson();
    const battleHandler = async () => {
        const moveResult: BattleResults = await goToServerWithJson(
            monsters, URLS.BATTLE_NEXT_TURN
       );
       console.log(moveResult.Monsters); // tmp //////////////////////////
       setMonsters([...moveResult.Monsters]); console.log(moveResult.Monsters);
       setBattleLog([...moveResult.BattleLog]);
       setBattleStarted(true);
       setMonsterCount([...moveResult.Monsters].length);
   };

    const insertResult = useServerWithJson();
    const insertBattleResult = useRegistResult();
    /**
     *  各モンスターのターン送り。戦闘ログを元に描画、表現を行う
     */
    const nextTurnHandler = async () => {
        const [shortLog, index] = createShortLog([...battleLog]); // モンスター１体分のログ
        setShortLog([...shortLog]);

        const afterLog = battleLog.slice(index); // 残りのログ
        setBattleLog([...afterLog]);
        setMonsterCount(monsterCount - 1);
        setBattleStarted(false);

        // モンスターの戦績を記録
        const lastLog: MetaDataDTO | undefined = shortLog.pop();
        insertBattleResult({
            monsters, lastLog, setResultLog, setShowResultDialog, insertResult
        });
        // ユーザの成績を記録
        if (!isEmpty(lastLog) && !isEmpty(lastLog!.WinnerMonsterId)) {
            const newShops: ShopDTO[] = await goToServer(
                `${URLS.RECORD_USER_RESULT}?betMonsterId=${betMonster?.MonsterId}
                                            &betGil=${betGil}
                                            &betRate=${betMonster!.BetRate}
                                            &winningMonsterId=${lastLog?.WinnerMonsterId}
                                            &loginId=${loginId}`
            );
            setNewShops(newShops);
        }
    }

    return (
        <>
            <SdivOutSideFrame style={{display: showBattleView ? "block" : "none", overflow: "hidden"}} >
                {/* モンスター画面 */}
                <SdivMonsterWindowFrame>
                    {
                        monsters.map((monster, index) => (
                            <MonsterWindowBlock
                                monster={monster}
                                shortLog={shortLog}
                                key={(monster.MonsterId + `_${index}`)}
                                />
                        ))
                    }
                </SdivMonsterWindowFrame>
                {/* 操作部 */}
                <CommandButtonBlock
                    battleStartHandler={battleHandler}
                    nextTurnHandler={nextTurnHandler}
                    monsterCount={monsterCount}
                    battleStarted={battleStarted}
                    />
                <MessageWindowBlock
                    shortLog={shortLog}
                    />
            </SdivOutSideFrame>
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