import { ItemDTO } from "../../types/Shop";
import Accent from "../common/Accent";
import BorderTd from "../common/BorderTd";

interface ArgProps {
    item: ItemDTO;
    myCash: number | null;
}

const ItemPrice = ({item, myCash}: ArgProps) => {
    return (
        <BorderTd>
        {
            myCash! < item.ItemPrice ? <Accent>{item.ItemPrice}</Accent>
                                     : item.ItemPrice
        } Gil
        </BorderTd>
    );
}

export default ItemPrice;