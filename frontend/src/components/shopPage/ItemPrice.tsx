import { ItemDTO } from "../../types/Shop";
import CommonAccent from "../common/CommonAccent";
import CommonBorderTd from "../common/CommonBorderTd";

interface ArgProps {
    item: ItemDTO;
    myCash: number | null;
}

const ItemPrice = ({item, myCash}: ArgProps) => {
    return (
        <CommonBorderTd>
        {
            myCash! < item.ItemPrice ? <CommonAccent>{item.ItemPrice}</CommonAccent>
                                     : item.ItemPrice
        } Gil
        </CommonBorderTd>
    );
}

export default ItemPrice;