export interface ShopDTO {
    ShopId: string;
    ShopName: string;
    WinMoneyUntilCanUse: number;
};

export interface ItemDTO {
    ItemId: string;
    ItemName: string;
    ItemPrice: number;
    Remarks: string;
    ItemImage: string;
    IsPurchased: boolean;
};