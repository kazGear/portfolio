import styled from "styled-components";
import Select from "../common/Select";

const SdivFrame = styled.div`
    margin: 10px;
`;

interface ArgProps {
    changeBattleScaleHandler: React.ChangeEventHandler<HTMLSelectElement>;
}

const BattleScaleListBlock = ({changeBattleScaleHandler}: ArgProps) => {
    return (
        <>
            <Select title="対戦規模" onChange={changeBattleScaleHandler}>
                <option value="0">指定なし</option>
                <option value="2">２匹戦</option>
                <option value="3">３匹戦</option>
                <option value="4">４匹戦</option>
                <option value="5">５匹戦</option>
                <option value="6">６匹戦</option>
            </Select>
        </>
    );
};

export default BattleScaleListBlock;
