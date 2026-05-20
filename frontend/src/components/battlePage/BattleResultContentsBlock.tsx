import styled from "styled-components";
import { COLORS, DECO } from "../../lib/Constants";
import Button from "../common/Button";
import { MetaDataDTO, MonsterDTO } from "../../types/MonsterBattle";
import { ShopDTO } from "../../types/Shop";
import Strong from "../common/Strong";

const SpLine = styled.p`
    color: ${COLORS.ACCENT_FONT_PINK};
    line-height: 0.5;
    margin: 5px 0 0 0;
`;
const SdivBetResultFrame = styled.div`
    height: 20%;
    text-align: center;
    align-content: center;
`;
const SdivMessageFrame = styled.div`
    height: 20%;
    text-align: center;
    align-content: center;
`;
const SdivButtonFrame = styled.div`
    margin-top: 15px;
    height: 10%;
    text-align: end;
    align-content: flex-end;
`;
const Sh1 = styled.h1`
    color: ${COLORS.CAPTION_FONT_COLOR};
    margin: 5px 0 5px 0;
`;
const Sspan = styled.span`
    color: ${COLORS.ACCENT_FONT_PINK};
`;

const gamesetHandler = () => {
    window.location.reload();
}

interface ArgProps {
    log: MetaDataDTO | null;
    betMonster: MonsterDTO | null;
    betGil: number;
    newShops: ShopDTO[];
}

const BattleResultContentsBlock = (
     {log, betMonster, betGil, newShops}: ArgProps
    ) => {
    const strike = log && betMonster &&
        log.WinnerMonsterId === betMonster.MonsterId;

    return (
        <>
        {log &&
            <>
            <SpLine>{ log.ExistWinner ? DECO.BLOCK_LINE : "" }</SpLine>
            <SpLine>{ log.ExistWinner ? DECO.BLOCK_LINE_R : "" }</SpLine>
            <Sh1>
                {
                    log.ExistWinner
                    ? `Winner: ${log.WinnerMonsterName} !!`
                    : ""
                }
            </Sh1>
            <SpLine>{ log.ExistWinner ? DECO.BLOCK_LINE : "" }</SpLine>
            <SpLine>{ log.ExistWinner ? DECO.BLOCK_LINE_R : "" }</SpLine>

            <h1>{log.AllLoser ? "draw ..." : ""}</h1>

            <SdivBetResultFrame>
                <h2>(∩´∀｀)∩<Sspan>獲得賞金&emsp;</Sspan>
                {
                    strike ? <Sspan>{Math.trunc(betGil * betMonster.BetRate)}</Sspan> : <Sspan>0</Sspan>
                } Gil
                </h2>
            </SdivBetResultFrame>

            <SdivMessageFrame>
                {
                    newShops.length > 0 ? <h3 style={{margin: 0}}>新しいショップが解禁されました。</h3> : ""
                }
                {
                    newShops.map((shop, index) => (
                        <Strong key={index}>{shop.ShopName}&emsp;</Strong>
                    ))
                }
            </SdivMessageFrame>

            <SdivButtonFrame>
                <Button text={"終了"} width={120} onClick={gamesetHandler}/>
            </SdivButtonFrame>
            </>
        }
        </>
    );
};
export default BattleResultContentsBlock;