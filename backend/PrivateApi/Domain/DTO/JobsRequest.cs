using Microsoft.AspNetCore.Mvc;
using System.ComponentModel.DataAnnotations;

namespace PrivateApi.Domain.DTO
{
    public record JobsRequest
    {
        [MaxLength(100)]
        [FromQuery(Name = "title")]
        public string? Title { get; init; }

        [MaxLength(5)]
        [FromQuery(Name = "location")]
        public string? Location { get; init; }

        [Range(-3, 10000000)]
        [FromQuery(Name = "minSalaryAtMonthSpecifiedMin")]
        public string? MinSalaryAtMonthSpecifiedMin { get; init; }

        [Range(-3, 10000000)]
        [FromQuery(Name = "minSalaryAtMonthSpecifiedMax")]
        public string? MinSalaryAtMonthSpecifiedMax { get; init; }

        [Range(-3, 10000000)]
        [FromQuery(Name = "maxSalaryAtMonthSpecifiedMin")]
        public string? MaxSalaryAtMonthSpecifiedMin { get; init; }

        [Range(-3, 10000000)]
        [FromQuery(Name = "maxSalaryAtMonthSpecifiedMax")]
        public string? MaxSalaryAtMonthSpecifiedMax { get; init; }

        [MaxLength(10)]
        [FromQuery(Name = "workPlace")]
        public string? WorkPlace { get; init; }

        [MaxLength(20)]
        [FromQuery(Name = "sourceSite")]
        public string? SourceSite { get; init; }

        [FromQuery(Name = "featureNames")]
        public IEnumerable<string>? FeatureNames { get; set; }

        [FromQuery(Name = "options")]
        public IEnumerable<string>? Options { get; set; }

        [Range(1, 100)]
        [FromQuery(Name = "page")]
        public int Page { get; init; } = 1;

        [Range(10, 100)]
        [FromQuery(Name = "pageSize")]
        public int PageSize { get; init; } = 50; // 大量取得防止
    }
}