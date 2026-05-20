using KazApi.Common._Filter;

namespace KazApi.Domain.DTO
{
    /// <summary>
    /// コードパラメータ for DB
    /// </summary>
    public class ItemDTO
    {
        private string _itemId;
        private string _itemName;
        private int _itemPrice;

        public string ItemId
        {
            get {  return _itemId; }
            set { _itemId = Validation.Id(value); }
        }

        public string ItemName
        {
            get { return _itemName; }
            set { _itemName = Validation.Name(value); }
        }

        public int ItemPrice
        {
            get { return _itemPrice; }
            set { _itemPrice = Validation.Amount(value); }
        }

        public string Remarks { get; set; }
        public string ItemImage { get; set; }
        public bool IsPurchased { get; set; }
    }
}
