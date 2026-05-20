import styled from "styled-components";
import { KEYS, URLS } from "../lib/Constants";
import { useEffect, useState } from "react";
import { useServerWithQuery } from "../hooks/useHooksOfCommon";
import { ItemDTO } from "../types/Shop";
import SelectShops from "../components/shopPage/SelectShops";
import ShopItemTable from "../components/shopPage/ShopItemTable";
import { UserDTO } from "../types/UserManage";
import PurchaseDialog from "../components/shopPage/PurchaseDialog";

const SshopPageFrame = styled.div`
    display: flex;
    height: 90%;
`;
const SdivControllerFrame = styled.div`
    width:25%;
    margin: 20px;
`;
const SdivItemFrame = styled.div`
    width: 75%;
    margin: 20px;
    overflow: overlay;
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

    /**
     * 店舗アイテム表示
     */
    const select = useServerWithQuery();
    useEffect(() => {
        const fetchShopItems = async () => {
            const loginId: string | null = localStorage.getItem(KEYS.USER_ID);

            const items: ItemDTO[] = await select(
                URLS.SELECT_SHOP_ITEMS + `?loginId=${loginId}&shopId=${selectedShop}`);
            const loginUser: UserDTO = await select(
                URLS.USER_INFO + `?loginId=${loginId}`);

            setShopItems(items);
            setUser(loginUser);
            setMyCash(loginUser.Cash);
        }
        fetchShopItems();
    }, [selectedShop]);

    return (
        <SshopPageFrame>
            <SdivControllerFrame>
                {/* コントローラ */}
                <SelectShops setSelectedShop={setSelectedShop}
                             user={user}
                             myCash={myCash}/>
            </SdivControllerFrame>

            <SdivItemFrame>
                {/* 販売商品テーブル */}
                <ShopItemTable shopItems={shopItems}
                               user={user}
                               myCash={myCash}
                               setMyCash={setMyCash}
                               setPurchaseItem={setPurchaseItem}
                               setShowPurchaseDialog={setShowPurchaseDialog}/>
            </SdivItemFrame>
            {/* 購入済ダイアログ */}
            <PurchaseDialog showDialog={showPurchaseDialog}
                            purchaseItem={purchaseItem}
                            setShowPurchaseDialog={setShowPurchaseDialog}/>
        </SshopPageFrame>
    )
};

export default ShopPage;
