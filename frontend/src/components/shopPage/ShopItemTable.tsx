import styled from "styled-components";
import { ItemDTO } from "../../types/Shop";
import CommonFrame from "../common/CommonOutSideFrame";
import BorderTd from "../common/CommonBorderTd";
import React from "react";
import { UserDTO } from "../../types/UserManage";
import ShopTableHeader from "./ShopTableHeader";
import ItemImage from "./ItemImage";
import ItemRemarks from "./ItemRemarks";
import ItemPrice from "./ItemPrice";
import PurchaseButton from "./PurchaseButton";

const Table = styled.table`
    margin: auto;
    width: 95%;
    border-collapse: collapse;
    margin-bottom: 20px;
`;

interface ArgProps {
    shopItems: ItemDTO[];
    user: UserDTO | null;
    myCash: number | null;
    setMyCash: React.Dispatch<React.SetStateAction<number | null>>;
    setPurchaseItem: React.Dispatch<React.SetStateAction<string>>;
    setShowPurchaseDialog: React.Dispatch<React.SetStateAction<boolean>>;
}

const ShopItemTable = ({
    shopItems,
    user,
    myCash,
    setMyCash,
    setPurchaseItem,
    setShowPurchaseDialog
}: ArgProps) => {

    return (
        <CommonFrame>
            <p style={{margin: "20px"}}>取り扱い商品</p>

            <Table>
                <ShopTableHeader/>
                <tbody>
                {
                    shopItems.map((item, index) => {
                        return (
                            <tr key={index} style={{height: "60px"}}>
                                <ItemImage item={item}/>
                                <BorderTd>{item.ItemName}</BorderTd>
                                <ItemRemarks item={item}/>
                                <ItemPrice item={item} myCash={myCash}/>
                                <PurchaseButton item={item}
                                                myCash={myCash}
                                                setMyCash={setMyCash}
                                                setPurchaseItem={setPurchaseItem}
                                                setShowPurchaseDialog={setShowPurchaseDialog}/>
                            </tr>
                        )
                    })
                }
                </tbody>
            </Table>
        </CommonFrame>
    );
}

export default ShopItemTable;