import Input from "../common/Input";

interface ArgProps {
    setInputLoginId: React.Dispatch<React.SetStateAction<string>>;
    setInputPassword: React.Dispatch<React.SetStateAction<string>>
}

const InputBlock = ({setInputLoginId, setInputPassword}: ArgProps) => {
    return (
        <div>
            <Input
                labelTitle="ログインID"
                inputType="text"
                id="loginId"
                name="loginId"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => setInputLoginId(e.target.value)}
                />
            <Input
                labelTitle="パスワード"
                inputType="password"
                id="password"
                name="password"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => setInputPassword(e.target.value)}
                />
        </div>
    );
}

export default InputBlock;