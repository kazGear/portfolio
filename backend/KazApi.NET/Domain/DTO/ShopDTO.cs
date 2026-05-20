using KazApi.Common._Filter;

namespace KazApi.Domain.DTO
{
    /// <summary>
    /// コードパラメータ for DB
    /// </summary>
    public class ShopDTO
    {
        private string _shopId;
        private string _shopName;

        public string ShopId
        { 
            get { return _shopId; }
            set { _shopId = Validation.Id(value); }
        }

        public string ShopName
        {
            get { return _shopName; }
            set { _shopName = Validation.Name(value); }
        }

        public int WinMoneyUntilCanUse { get; set; }
    }
}
