using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using KazApi.Repository;
using KazApi.Domain.DTO;
using KazApi.Service;
using CSLib.Lib;
using KazApi.Common;
using System.Transactions;

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
        public IActionResult SelectItemInfo([FromBody] string itemId)
        {
            if (string.IsNullOrEmpty(itemId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                ItemDTO item = _service.SelectItemOne(itemId);
                return StatusCode(HttpStatus.OK, item);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/shop/init")]
        public IActionResult Init([FromBody] string? loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                // ショップリストを取得
                IEnumerable<ShopDTO> shops = _service.SelectShops(loginId);
                return StatusCode(HttpStatus.OK, shops);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        //
        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/shop/items")]
        public IActionResult SelectShopItem([FromBody] ReqShopItem req)
        {
            if (req.selectedShop == null)
                return StatusCode(HttpStatus.OK, new List<string>());

            try
            {
                // アイテムリストを取得
                IEnumerable<ItemDTO> shops = _service.SelectShopItems(req.loginId, req.selectedShop);
                return StatusCode(HttpStatus.OK, shops);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        [HttpPut("api/shop/purchase")]
        public IActionResult InsertMyItem([FromBody] ReqMyItem req)
        {
            using (TransactionScope transaction = new TransactionScope())
            {
                try
                {
                    // クレンジング
                    string loginId = req.loginId.Trim();
                    string itemId  = req.itemId.Trim();

                    // 残金取得
                    int cash = _userService.SelectUserOne(loginId).Cash;
                    // 購入品取得
                    ItemDTO item = _service.SelectItemOne(itemId);

                    if (cash < item.ItemPrice)
                    {
                        return StatusCode(HttpStatus.InternalServerError, Message.Create("資金が不足しています。"));
                    }

                    // 各種登録
                    _service.Purchase(loginId, itemId);
                    _userService.Purchase(loginId, (cash - item.ItemPrice));

                    transaction.Complete();
                    return StatusCode(HttpStatus.OK);
                }
                catch (Exception e)
                {
                    return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
                }
            }
        }
    }
}
