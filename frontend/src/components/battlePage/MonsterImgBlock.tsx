import { useEffect, useRef, useState } from "react";
import monsterImages from "../../lib/MonsterImages";
import styled from "styled-components";
import effectImages from "../../lib/effectImages"
import { DAMAGE_VIEW } from "../../lib/Constants";
import { isEmpty } from "../../lib/CommonLogic";
import { MetaDataDTO, MonsterDTO } from "../../types/MonsterBattle";

const SdivMonsterImgFrame = styled.div`
    margin: 10px 0 5px 0;
`;
const SdivMonsterImg = styled.div`
    position: relative;
`;

const SimgMonster = styled.img`
    position: relative;
    width: 100px;
    height: 100px;
    border-radius: 100%;
`;
interface SimgEffectProp { display: string; }
const SimgEffect = styled.img<SimgEffectProp>`
    display: ${props => props.display};
    position: absolute;
    left: 0;
    width: 100%;
    height: 100%;
`;

interface ArgProps {
    monster: MonsterDTO;
    shortLog: MetaDataDTO[];
}

const MonsterImgBlock = ({monster, shortLog}: ArgProps) => {
    const imgRef = useRef<HTMLImageElement>(null);
    const [effectImage, setEffectImage] = useState("");
    const [showEffect, setShowEffect] = useState(false);
    const monsterImage = monsterImages(monster.MonsterId);
    // モンスターの動き
    const [monsterMove, setMonsterMove] = useState({});
    const noMove = {animation: ""}
    const damageFlash = {animation: "0.8s damageFlash"};
    const dodge = {animation: `${DAMAGE_VIEW.DODGE_END / 1000}s dodge`}

    // ダメージ表現
    useEffect(() => {
        for (const log of shortLog) {
            if (   log.TargetMonsterId === monster.MonsterId
                && !isEmpty(log.SkillId)) // 受けるスキルがある
            {
                setEffectImage(effectImages(log.SkillId));
                setShowEffect(true);

                if (log.IsDodge) { // 回避した
                    setTimeout(() => setShowEffect(false), log.EffectTime);
                    setTimeout(() => setMonsterMove(dodge), DAMAGE_VIEW.DODGE_START);
                    setTimeout(() => setMonsterMove(noMove), DAMAGE_VIEW.DODGE_END);
                } else { // ヒットした
                    setTimeout(() => setShowEffect(false), log.EffectTime);
                    setTimeout(() => setMonsterMove(damageFlash), log.EffectTime);
                    setTimeout(() => setMonsterMove(noMove), DAMAGE_VIEW.DAMAGE_END);
                }
            }
        }
    }, [shortLog]);

    return (
        <SdivMonsterImgFrame>
            <SdivMonsterImg>
                <SimgMonster
                    id={"monsterImage" + monster.MonsterId}
                    src={monsterImage}
                    alt="Loding ..."
                    ref={imgRef}
                    style={monsterMove}
                    />
                <SimgEffect
                    src={effectImage}
                    alt="Loding ..."
                    display={showEffect ? "inline" : "none"}
                    />
            </SdivMonsterImg>
        </SdivMonsterImgFrame>
    );
}

export default MonsterImgBlock;