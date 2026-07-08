import { useCallback, useEffect, useState } from "react";
import { KEYS, URLS } from "../../../lib/Constants";
import MonsterTableBody from "./MonsterTableBody";
import styled from "styled-components";
import { EditMonsterDTO } from "../../../types/Edit";
import EditStatusFinishedDialog from "./EditStatusFinishedDialog";
import CommonHeaderOfBody from "../../common/CommonHeaderOfBody";
import { api } from "../../../lib/apiClient";

const Table = styled.table`
    margin: auto;
    width: 90%;
`;

interface ArgProps {
    editMonsters: EditMonsterDTO[];
    setEditMonsters: React.Dispatch<React.SetStateAction<EditMonsterDTO[]>>;
}

const MonsterStatusEditBlock = ({editMonsters, setEditMonsters}: ArgProps) => {
    const [showDialog, setShowDialog] = useState(false);
    const [isNowLoading, setIsNowLoading] = useState(true);

    /**
     * 編集用モンスター情報
     */
    useEffect(() => {
        const fetchEditMonsters = async () => {
            const loginId: string | null = localStorage.getItem(KEYS.USER_ID);

            const monsters = await api.POST<EditMonsterDTO[]>(URLS.FETCH_EDIT_MONSTERS, loginId);

            setEditMonsters([...monsters!]);
            setIsNowLoading(false);
        }
        fetchEditMonsters();
    }, []);
    /**
     * 更新実行
     */
    const updateStatusHandler = useCallback( async () => {
        await api.PUT(URLS.UPDATE_MONSTER_STATUS, editMonsters);

        setEditMonsters([...editMonsters]);
        setShowDialog(true);
    }, [editMonsters]);

    return (
        <div>
            <CommonHeaderOfBody message="変更したいパラメータを入力してください。"
                                buttonText="ステータス変更"
                                callback={updateStatusHandler}/>
            {/* ステータス編集部 */}
            <Table>
                <MonsterTableBody editMonsters={editMonsters}
                                  isNowLoading={isNowLoading}/>
            </Table>
            {/* 完了ダイアログ */}
            <EditStatusFinishedDialog isShow={showDialog}
                                      setShowDialog={setShowDialog}
                                      setEditMonsters={setEditMonsters}/>
        </div>
    );
}

export default MonsterStatusEditBlock;