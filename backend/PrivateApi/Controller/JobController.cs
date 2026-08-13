using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;
using PrivateApi.Domain.DTO;
using PrivateApi.Service;

namespace PublicApi.Controllers
{
    [ApiController]
    public class JobController : ControllerBase
    {
        private readonly ILogger<JobController> _logger;
        private readonly JobService _service;

        public JobController(IConfiguration configuration, ILogger<JobController> logger)
        {
            _logger = logger;
            _service = new JobService(configuration);
        }

        [HttpGet("api/job/get")]
        public async Task<IActionResult> Get([FromQuery] JobsRequest req)
        {
            JobsResponse jobs = await _service.Get(req);
            return StatusCode(HttpStatus.OK, jobs);
        }
    }
}
