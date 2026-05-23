import Strong from "../../common/Strong";

const MonsterTableHeader = () => {
    return (
        <table style={{width: "620px"}}>
            <thead>
                <tr style={{textAlign: "center"}}>
                    <td style={{width: "18px", paddingLeft: "35px"}}>
                        <Strong>ID</Strong>
                    </td>
                    <td style={{width: "64px", paddingLeft: "30px"}}>
                        <Strong>イメージ</Strong>
                    </td>
                    <td style={{width: "96px", paddingLeft: "20px"}}>
                        <Strong>モンスター名</Strong>
                    </td>
                    <td style={{width: "24px", paddingLeft: "45px"}}>
                        <Strong>HP</Strong>
                    </td>
                    <td style={{width: "48px", paddingLeft: "40px"}}>
                        <Strong>攻撃力</Strong>
                    </td>
                    <td style={{width: "32px", paddingLeft: "30px"}}>
                        <Strong>速さ</Strong>
                    </td>
                    <td style={{width: "32px", paddingLeft: "15px"}}>
                        <Strong>弱点</Strong>
                    </td>
                    <td> ⇒ 変更後</td>
                </tr>
            </thead>
        </table>
    );
}

export default MonsterTableHeader;