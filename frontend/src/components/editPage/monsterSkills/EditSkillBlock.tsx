import { useEffect, useState } from "react";
import { URLS } from "../../../lib/Constants";
import { AllSkillDTO, EditSkillsDTO } from "../../../types/Edit";
import MonsterSkillBlock from "./MonsterSkillBlock";
import MonsterStatusBlock from "./MonsterStatusBlock";
import NowLoading from "../../common/NowLoading";
import { api } from "../../../lib/apiClient";

interface ArgProps {
    editMonsterSkills: EditSkillsDTO[];
    isNowLoading: boolean | null;
}

const EditSkillBlock = ({editMonsterSkills, isNowLoading}: ArgProps) => {
    const [allSkills, setAllSkills] = useState<AllSkillDTO[]>([]);
    /**
     * スキルリストを取得
     */
    useEffect(() => {
        const fetchAllSkills = async () => {
            const allSkills = await api.GET<AllSkillDTO[]>(URLS.FETCH_ALL_SKILLS);
            setAllSkills(allSkills!);
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