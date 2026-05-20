import Button from "../common/Button";

interface ArgProps {
    loginHandler: () => void;
    setShowRegistForm: React.Dispatch<React.SetStateAction<boolean>>;
    showRegistForm: boolean;
}

const ButtonBlock = ({loginHandler, setShowRegistForm, showRegistForm}: ArgProps) => {
    return (
        <div>
            <Button text="ログイン" onClick={loginHandler}/>
            <br/>
            <Button text="ユーザー登録がお済でない方"
                    onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => setShowRegistForm(!showRegistForm)}
                    styleObj={{width: "220px", marginTop: "15px"}}
            />
        </div>
    );
}

export default ButtonBlock;