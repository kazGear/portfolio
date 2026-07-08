import { EditMonsterDTO } from "../../../types/Edit";
import CommonBorderTd from "../../common/CommonBorderTd";
import CommonInput from "../../common/CommonInput";

interface ArgProps {
    monster: EditMonsterDTO;
}

const EditMonsterSpeedBlock = ({monster}: ArgProps) => {
    return (
        <CommonBorderTd>
            <CommonInput styleObj={{width: "50px", textAlign: "right"}}
                    inputType="number"
                    defaultValue={monster.Speed}
                    labelTitle=""
                    onChange={(e: React.ChangeEvent<HTMLInputElement> | undefined) => {
                        monster.AfterSpeed = parseInt(e!.target.value);
                        monster.IsChanged = true;
                    }}/>
        </CommonBorderTd>
    );
}

export default EditMonsterSpeedBlock;