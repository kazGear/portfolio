import styled from "styled-components";
import CommonStrong from "../common/CommonStrong";
import monsterImages from "../../lib/MonsterImages";
import { MonsterDTO } from "../../types/MonsterBattle";
import { useEffect, useState } from "react";
import { URLS } from "../../lib/Constants";
import CommonNowLoading from "../common/CommonNowLoading";
import { api } from "../../lib/apiClient";

const MonstersFrame = styled.div`
    height: 100%;
    margin: 20px;
`;
const Img = styled.img`
    width: 50px;
    height: 50px;
    vertical-align: middle;
    margin: 5px 0 5px 0;
}
`;
const Table = styled.table`
    height: 100%;
`;
const Td1 = styled.td`
    min-width: 150px;
`;
const Td2 = styled.td`
    min-width: 80px;
`;
const Td3 = styled.td`
    width: 100px;
    min-width: 80px;
`;
const Td4 = styled.td`
    width: 110px;
    min-width: 100px;
`;
const Td5 = styled.td`
    width: 120px;
    min-width: 110px;
`;
const Td6 = styled.td`
    min-width: 110px;
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
    useEffect(() => {
        const selectMonsterCount = async () => {
            const monsterCount = await api.POST<string>(URLS.GET_MONSTER_COUNT, loginId);
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
                <CommonNowLoading alt="ローディング"/>
            </div>
        );
    }

    return (
        <MonstersFrame>
            <CommonStrong>開放モンスター&emsp;{getMonsterCount}</CommonStrong>

            <Table>
                <tbody>
                    { monsters !== null ?
                        monsters.map((monster, index) => (
                            <tr key={index}>
                                <Td1>{monster.MonsterName}</Td1>
                                <Td2>
                                    <Img src={monsterImages(monster.MonsterId)} alt={monster.MonsterName}/>
                                </Td2>
                                <Td3>HP：{monster.Hp}</Td3>
                                <Td4>攻撃力：{monster.Attack}</Td4>
                                <Td5>スピード：{monster.Speed}</Td5>
                                <Td6>回避力：{monster.Dodge}</Td6>
                            </tr>
                        ))
                        : ""
                    }
                </tbody>
            </Table>
        </MonstersFrame>
    );
}

export default MonstersBlock;