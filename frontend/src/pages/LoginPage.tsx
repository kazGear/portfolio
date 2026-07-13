import styled from "styled-components";
import { COLORS, KEYS } from "../lib/Constants";
import CommonDialogFrame from "../components/common/CommonDialogFrame";
import { useCallback, useEffect, useState } from "react";
import { useLogin } from "../hooks/useHooksOfUser";
import UserRegistBlock from "../components/loginPage/UserRegistBlock";
import CommonFrame from "../components/common/CommonFrame";
import InputBlock from "../components/loginPage/InputBlock";
import ButtonBlock from "../components/loginPage/ButtonBlock";

const LoginFrame = styled.div`
    text-align: center;
    display: flex;
`;

const frameStyle: React.CSSProperties = {
    width: "40%",
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

const Span = styled.span`
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
    useEffect(() => {
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
            <LoginFrame>
                <CommonFrame styleObj={frameStyle} >
                    <form action="" method="post">
                        <InputBlock setInputLoginId={setInputLoginId}
                                    setInputPassword={setInputPassword}
                         />
                        {
                            showAlert ? <Span>入力に誤りがあるか、無効なユーザーです。</Span> : ""
                        }
                        <br />
                        <ButtonBlock loginHandler={loginHandler}
                                     setShowRegistForm={setShowRegistForm}
                                     showRegistForm={showRegistForm}/>
                    </form>
                </CommonFrame>
            </LoginFrame>

            {/* ユーザ登録フォーム */}
            <CommonDialogFrame showDialog={showRegistForm}>
                <UserRegistBlock setShowRegistForm={setShowRegistForm}/>
            </CommonDialogFrame>
         </>
    );
};

export default LoginPage;
