import styled from "styled-components";
import Strong from "../common/Strong";
import monsterImages from "../../lib/MonsterImages";
import { MonsterDTO } from "../../types/MonsterBattle";
import { useLayoutEffect, useState } from "react";
import { useServerWithQuery } from "../../hooks/useHooksOfCommon";
import { URLS } from "../../lib/Constants";
import BorderTd from "../common/BorderTd";
import NowLoading from "../common/NowLoading";

const SdivMonsters = styled.div`
    height: 100%;
    margin: 20px;
`;
const SImg = styled.img`
    width: 50px;
    height: 50px;
    vertical-align: middle;
    margin: 5px 0 5px 0;
}
`;
const Stable = styled.table`
    height: 100%;
`;
const Std1 = styled.td`
    width: 180px;
    min-width: 90px;
`;
const Std2 = styled.td`
    width: 80px;
`;
const Std3 = styled.td`
    width: 80px;
    min-width: 60px;
`;
const Std4 = styled.td`
    width: 90px;
    min-width: 70px;
`;
const Std5 = styled.td`
    width: 90px;
    min-width: 80px;
`;

interface ArgProps {
    monsters: MonsterDTO[] | null;
    loginId: string | null;
}

const MonstersBlock = ({monsters, loginId}: ArgProps) => {
    const [getMonsterCount, setGetMonsterCount] = useState<string | null>(null);
    const [isNowLoading, setIsNowLoading] = useState(true);

    /**
     * 使用権開放済モンスターの取得
     */
    const goToServer = useServerWithQuery();
    useLayoutEffect(() => {
        const selectMonsterCount = async () => {
            const monsterCount: string = await goToServer(URLS.GET_MONSTER_COUNT + `?loginId=${loginId}`);
            setGetMonsterCount(monsterCount);
        }
        selectMonsterCount();
        setIsNowLoading(false);
    }, []);

    /**
     * ローディング
     */
    if (isNowLoading) {
        return (
            <div style={{margin: "100px"}}>
                <NowLoading alt="ローディング"/>
            </div>
        );
    }

    return (
        <SdivMonsters>
            <Strong>開放モンスター&emsp;{getMonsterCount}</Strong>

            <Stable>
                <tbody>
                    { monsters !== null ?
                        monsters.map((monster, index) => (
                            <tr key={index}>
                                <Std1>{monster.MonsterName}</Std1>
                                <Std2><SImg src={monsterImages(monster.MonsterId)} alt="アイコン"/>{}</Std2>
                                <Std3>HP：{monster.Hp}</Std3>
                                <Std4>攻撃力：{monster.Attack}</Std4>
                                <Std5>スピード：{monster.Speed}</Std5>
                            </tr>
                        ))
                        : ""
                    }
                </tbody>
            </Stable>
        </SdivMonsters>
    );
}

export default MonstersBlock;