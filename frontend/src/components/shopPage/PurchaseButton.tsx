import { ItemDTO } from "../../types/Shop";
import CommonBorderTd from "../common/CommonBorderTd";
import React, { useCallback } from "react";
import CommonButton from "../common/CommonButton";
import { KEYS, URLS } from "../../lib/Constants";
import { UserDTO } from "../../types/UserManage";
import { api } from "../../lib/apiClient";

interface ArgProps {
    item:                  ItemDTO;
    myCash:                number | null;
    setMyCash:             React.Dispatch<React.SetStateAction<number | null>>;
    setPurchaseItem:       React.Dispatch<React.SetStateAction<string>>;
    setShowPurchaseDialog: React.Dispatch<React.SetStateAction<boolean>>;
}

const PurchaseCommonButton = ({item,
                               myCash,
                               setMyCash,
                               setPurchaseItem,
                               setShowPurchaseDialog}: ArgProps
) => {
        /**
         * 購入処理
         */
        const purchase = useCallback((itemRow: ItemDTO) => {
            const update = async (itemRow: ItemDTO) => {
                const loginId: string | null = localStorage.getItem(KEYS.USER_ID);

                await api.PUT(URLS.PURCHASE_ITEM, {
                    loginId: `${loginId}`,
                    itemId:  itemRow.ItemId,
                });
                const user = await api.POST<UserDTO>(URLS.USER_INFO, loginId);

                setMyCash(user!.Cash);
                setPurchaseItem(itemRow.ItemName);
                setShowPurchaseDialog(true);
            }
            update(itemRow);
        }, []);

    return (
        <CommonBorderTd>
        {
            item.IsPurchased ? <CommonButton text="購入済"
                                             onClick={() => {}}
                                             disabled={true}
                                             styleObj={{width: "80px"}}/>
                             : <CommonButton text={myCash! < item.ItemPrice ? "資金不足" : "購入"}
                                             onClick={() => purchase(item)}
                                             styleObj={{width: "80px"}}
                                             disabled={myCash! < item.ItemPrice}/>
        }
            <input type="hidden" value={item.ItemId}/>
        </CommonBorderTd>
    );
}

export default PurchaseCommonButton;