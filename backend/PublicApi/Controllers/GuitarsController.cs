using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;
using PublicApi.RequestDtos;
using PublicApi.ResponseDtos;
using PublicApi.Services;

namespace PublicApi.Controllers
{
    [ApiController]
    public class GuitarsController : ControllerBase
    {
        private readonly ILogger<GuitarsController> _logger;
        private readonly GuitarsService _service;

        public GuitarsController(IConfiguration configuration, ILogger<GuitarsController> logger)
        {
            _logger = logger;
            _service = new GuitarsService(configuration);
        }

        [HttpGet("/public/v1/guitars")]
        public IActionResult Get([FromQuery] GuitarsRequest req)
        {
            try
            {
                GuitarsResponse guitars = _service.Get(req);
                return StatusCode(HttpStatus.OK, guitars);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }
    }
}
