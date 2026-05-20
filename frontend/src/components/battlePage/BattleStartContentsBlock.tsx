import styled from "styled-components";
import Button from "../common/Button";
import { COLORS } from "../../lib/Constants";
import MonsterCountSelectorBlock from "./MonsterCountSelectorBlock";


const SdivButtonFrame = styled.div`
    height: 50%;
    align-content: flex-end;
    text-align: end;
`;
const Sh1 = styled.h1`
    margin: 0;
    color: ${COLORS.CAPTION_FONT_COLOR};
`;

interface ArgProps {
    battleStartHandler: (e: any) => Promise<void>;
    selectMonstersCountHandler: (e: any) => void;
}

const BattleStartContentsBlock = ({
    battleStartHandler, selectMonstersCountHandler}: ArgProps
) => {
    return (
        <div style={{
            display:  "block",
            height: "100%"}}
        >
            <div style={{ height: "50%" }}>
                <Sh1>モンスタ－闘技場</Sh1>
                <p style={{ marginTop: 0 }}>
                    参戦モンスター数を選択してください
                </p>
                <MonsterCountSelectorBlock
                    selectMonstersCountHandler={selectMonstersCountHandler}
                />
            </div>
            <SdivButtonFrame>
                <Button
                    id="battleStartBtn"
                    text={"次へ"}
                    width={150}
                    onClick={battleStartHandler}
                />

            </SdivButtonFrame>
        </div>
    );
}
export default BattleStartContentsBlock;