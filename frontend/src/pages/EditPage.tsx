import styled from "styled-components";
import EditSelector from "../components/editPage/monsterParam/EditSelectorBlock";
import EditMonsterStatusBlock from "../components/editPage/monsterParam/EditMonsterStatusBlock";
import { useState } from "react";
import OutSideFrame from "../components/common/OutSideFrame";
import { EditMonsterDTO, EditSkillsDTO } from "../types/Edit";
import ClearAllSkillsBlock from "../components/editPage/monsterSkills/ClearAllSkillsBlock";
import ClearAllStatusBlock from "../components/editPage/monsterParam/ClearAllStatusBlock";
import EditMonsterSkillsBlock from "../components/editPage/monsterSkills/EditMonsterSkillsBlock";
import { useCheckToken } from "../hooks/useHooksOfCommon";

const SdivEditFrame = styled.div`
    width: 90%;
    margin: auto;
`;
const styleForController = {
    width: "100%",
    marginBottom: "0",
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center"
}
const styleForEdit = {
    height: "75vh"
}

const EditPage = () => {
    const [editMonsters, setEditMonsters] = useState<EditMonsterDTO[]>([]);
    const [editMonsterSkills, setEditMonsterSkills] = useState<EditSkillsDTO[]>([]);
    const [selectEditType, setSelectEditType] = useState(1);

    useCheckToken();

    return (
        <SdivEditFrame>
            <OutSideFrame styleObj={styleForController}>
                {/* 設定種類コントローラ */}
                <EditSelector setSelectEditType={setSelectEditType}/>
                <ClearAllStatusBlock setEditMonsters={setEditMonsters} selectEditType={selectEditType}/>
                <ClearAllSkillsBlock setEditMonsterSkills={setEditMonsterSkills} selectEditType={selectEditType}/>
            </OutSideFrame>
            <OutSideFrame styleObj={styleForEdit}>
                {/* 設定部 */}
                { selectEditType === 1 ? <EditMonsterStatusBlock editMonsters={editMonsters} setEditMonsters={setEditMonsters} /> : ""}
                { selectEditType === 2 ? <EditMonsterSkillsBlock editMonsterSkills={editMonsterSkills} setEditMonsterSkills={setEditMonsterSkills}/> : ""}
                { selectEditType === 3 ? <h1 style={{marginLeft: "40px"}}>使用可能モンスター（工事予定）</h1> : ""}
            </OutSideFrame>
        </SdivEditFrame>
    );
}

export default EditPage;