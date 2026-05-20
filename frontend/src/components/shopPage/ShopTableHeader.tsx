const ShopTableHeader = () => {
    return (
        <>
            <thead>
                <tr>
                    <td style={{border: "none", minWidth: "70px"}}>
                        <span style={{marginLeft: "10px"}}>イメージ</span>
                    </td>
                    <td style={{border: "none", minWidth: "120px"}}>商品名</td>
                    <td style={{border: "none"}}>備考</td>
                    <td style={{border: "none"}}>価格</td>
                    <td style={{border: "none", width: "100px"}}> </td>
                </tr>
            </thead>
        </>
    );
}

export default ShopTableHeader;