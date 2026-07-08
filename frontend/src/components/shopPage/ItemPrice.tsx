import { ItemDTO } from "../../types/Shop";
import CommonAccent from "../common/CommonAccent";
import BorderTd from "../common/CommonBorderTd";

interface ArgProps {
    item: ItemDTO;
    myCash: number | null;
}

const ItemPrice = ({item, myCash}: ArgProps) => {
    return (
        <BorderTd>
        {
            myCash! < item.ItemPrice ? <CommonAccent>{item.ItemPrice}</CommonAccent>
                                     : item.ItemPrice
        } Gil
        </BorderTd>
    );
}

export default ItemPrice;