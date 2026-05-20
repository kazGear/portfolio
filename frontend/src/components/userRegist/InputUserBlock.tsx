import Input from "../common/Input";


interface ArgProps {
    // イベントハンドラ
    inputHandlerLoginId: React.ChangeEventHandler<HTMLInputElement> | undefined;
    inputHandlerPassword: React.ChangeEventHandler<HTMLInputElement> | undefined;
    inputHandlerDispName: React.ChangeEventHandler<HTMLInputElement> | undefined;
    inputHandlerDispShortName: React.ChangeEventHandler<HTMLInputElement> | undefined;
    // 入力値が使用可能か
    usableLoginId: boolean;
    usablePassword: boolean;
    usableDispName: boolean;
    usableDispShortName: boolean;
    // 入力欄への参照
    refLoginId: React.LegacyRef<HTMLInputElement>;
    refPassword: React.LegacyRef<HTMLInputElement>;
    refDispName: React.LegacyRef<HTMLInputElement>;
    refDispShortName: React.LegacyRef<HTMLInputElement>;
}

const InputUserBlock = ({
    inputHandlerLoginId,
    inputHandlerPassword,
    inputHandlerDispName,
    inputHandlerDispShortName,
    usableLoginId,
    usablePassword,
    usableDispName,
    usableDispShortName,
    refLoginId,
    refPassword,
    refDispName,
    refDispShortName
}: ArgProps
) => {

    return (
        <div>
            <Input
                id="registLoginId"
                labelTitle="ログインID"
                inputType="text"
                onChange={inputHandlerLoginId}
                showMessage={usableLoginId}
                alertMessage="既に使用されています。"
                ref={refLoginId}
            />

            <Input
                id="registPassword"
                labelTitle="パスワード"
                inputType="password"
                onChange={inputHandlerPassword}
                showMessage={usablePassword}
                alertMessage="任意の４文字以上で設定してください。"
                ref={refPassword}
            />

            <Input
                id="registDispName"
                labelTitle="ユーザー名"
                inputType="text"
                onChange={inputHandlerDispName}
                showMessage={usableDispName}
                alertMessage="既に使用されている名称です。"
                ref={refDispName}
            />

            <Input
                id="registDispShortName"
                labelTitle="ユーザー略称"
                inputType="text"
                onChange={inputHandlerDispShortName}
                showMessage={usableDispShortName}
                alertMessage="既に使用されている略称です。"
                ref={refDispShortName}
            />
        </div>
    );
}

export default InputUserBlock;