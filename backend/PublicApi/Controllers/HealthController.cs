using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;

namespace PublicApi.Controllers
{
    [ApiController]
    public class HealthController : ControllerBase
    {
        private readonly ILogger<HealthController> _logger;

        public HealthController(ILogger<HealthController> logger)
        {
            _logger = logger;
        }

        [HttpGet("/health")]
        public IActionResult Health()
        {
            return StatusCode(HttpStatus.OK, "This API is healthy !!");
        }
    }
}
