using CSLib.Lib;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;

namespace PrivateApi.Service
{
    public class ShopService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public ShopService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        public async Task<ItemDTO> SelectItemOne(string itemId)
        {
            var param = new { item_id = itemId };
            var items = await _posgre.Select<ItemDTO>(ShopSQL.SelectItemOne(), param);
            return items.Single();                                   
        }
        /// <summary>
        /// 店舗リスト取得
        /// </summary>
        public async Task<IEnumerable<ShopDTO>> SelectShops(string loginId)
        {
            var param = new { login_id = loginId };
            return await _posgre.Select<ShopDTO>(ShopSQL.SelectShops(), param);
        }

        /// <summary>
        /// 店舗リスト取得
        /// </summary>
        public async Task<IEnumerable<ItemDTO>> SelectShopItems(string loginId, string shopId)
        {
            var param = new
            {
                login_id = loginId,
                shop_id = shopId
            };
            return await _posgre.Select<ItemDTO>(ShopSQL.SelectItems(), param);
        }

        /// <summary>
        /// 開放可能な店舗一覧を取得
        /// </summary>
        public async Task<IEnumerable<ShopDTO>> ExistsUsableShop(string loginId)
        {
            var param = new { login_id = loginId };
            return await _posgre.Select<ShopDTO>(ShopSQL.ExistsUsableShop(), param);
        }

        /// <summary>
        /// 開放可能な店舗一覧を登録
        /// </summary>
        public async Task InsertUsableShop(string loginId, IEnumerable<ShopDTO> shops)
        {
            foreach (ShopDTO shop in shops)
            {
                var param = new
                {
                    login_id = loginId,
                    shop_id = shop.ShopId
                };
                await _posgre.Execute(ShopSQL.InsertUsableStore(), param);
            }
        }

        public async Task Purchase(string loginId, string itemId)
        {
            var param = new
            {
                login_id = loginId,
                item_id  = itemId
            };
            await _posgre.Execute(ShopSQL.Purchase(), param);
        }

    }
}
