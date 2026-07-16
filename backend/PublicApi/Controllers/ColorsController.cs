using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;
using PublicApi.ResponseDtos;
using PublicApi.Services;

namespace PublicApi.Controllers
{
    [ApiController]
    public class ColorsController : ControllerBase
    {
        private readonly ILogger<ColorsController> _logger;
        private readonly ColorsService _service;

        public ColorsController(IConfiguration configuration, ILogger<ColorsController> logger)
        {
            _logger = logger;
            _service = new ColorsService(configuration);
        }

        [HttpGet("/public/v1/Colors")]
        public async Task<IActionResult> Get()
        {
            IEnumerable<CodeResponse> colors = await _service.Get();
            return StatusCode(HttpStatus.OK, colors);
        }
    }
}
