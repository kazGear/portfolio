import { useLayoutEffect, useState } from "react";
import { useServer } from "../../../hooks/useHooksOfCommon";
import { URLS } from "../../../lib/Constants";
import { AllSkillDTO, EditSkillsDTO } from "../../../types/Edit";

import MonsterSkillBlock from "./MonsterSkillBlock";
import MonsterStatusBlock from "./MonsterStatusBlock";
import NowLoading from "../../common/NowLoading";

interface ArgProps {
    editMonsterSkills: EditSkillsDTO[];
    isNowLoading: boolean | null;
}

const EditSkillBlock = ({editMonsterSkills, isNowLoading}: ArgProps) => {
    const [allSkills, setAllSkills] = useState<AllSkillDTO[]>([]);
    /**
     * スキルリストを取得
     */
    const goToServer = useServer();
    useLayoutEffect(() => {
        const fetchAllSkills = async () => {
            const allSkills: AllSkillDTO[] = await goToServer(URLS.FETCH_ALL_SKILLS);
            setAllSkills(allSkills);
        }
        fetchAllSkills();
    }, []);
    /**
     * ローディング中
     */
    if (isNowLoading) {
        return (
            <div style={{margin: "100px"}}>
                <NowLoading alt="ローディング"/>
            </div>
        );
    }

    return (
        <div>
            {
                editMonsterSkills.map((monster, index) => (
                    <div key={index}>
                        <MonsterStatusBlock monster={monster}/>
                        <MonsterSkillBlock monster={monster}
                                           allSkills={allSkills}
                                          />
                        <hr/>
                    </div>
                ))
            }
        </div>
    );
}

export default EditSkillBlock;