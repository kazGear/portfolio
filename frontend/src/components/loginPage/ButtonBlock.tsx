import CommonButton from "../common/CommonButton";

interface ArgProps {
    loginHandler: () => void;
    setShowRegistForm: React.Dispatch<React.SetStateAction<boolean>>;
    showRegistForm: boolean;
}

const ButtonBlock = ({loginHandler, setShowRegistForm, showRegistForm}: ArgProps) => {
    return (
        <div>
            <CommonButton text="ログイン" onClick={loginHandler}/>
            <br/>
            <CommonButton text="ユーザー登録がお済でない方"
                    onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => setShowRegistForm(!showRegistForm)}
                    styleObj={{width: "220px", marginTop: "15px"}}
            />
        </div>
    );
}

export default ButtonBlock;