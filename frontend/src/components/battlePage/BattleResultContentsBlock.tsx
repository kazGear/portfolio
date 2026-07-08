import styled from "styled-components";
import { COLORS, DECO } from "../../lib/Constants";
import CommonButton from "../common/CommonButton";
import { MetaDataDTO, MonsterDTO } from "../../types/MonsterBattle";
import { ShopDTO } from "../../types/Shop";
import CommonStrong from "../common/CommonStrong";

const Line = styled.p`
    color: ${COLORS.ACCENT_FONT_PINK};
    line-height: 0.5;
    margin: 5px 0 0 0;
`;
const BetResultFrame = styled.div`
    height: 20%;
    text-align: center;
    align-content: center;
`;
const MessageFrame = styled.div`
    height: 20%;
    text-align: center;
    align-content: center;
`;
const ButtonFrame = styled.div`
    margin-top: 15px;
    height: 10%;
    text-align: end;
    align-content: flex-end;
`;
const H1 = styled.h1`
    color: ${COLORS.CAPTION_FONT_COLOR};
    margin: 5px 0 5px 0;
`;
const Span = styled.span`
    color: ${COLORS.ACCENT_FONT_PINK};
`;

const gameSetHandler = () => {
    globalThis.location.reload();
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
    const strike = log &&
                   betMonster &&
                   log.WinnerMonsterId === betMonster.MonsterId;

    return (
        <>
        {log &&
            <>
            <Line>{ log.ExistWinner ? DECO.BLOCK_LINE : "" }</Line>
            <Line>{ log.ExistWinner ? DECO.BLOCK_LINE_R : "" }</Line>
            <H1>
                {
                    log.ExistWinner ? `Winner: ${log.WinnerMonsterName} !!` : ""
                }
            </H1>
            <Line>{ log.ExistWinner ? DECO.BLOCK_LINE : "" }</Line>
            <Line>{ log.ExistWinner ? DECO.BLOCK_LINE_R : "" }</Line>

            <h1>{log.AllLoser ? "draw ..." : ""}</h1>

            <BetResultFrame>
                <h2>(∩´∀｀)∩<Span>獲得賞金&emsp;</Span>
                {
                    strike ? <Span>{Math.trunc(betGil * betMonster.BetRate)}</Span>
                           : <Span>0</Span>
                } Gil
                </h2>
            </BetResultFrame>

            <MessageFrame>
                {
                    newShops.length > 0 ? <h3 style={{margin: 0}}>新しいショップが解禁されました。</h3>
                                        : ""
                }
                {
                    newShops.map((shop, index) => (
                        <CommonStrong key={shop.ShopName + index}>
                            {shop.ShopName}&emsp;
                        </CommonStrong>
                    ))
                }
            </MessageFrame>

            <ButtonFrame>
                <CommonButton text={"終了"} width={120} onClick={gameSetHandler}/>
            </ButtonFrame>
            </>
        }
        </>
    );
};
export default BattleResultContentsBlock;