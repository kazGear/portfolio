import { ItemDTO } from "../../types/Shop";
import BorderTd from "../common/BorderTd";
import React from "react";

interface ArgProps {
    item: ItemDTO;
}

const ItemRemarks = ({item}: ArgProps) => {

    return (
        <BorderTd>
        {
            item.Remarks.split("\n").map((line, index) => (
                <React.Fragment key={index}>
                    {line}<br/>
                </React.Fragment>
            ))
        }
        </BorderTd>
    );
}

export default ItemRemarks;