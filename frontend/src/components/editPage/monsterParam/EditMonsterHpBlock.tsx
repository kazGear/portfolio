import { EditMonsterDTO } from "../../../types/Edit";
import BorderTd from "../../common/BorderTd";
import Input from "../../common/Input";

interface ArgProps {
    monster: EditMonsterDTO;
}

const EditMonsterHpBlock = ({monster}: ArgProps) => {
    return (
        <>
            <BorderTd>
                <Input styleObj={{width: "50px"}}
                       inputType="number"
                       defaultValue={monster.Hp}
                       labelTitle=""
                       onChange={(e: React.ChangeEvent<HTMLInputElement> | undefined) => {
                            monster.AfterHp = parseInt(e!.target.value);
                            monster.IsChanged = true;
                       }}/>
            </BorderTd>
        </>
    );
}

export default EditMonsterHpBlock;