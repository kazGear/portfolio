import { MonsterDTO } from "../../types/MonsterBattle";

interface ArgProps {
    monsters: MonsterDTO[];
    rowClickHandler: (row: any) => void;
    validSelect: (row: any) => void;
    selectError: boolean;
}

const MonsterSelectorBlock = ({
    monsters,
    rowClickHandler,
    validSelect,
    selectError
}: ArgProps) => {
    return (
        <>
            <table>
                <tbody>
                {
                    monsters.map((monster, index) => (
                    <tr
                        key={index}
                        onClick={() => rowClickHandler(monster)}
                        >
                        {/* モンスター名、ラジオボタン */}
                        <td>
                            <label>
                                <input
                                    type="radio"
                                    name="betMonster"
                                    value={monster.MonsterId}
                                    onChange={(e) => validSelect(e)}
                                />
                                &emsp;{monster.MonsterName}
                            </label>
                        </td>
                        {/* 賭けレート */}
                        <td>
                            <span style={{marginLeft: "20px"}}>（レート：{monster.BetRate}倍）</span>
                        </td>
                    </tr>
                    ))
                }
                </tbody>
            </table>

            {
                selectError ? <span style={{color: "red"}}>※モンスターを選択してください。</span> : ""
            }
        </>
    );
}

export default MonsterSelectorBlock;