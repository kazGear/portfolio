using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;
using PublicApi.ResponseDtos;
using PublicApi.Services;

namespace PublicApi.Controllers
{
    [ApiController]
    public class MakersController : ControllerBase
    {
        private readonly ILogger<MakersController> _logger;
        private readonly MakersService _service;

        public MakersController(IConfiguration configuration, ILogger<MakersController> logger)
        {
            _logger = logger;
            _service = new MakersService(configuration);
        }

        [HttpGet("/public/v1/makers")]
        public async Task<IActionResult> Get()
        {
            IEnumerable<CodeResponse> makers = await _service.Get();
            return StatusCode(HttpStatus.OK, makers);
        }
    }
}
