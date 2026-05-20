import { EditMonsterDTO } from "../../../types/Edit";
import BorderTd from "../../common/BorderTd";
import Input from "../../common/Input";

interface ArgProps {
    monster: EditMonsterDTO;
}

const EditMonsterAttackBlock = ({monster}: ArgProps) => {
    return (
        <>
            <BorderTd>
                <Input styleObj={{width: "50px"}}
                       inputType="number"
                       defaultValue={monster.Attack}
                       labelTitle=""
                       onChange={(e: React.ChangeEvent<HTMLInputElement> | undefined) => {
                            monster.AfterAttack = parseInt(e!.target.value);
                            monster.IsChanged = true;
                       }}/>
            </BorderTd>
        </>
    );
}

export default EditMonsterAttackBlock;