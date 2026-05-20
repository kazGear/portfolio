import styled from "styled-components";
import { COLORS, KEYS } from "../lib/Constants";
import DialogFrame from "../components/common/DialogFrame";
import { useCallback, useLayoutEffect, useState } from "react";
import { useLogin } from "../hooks/useHooksOfUser";
import UserRegistBlock from "../components/loginPage/UserRegistBlock";
import OutSideFrame from "../components/common/OutSideFrame";
import InputBlock from "../components/loginPage/InputBlock";
import ButtonBlock from "../components/loginPage/ButtonBlock";

const LoginContainer = styled.div`
    text-align: center;
    display: flex;
`;

const frameStyle: React.CSSProperties = {
    width: "50%",
    minWidth: "300px",
    height: "200px",
    margin: "auto",
    alignContent: "center",
    textAlign: "center",
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
};

const Sspan = styled.span`
    color: ${COLORS.ALERT_MESSAGE_COLOR};
    font-size: 12px;
`;

const LoginPage = () => {
    const [inputLoginId, setInputLoginId] = useState("");
    const [inputPassword, setInputPassword] = useState("");
    const [showRegistForm, setShowRegistForm] = useState(false);
    const [token, setToken] = useState<string | null>(null);
    const [showAlert, setShowAlert] = useState(false);

    /**
     * 初期処理
     */
    useLayoutEffect(() => {
        const token = localStorage.getItem(KEYS.TOKEN);
        if (token) setToken(token);
    }, []);
    /**
     * ログイン処理
     */
    const login = useLogin();
    const loginHandler = useCallback(() => {
        login({inputLoginId, inputPassword, setToken, setShowAlert});
    }, [inputLoginId, inputPassword]);

    return (
        <>
            <LoginContainer>
                <OutSideFrame styleObj={frameStyle} >
                    <form action="" method="post">
                        <InputBlock setInputLoginId={setInputLoginId}
                                    setInputPassword={setInputPassword}
                         />
                        {
                            showAlert ? <Sspan>入力に誤りがあるか、無効なユーザーです。</Sspan> : ""
                        }
                        <br />
                        <ButtonBlock loginHandler={loginHandler}
                                     setShowRegistForm={setShowRegistForm}
                                     showRegistForm={showRegistForm}/>
                    </form>
                </OutSideFrame>
            </LoginContainer>

            {/* ユーザ登録フォーム */}
            <DialogFrame showDialog={showRegistForm}>
                <UserRegistBlock setShowRegistForm={setShowRegistForm}/>
            </DialogFrame>
         </>
    );
};

export default LoginPage;
