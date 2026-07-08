import styled from "styled-components";
import { KEYS, URLS } from "../lib/Constants";
import { useEffect, useState } from "react";
import { ItemDTO } from "../types/Shop";
import SelectShops from "../components/shopPage/SelectShops";
import ShopItemTable from "../components/shopPage/ShopItemTable";
import { UserDTO } from "../types/UserManage";
import PurchaseDialog from "../components/shopPage/PurchaseDialog";
import { api } from "../lib/apiClient";
import { useCheckToken } from "../hooks/useHooksOfCommon";

const ShopPageFrame = styled.div`
    display: flex;
    height: 90%;
`;
const ControllerFrame = styled.div`
    width: 25%;
    margin: 0px 20px;
`;
const ItemFrame = styled.div`
    width: 75%;
    height: 90vh;
    margin: 0px 20px 0px 0px;
    overflow: scroll;
`;

const ShopPage = () => {
    // ショップ関係
    const [selectedShop, setSelectedShop] = useState<string | undefined>("shop001");
    const [shopItems, setShopItems] = useState<ItemDTO[]>([]);
    const [purchaseItem, setPurchaseItem] = useState("");
    // ユーザー関係
    const [user, setUser] = useState<UserDTO | null>(null);
    const [myCash, setMyCash] = useState<number | null>(null);

    const [showPurchaseDialog, setShowPurchaseDialog] = useState(false);

    useCheckToken();

    /**
     * 店舗アイテム表示
     */
    useEffect(() => {
        const fetchShopItems = async () => {
            const loginId: string | null = localStorage.getItem(KEYS.USER_ID);

            const items = await api.POST<ItemDTO[]>(URLS.SELECT_SHOP_ITEMS, {
                loginId:      `${loginId}`,
                selectedShop: `${selectedShop}`,
            });

            const loginUser = await api.POST<UserDTO>(URLS.USER_INFO, loginId);

            setShopItems(items!);
            setUser(loginUser);
            setMyCash(loginUser!.Cash);
        }
        fetchShopItems();
    }, [selectedShop]);

    return (
        <ShopPageFrame>
            <ControllerFrame>
                {/* コントローラ */}
                <SelectShops setSelectedShop={setSelectedShop}
                             user={user}
                             myCash={myCash}/>
            </ControllerFrame>

            <ItemFrame>
                {/* 販売商品テーブル */}
                <ShopItemTable shopItems={shopItems}
                               user={user}
                               myCash={myCash}
                               setMyCash={setMyCash}
                               setPurchaseItem={setPurchaseItem}
                               setShowPurchaseDialog={setShowPurchaseDialog}/>
            </ItemFrame>
            {/* 購入済ダイアログ */}
            <PurchaseDialog showDialog={showPurchaseDialog}
                            purchaseItem={purchaseItem}
                            setShowPurchaseDialog={setShowPurchaseDialog}/>
        </ShopPageFrame>
    )
};

export default ShopPage;
