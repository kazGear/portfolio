import { useEffect, useState } from "react";
import CommonFrame from "../common/CommonFrame";
import CommonSelect from "../common/CommonSelect";
import { ShopDTO } from "../../types/Shop";
import { KEYS, URLS } from "../../lib/Constants";
import { UserDTO } from "../../types/UserManage";
import CommonAccent from "../common/CommonAccent";
import { api } from "../../lib/apiClient";
import useApiErrorHandler from "../../hooks/useApiErrorHandler";
import { isEmpty } from "../../lib/CommonLogic";
import { ApiError } from "../../types/ApiError";

interface ArgProps {
    setSelectedShop: React.Dispatch<React.SetStateAction<string | undefined>>;
    user: UserDTO | null;
    myCash: number | null;
}

const SelectShops = ({setSelectedShop, user, myCash}: ArgProps) => {
    const [shopsOfSelectBox, setShopsOfSelectBox] = useState<ShopDTO[]>([]);
    const errorHandler                            = useApiErrorHandler();

    /**
     * 店舗の選択肢を取得
     */
    useEffect(() => {
        const selectShops = async () => {
            const loginId = localStorage.getItem(KEYS.USER_ID);

            try {
                const shops = await api.POST<ShopDTO[]>(URLS.SHOP_INIT, loginId);
                setShopsOfSelectBox(shops!);

                if (isEmpty(shops)) throw new ApiError(500, "Get shops failed ...")

            } catch (e) {
                console.log(e);
                errorHandler(e);
            }
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
        <CommonFrame>
            <h3 style={{margin: "10px"}}>
                所持金：<CommonAccent>{myCash?.toLocaleString()}</CommonAccent> Gil
            </h3>
            <CommonSelect title="店舗" onChange={changeShopHandler}>
                {
                    shopsOfSelectBox.map((shop, index) => (
                        <option value={shop.ShopId} key={shop.ShopName + index}>
                            {shop.ShopName}
                        </option>
                    ))
                }
            </CommonSelect>
        </CommonFrame>
    );
}

export default SelectShops;