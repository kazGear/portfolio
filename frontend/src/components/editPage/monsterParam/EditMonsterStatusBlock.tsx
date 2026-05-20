import { useCallback, useLayoutEffect, useState } from "react";
import { useServerWithJson, useServerWithQuery } from "../../../hooks/useHooksOfCommon";
import { KEYS, URLS } from "../../../lib/Constants";
import MonsterTableHeader from "./MonsterTableHeader";
import MonsterTableBody from "./MonsterTableBody";
import styled from "styled-components";
import { EditMonsterDTO } from "../../../types/Edit";
import EditStatusFinishedDialog from "./EditStatusFinishedDialog";
import HeaderOfBody from "../../common/HeaderOfBody";

const Stable = styled.table`
    margin: auto;
    width: 90%;
`;

interface ArgProps {
    editMonsters: EditMonsterDTO[];
    setEditMonsters: React.Dispatch<React.SetStateAction<EditMonsterDTO[]>>;
}

const MonsterStatusEditBlock = ({editMonsters, setEditMonsters}: ArgProps) => {
    const [showDialog, setShowDialog] = useState(false);
    const [isNowLoading, setIsLowLoading] = useState(true);

    /**
     * 編集用モンスター情報
     */
    const goToServer = useServerWithQuery();
    useLayoutEffect(() => {
        const fetchEditMonsters = async () => {
            const loginId: string | null = localStorage.getItem(KEYS.USER_ID);
            const monsters: EditMonsterDTO[] = await goToServer(
                URLS.FETCH_EDIT_MONSTERS + `?loginId=${loginId}`
            );
            setEditMonsters([...monsters]);
            setIsLowLoading(false);
        }
        fetchEditMonsters();
    }, []);
    /**
     * 更新実行
     */
    const goToServerWithJson = useServerWithJson();
    const updateStatusHandler = useCallback(() => {
        goToServerWithJson(
            editMonsters, URLS.UPDATE_MONSTER_STATUS
        );
        setEditMonsters([...editMonsters]);
        setShowDialog(true);
    }, [editMonsters]);

    return (
        <div>
            <HeaderOfBody message="変更したいパラメータを入力してください。"
                          buttonText="ステータス変更"
                          callback={updateStatusHandler}/>
            {/* ステータス編集部 */}
            <Stable>
                <MonsterTableHeader />
                <MonsterTableBody editMonsters={editMonsters}
                                  isNowLoading={isNowLoading}/>
            </Stable>
            {/* 完了ダイアログ */}
            <EditStatusFinishedDialog isShow={showDialog}
                                      setShowDialog={setShowDialog}
                                      setEditMonsters={setEditMonsters}/>
        </div>
    );
}

export default MonsterStatusEditBlock;