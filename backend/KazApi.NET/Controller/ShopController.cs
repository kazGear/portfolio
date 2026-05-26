using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using KazApi.Repository;
using KazApi.Domain.DTO;
using KazApi.Service;

namespace KazApi.Controller
{
    [ApiController]
    public class ShopController : ControllerBase
    {
        private readonly ShopService _service;
        private readonly UserService _userService;
        private readonly IDatabase _posgre;

        public ShopController(IConfiguration configuration)
        {
            _service = new ShopService(configuration);
            _userService = new UserService(configuration);
            _posgre = new PostgreSQL(configuration);
        }

        [HttpPost("api/shop/itemInfo")]
        public ActionResult<string> SelectItemInfo([FromForm] string itemId)
        {
            ItemDTO item = _service.SelectItemOne(itemId);
            return JsonConvert.SerializeObject(item);
        }

        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/shop/init")]
        public ActionResult<string> Init([FromBody] string loginId)
        {
            // ショップリストを取得
            IEnumerable<ShopDTO> shops = _service.SelectShops(loginId);
            return JsonConvert.SerializeObject(shops);
        }

        //
        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/shop/items")]
        public ActionResult<string> SelectShopItem([FromForm] string loginId,
                                                   [FromForm] string? selectedShop)
        {
            if (selectedShop == null) return JsonConvert.SerializeObject(new List<string>());

            // アイテムリストを取得
            IEnumerable<ItemDTO> shops = _service.SelectShopItems(loginId, selectedShop);
            return JsonConvert.SerializeObject(shops);
        }

        [HttpPut("api/shop/purchase")]
        public ActionResult InsertMyItem([FromForm] string loginId,
                                         [FromForm] string itemId)
        {
            // クレンジング
            loginId = loginId.Trim();
            itemId = itemId.Trim();

            // 残金取得
            int cash = _userService.SelectUserOne(loginId).Cash;
            // 購入品取得
            ItemDTO item = _service.SelectItemOne(itemId);

            if (cash < item.ItemPrice) throw new Exception("資金が不足しています。");

            // 各種登録
            _service.Purchase(loginId, itemId);
            _userService.Purchase(loginId, (cash - item.ItemPrice));
                        
            return Ok(200);
        }
    }
}
