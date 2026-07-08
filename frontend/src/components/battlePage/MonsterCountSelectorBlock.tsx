
import CommonSelect from "../common/CommonSelect";

interface ArgProps {
    selectMonstersCountHandler: (e: any) => void;
}

const MonsterCountSelectorBlock = ({selectMonstersCountHandler}: ArgProps) => {
    return (
        <form method="POST" onChange={selectMonstersCountHandler}>
            <CommonSelect>
                <option value="2">２匹戦</option>
                <option value="3">３匹戦</option>
                <option value="4">４匹戦</option>
                <option value="5">５匹戦</option>
                <option value="6">６匹戦</option>
            </CommonSelect>
        </form>
    );
}

export default MonsterCountSelectorBlock;