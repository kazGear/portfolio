using CSLib.Lib;
using PrivateApi.Common;
using PrivateApi.Domain.DTO;
using PrivateApi.Service;
using Microsoft.AspNetCore.Mvc;

namespace PrivateApi.Controller
{
    [ApiController]
    public class CommonController : ControllerBase
    {
        private readonly CommonService _serviceCommon;
        private readonly BattleReportService _serviceReport;

        public CommonController(IConfiguration configuration)
        {
            _serviceCommon = new CommonService(configuration);
            _serviceReport = new BattleReportService(configuration);
        }

        /// <summary>
        /// 画像ファイルアップロード
        /// </summary>
        [HttpPut("api/common/imgUpload")]
        public async Task<IActionResult> UploadImage([FromForm] IFormFile? image,
                                                     [FromForm] string? loginId
        ) {
            if (image == null || image.Length == 0)
                return StatusCode(HttpStatus.BadRequest, Message.Create("No file uploaded."));

            if (string.IsNullOrEmpty(loginId))
                return StatusCode(HttpStatus.BadRequest);

            try
            {
                // 画像のバイナリ化
                using (MemoryStream ms = new MemoryStream())
                {
                    await image.CopyToAsync(ms);
                    byte[] imageByte = ms.ToArray();
                    string imageBASE64 = Convert.ToBase64String(imageByte);

                    _serviceCommon.UpdateImage(loginId, imageBASE64);
                }
                return StatusCode(HttpStatus.OK, Message.Create("Image upload success."));
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e, "Image upload failed."));
            }
        }

        [HttpGet("api/common/FetchElementCode")]
        public IActionResult FetchElementCode()
        {
            try
            {
                IEnumerable<CodeDTO> result = _serviceCommon.FetchElementCode();
                return StatusCode(HttpStatus.OK, result);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }
    }
}
