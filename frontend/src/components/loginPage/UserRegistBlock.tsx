import { useEffect, useLayoutEffect, useRef, useState } from "react";
import Button from "../common/Button";
import { useServerWithQuery } from "../../hooks/useHooksOfCommon";
import { COLORS, KEYS, URLS } from "../../lib/Constants";
import { UserDTO } from "../../types/UserManage";
import { useCreateUsedList } from "../../hooks/useHooksOfUser";
import InputUserContents from "../userRegist/InputUserBlock";
import styled from "styled-components";

const SbuttonFrame = styled.div`
    height: 20%;
    min-height: 60px;
    text-align: end;
    align-content: flex-end;
`;

interface ArgProps {
    setShowRegistForm: React.Dispatch<React.SetStateAction<boolean>>
}

const UserRegistBlock = ({setShowRegistForm}: ArgProps) => {
    // ユーザ関連情報
    const [users, setUserList] = useState<UserDTO[] | null>(null);
    const [usedLoginIdList, setUsedLoginIdList] = useState<string[] | null>(null);
    const [usedDispNameList, setUsedDispNameList] = useState<string[] | null>(null);
    const [usedDispShortNameList, setUsedDispShortNameList] = useState<string[] | null>(null);
    // 入力値
    const [inputLoginId, setInputLoginId] = useState("");
    const [inputPassword, setInputPassword] = useState("");
    const [inputDispName, setInputDispName] = useState("");
    const [inputDispShortName, setInputDispShortName] = useState("");
    // 入力値が使用可能か、既に登録済でないか
    const [usableLoginId, setUsableLoginId] = useState(true);
    const [usablePassword, setUsablePassword] = useState(false);
    const [usableDispName, setCanUseDispName] = useState(true);
    const [usableDispShortName, setCanUseDispShortName] = useState(true);
    // 入力欄参照
    const refLoginId = useRef<HTMLInputElement>(null!);
    const refPassword = useRef<HTMLInputElement>(null!);
    const refDispName = useRef<HTMLInputElement>(null!);
    const refDispShortName = useRef<HTMLInputElement>(null!);

    const [canRegist, setCanRegist] = useState(true);
    const [registResult, setRegistResult] = useState("");

    /**
     * 既に登録されているログインID等と取得（検証に使用）
     */
    const goToServer = useServerWithQuery();
    useLayoutEffect(() => {
        const fetchUsers = async () => {
            const usersFromDb: UserDTO[] = await goToServer(`${URLS.REGIST_USER_INIT}`);
            setUserList(usersFromDb);
        };
        fetchUsers();
    }, []);
    /**
     * 既に登録されているデータ群
     */
    useCreateUsedList({
        users,
        setUsedLoginIdList,
        setUsedDispNameList,
        setUsedDispShortNameList
    });
    /**
     * 登録可能な状態か
     */
    useEffect(() => {
        if (   usableLoginId
            && usablePassword
            && usableDispName
            && usableDispShortName
        ) {
            setCanRegist(false);
        } else {
            setCanRegist(true);
        }
    }, [usableLoginId, usablePassword, usableDispName, usableDispShortName]);
    /**
     * 登録内容の送信
     */
    const insertUser = useServerWithQuery();
    const exeUserInsert = async () => {
        try {
            await insertUser(`${URLS.REGIST_USER}?LoginId=${inputLoginId}
                                                &Password=${inputPassword}
                                                &DispName=${inputDispName}
                                                &DispShortName=${inputDispShortName}`);
            localStorage.setItem(KEYS.USER_ID, inputLoginId);
            setRegistResult("正常に登録されました。");
            setCanRegist(true);
        } catch (err) {
            localStorage.removeItem(KEYS.USER_ID);
            setRegistResult("登録に失敗しました。");
        }
    };
    /**
     * ログインID用ハンドラ
     */
    const inputHandlerLoginId = (e: React.ChangeEvent<HTMLInputElement>) => {
        const usable = usedLoginIdList?.includes(e.target.value);
        setUsableLoginId(!usable);
        setInputLoginId(e.target.value);
    };
    /**
     * パスワード用ハンドラ
     */
    const inputHandlerPassword = (e: React.ChangeEvent<HTMLInputElement>) => {
        const match: any = e.target.value.match(/.{4,}/g); // 何かしら４文字以上必須
        setUsablePassword(match);
        setInputPassword(e.target.value);
    };
    /**
     * ユーザ名用ハンドラ
     */
    const inputHandlerDispName = (e: React.ChangeEvent<HTMLInputElement>) => {
        const usable = usedDispNameList?.includes(e.target.value);
        setCanUseDispName(!usable);
        setInputDispName(e.target.value);
    };
    /**
     * ユーザ略称用ハンドラ
     */
    const inputHandlerDispShortName = (e: React.ChangeEvent<HTMLInputElement>) => {
        const usable = usedDispShortNameList?.includes(e.target.value);
        setCanUseDispShortName(!usable);
        setInputDispShortName(e.target.value);
    };

    return (
        <div>
            <h1>ユーザー登録</h1>
            <InputUserContents inputHandlerLoginId={inputHandlerLoginId}
                               inputHandlerPassword={inputHandlerPassword}
                               inputHandlerDispName={inputHandlerDispName}
                               inputHandlerDispShortName={inputHandlerDispShortName}
                               usableLoginId={usableLoginId}
                               usablePassword={usablePassword}
                               usableDispName={usableDispName}
                               usableDispShortName={usableDispShortName}
                               refLoginId={refLoginId}
                               refPassword={refPassword}
                               refDispName={refDispName}
                               refDispShortName={refDispShortName}
                               />
            <div style={{height: "20px"}}>
            {
                registResult === "" ?  "" : <span style={{color: `${COLORS.ALERT_MESSAGE_COLOR}`}}>
                                                {registResult}
                                            </span>
            }
            </div>
            <SbuttonFrame>
                <Button text="閉じる"
                        onClick={() => {
                            setShowRegistForm(false);
                            setRegistResult("");
                            // 入力値初期化
                            refLoginId.current.value = "";
                            refPassword.current.value = "";
                            refDispName.current.value = "";
                            refDispShortName.current.value = "";
                        }}
                />
                <Button text="登録"
                        disabled={canRegist}
                        onClick={exeUserInsert} />
            </SbuttonFrame>
        </div>
    );
}
export default UserRegistBlock;