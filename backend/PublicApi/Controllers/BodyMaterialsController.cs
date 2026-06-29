using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;
using PublicApi.ResponseDtos;
using PublicApi.Services;

namespace PublicApi.Controllers
{
    [ApiController]
    public class BodyMaterialsController : ControllerBase
    {
        private readonly ILogger<BodyMaterialsController> _logger;
        private readonly BodyMaterialsService _service;

        public BodyMaterialsController(IConfiguration configuration, ILogger<BodyMaterialsController> logger)
        {
            _logger = logger;
            _service = new BodyMaterialsService(configuration);
        }

        [HttpGet("/api/v1/bodyMaterials")]
        public IActionResult Get()
        {
            try
            {
                IEnumerable<CodeResponse> bodyMaterials = _service.Get();
                return StatusCode(HttpStatus.OK, bodyMaterials);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }
    }
}
