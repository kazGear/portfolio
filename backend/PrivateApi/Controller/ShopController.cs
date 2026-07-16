using Microsoft.AspNetCore.Mvc;
using PrivateApi.Domain.DTO;
using PrivateApi.Service;
using CSLib.Lib;
using PrivateApi.Common;
using System.Transactions;

namespace PrivateApi.Controller
{
    [ApiController]
    public class ShopController : ControllerBase
    {
        private readonly ShopService _service;
        private readonly UserService _userService;

        public ShopController(IConfiguration configuration)
        {
            _service = new ShopService(configuration);
            _userService = new UserService(configuration);
        }

        [HttpPost("api/shop/itemInfo")]
        public async Task<IActionResult> SelectItemInfo([FromBody] string itemId)
        {
            if (string.IsNullOrEmpty(itemId)) return StatusCode(HttpStatus.BadRequest);

            ItemDTO item = await _service.SelectItemOne(itemId);
            return StatusCode(HttpStatus.OK, item);
        }

        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/shop/init")]
        public async Task<IActionResult> Init([FromBody] string? loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);
  
            // ショップリストを取得
            IEnumerable<ShopDTO> shops = await _service.SelectShops(loginId);
            return StatusCode(HttpStatus.OK, shops);
        }

        //
        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/shop/items")]
        public async Task<IActionResult> SelectShopItem([FromBody] ReqShopItem req)
        {
            if (req.selectedShop == null)
                return StatusCode(HttpStatus.OK, new List<string>());

            // アイテムリストを取得
            IEnumerable<ItemDTO> shops = await _service.SelectShopItems(req.loginId, req.selectedShop);
            return StatusCode(HttpStatus.OK, shops);
        }

        [HttpPut("api/shop/purchase")]
        public async Task<IActionResult> InsertMyItem([FromBody] ReqMyItem req)
        {
            using (var transaction = new TransactionScope(TransactionScopeAsyncFlowOption.Enabled))
            {
                // クレンジング
                string loginId = req.loginId.Trim();
                string itemId  = req.itemId.Trim();

                // 残金取得
                UserDTO? user = await _userService.SelectUserOne(loginId);
                int cash      = user.Cash;

                // 購入品取得
                ItemDTO item = await _service.SelectItemOne(itemId);

                if (cash < item.ItemPrice)
                {
                    return StatusCode(HttpStatus.InternalServerError, Message.Create("資金が不足しています。"));
                }

                // 各種登録
                await _service.Purchase(loginId, itemId);
                await _userService.Purchase(loginId, (cash - item.ItemPrice));

                transaction.Complete();
                return StatusCode(HttpStatus.OK);
            }
        }
    }
}
