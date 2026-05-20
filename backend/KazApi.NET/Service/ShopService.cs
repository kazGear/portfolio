using KazApi.Domain.DTO;
using KazApi.Repository;
using KazApi.Repository.sql;

namespace KazApi.Service
{
    public class ShopService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public ShopService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(configuration);
        }

        public ItemDTO SelectItemOne(string itemId)
        {
            var param = new { item_id = itemId };
            return _posgre.Select<ItemDTO>(ShopSQL.SelectItemOne(), param)
                          .Single();
        }
        /// <summary>
        /// 店舗リスト取得
        /// </summary>
        public IEnumerable<ShopDTO> SelectShops(string loginId)
        {
            var param = new { login_id = loginId };
            return _posgre.Select<ShopDTO>(ShopSQL.SelectShops(), param);
        }

        /// <summary>
        /// 店舗リスト取得
        /// </summary>
        public IEnumerable<ItemDTO> SelectShopItems(string loginId, string shopId)
        {
            var param = new
            {
                login_id = loginId,
                shop_id = shopId
            };
            return _posgre.Select<ItemDTO>(ShopSQL.SelectItems(), param);
        }

        /// <summary>
        /// 開放可能な店舗一覧を取得
        /// </summary>
        public IEnumerable<ShopDTO> ExistsUsableShop(string loginId)
        {
            var param = new { login_id = loginId };
            return _posgre.Select<ShopDTO>(ShopSQL.ExistsUsableShop(), param);
        }

        /// <summary>
        /// 開放可能な店舗一覧を登録
        /// </summary>
        public void InsertUsableShop(string loginId, IEnumerable<ShopDTO> shops)
        {
            foreach (ShopDTO shop in shops)
            {
                var param = new
                {
                    login_id = loginId,
                    shop_id = shop.ShopId
                };
                _posgre.Execute(ShopSQL.InsertUsableStore(), param);
            }
        }

        public void Purchase(string loginId, string itemId)
        {
            var param = new
            {
                login_id = loginId,
                item_id = itemId
            };
            _posgre.Execute(ShopSQL.Purchase(), param);
        }

    }
}
