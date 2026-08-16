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

        [HttpGet("api/jobs/get")]
        public async Task<IActionResult> GetJobs([FromQuery] JobsRequest req)
        {
            JobsResponse jobs = await _service.GetJobs(req);
            return StatusCode(HttpStatus.OK, jobs);
        }

        [HttpGet("api/languages/get")]
        public async Task<IActionResult> GetLanguages([FromQuery] JobsRequest req)
        {
            IEnumerable<string> features = await _service.GetLanguages(req);
            return StatusCode(HttpStatus.OK, features);
        }
    }
}
