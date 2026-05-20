import styled from "styled-components";
import { PREFIX } from "../../lib/Constants";
import { ItemDTO } from "../../types/Shop";
import BorderTd from "../common/BorderTd";

const Simg = styled.img`
    width: 50px;
    height: 50px;
    border-radius: 100%;
    vertical-align: middle;
`

interface ArgProps {
    item: ItemDTO;
}

const ItemImage = ({item}: ArgProps) => {
    return (
        <BorderTd>
            <span style={{marginLeft: "10px", display: "inline-block"}}>
                <Simg src={PREFIX.BASE64 + item.ItemImage} alt="書品"/>
            </span>
        </BorderTd>
    );
}

export default ItemImage;