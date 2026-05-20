import { ChangeEvent, useLayoutEffect, useState } from "react";
import { useServerWithQuery } from "../../hooks/useHooksOfCommon";
import Select from "../common/Select";
import { URLS } from "../../lib/Constants";
import { MonsterTypeDTO } from "../../types/BattleReport";

interface ArgProps {
    setMonsterTypeId: React.Dispatch<React.SetStateAction<string>>;
}

const MonsterTypesListBlock = ({setMonsterTypeId}: ArgProps) => {
    const [monsterTypes, setMonsterTypes] = useState<MonsterTypeDTO[]>([]);
    const goToServer = useServerWithQuery();

    useLayoutEffect(() => {
        const fetchTypes = async () => {
            const types: MonsterTypeDTO[] = await goToServer(URLS.INIT_BATTLE_REPORT);
            setMonsterTypes(types);
        }
        fetchTypes();
    }, []);

    const changeMonsterTypeHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        setMonsterTypeId(e.target.value);
    }

    return (
        <>
            <Select title="モンスター種" onChange={changeMonsterTypeHandler}>
                <option value="0">指定なし</option>
                {
                    monsterTypes.map((monster, index) => {
                        return (
                            <option value={monster.MonsterTypeId} key={index}>
                                {monster.MonsterTypeName}
                            </option>
                    )})
                }
            </Select>
        </>
    );
}

export default MonsterTypesListBlock;