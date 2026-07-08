import { EditMonsterDTO } from "../../../types/Edit";
import CommonBorderTd from "../../common/CommonBorderTd";
import CommonInput from "../../common/CommonInput";

interface ArgProps {
    monster: EditMonsterDTO;
}

const EditMonsterNameBlock = ({monster}: ArgProps) => {
    return (
        <CommonBorderTd>
            <CommonInput styleObj={{width: "120px"}}
                    inputType="text"
                    defaultValue={monster.MonsterName}
                    labelTitle=""
                    onChange={(e: React.ChangeEvent<HTMLInputElement> | undefined) => {
                        monster.AfterName = e!.target.value;
                        monster.IsChanged = true;
                    }}/>
        </CommonBorderTd>
    );
}

export default EditMonsterNameBlock;