import { ItemDTO } from "../../types/Shop";
import BorderTd from "../common/BorderTd";
import React, { useCallback } from "react";
import Button from "../common/Button";
import { KEYS, URLS } from "../../lib/Constants";
import { UserDTO } from "../../types/UserManage";
import { api } from "../../lib/apiClient";

interface ArgProps {
    item: ItemDTO;
    myCash: number | null;
    setMyCash: React.Dispatch<React.SetStateAction<number | null>>;
    setPurchaseItem: React.Dispatch<React.SetStateAction<string>>;
    setShowPurchaseDialog: React.Dispatch<React.SetStateAction<boolean>>;
}

const PurchaseButton = ({
    item,
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
        }, [item]);

    return (
        <BorderTd>
        {
            item.IsPurchased ? <Button text="購入済"
                                       onClick={() => {}}
                                       disabled={true}
                                       styleObj={{width: "80px"}}/>
                             : <Button text={myCash! < item.ItemPrice ? "資金不足"
                                                                      : "購入"
                                       }
                                       onClick={() => purchase(item)}
                                       styleObj={{width: "80px"}}
                                       disabled={myCash! < item.ItemPrice}/>
        }
            <input type="hidden" value={item.ItemId}/>
        </BorderTd>
    );
}

export default PurchaseButton;