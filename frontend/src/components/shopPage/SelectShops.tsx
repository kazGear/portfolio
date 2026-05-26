import { useEffect, useState } from "react";
import OutSideFrame from "../common/OutSideFrame";
import Select from "../common/Select";
import { ShopDTO } from "../../types/Shop";
import { KEYS, URLS } from "../../lib/Constants";
import { UserDTO } from "../../types/UserManage";
import Accent from "../common/Accent";
import { api } from "../../lib/apiClient";

interface ArgProps {
    setSelectedShop: React.Dispatch<React.SetStateAction<string | undefined>>;
    user: UserDTO | null;
    myCash: number | null;
}

const SelectShops = ({setSelectedShop, user, myCash}: ArgProps) => {
    const [shopsOfSelectBox, setShopsOfSelectBox] = useState<ShopDTO[]>([])

    /**
     * 店舗の選択肢を取得
     */
    useEffect(() => {
        const selectShops = async () => {
            const loginId = localStorage.getItem(KEYS.USER_ID);
            const shops = await api.POST<ShopDTO[]>(URLS.SHOP_INIT, loginId);
            setShopsOfSelectBox(shops!);
        }
        selectShops();
    }, []);
    /**
     * ショップ選択時
     */
    const changeShopHandler = (e: React.ChangeEvent<HTMLSelectElement> | undefined) => {
        setSelectedShop(e?.target.value);
    }

    return (
        <OutSideFrame >
            <h3 style={{margin: "10px"}}>
                所持金：<Accent>{myCash?.toLocaleString()}</Accent> Gil
            </h3>
            <Select title="店舗" onChange={changeShopHandler}>
                {
                    shopsOfSelectBox.map((shop, index) => (
                        <option value={shop.ShopId} key={index}>{shop.ShopName}</option>
                    ))
                }
            </Select>
        </OutSideFrame>
    );
}

export default SelectShops;