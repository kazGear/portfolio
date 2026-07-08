import { ItemDTO } from "../../types/Shop";
import CommonBorderTd from "../common/CommonBorderTd";
import React from "react";

interface ArgProps {
    item: ItemDTO;
}

const ItemRemarks = ({item}: ArgProps) => {

    return (
        <CommonBorderTd>
        {
            item.Remarks.split("\n").map((line, index) => (
                <React.Fragment key={line + index}>
                    {line}<br/>
                </React.Fragment>
            ))
        }
        </CommonBorderTd>
    );
}

export default ItemRemarks;