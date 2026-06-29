using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;
using PublicApi.RequestDtos;
using PublicApi.ResponseDtos;
using PublicApi.Services;

namespace PublicApi.Controllers
{
    [ApiController]
    public class SeriesController : ControllerBase
    {
        private readonly ILogger<SeriesController> _logger;
        private readonly SeriesService _service;

        public SeriesController(IConfiguration configuration, ILogger<SeriesController> logger)
        {
            _logger = logger;
            _service = new SeriesService(configuration);
        }

        [HttpGet("/api/v1/series")]
        public IActionResult Get([FromQuery] SeriesRequest req)
        {
            try
            {
                IEnumerable<CodeResponse> colors = _service.Get(req);
                return StatusCode(HttpStatus.OK, colors);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }
    }
}
