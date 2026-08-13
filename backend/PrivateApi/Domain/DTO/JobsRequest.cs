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

        [Range(0, 10000000)]
        [FromQuery(Name = "min_salary_at_month_specified_min")]
        public string? MinSalaryAtMonthSpecifiedMin { get; init; }

        [Range(0, 10000000)]
        [FromQuery(Name = "min_salary_at_month_specified_max")]
        public string? MinSalaryAtMonthSpecifiedMax { get; init; }

        [Range(0, 10000000)]
        [FromQuery(Name = "max_salary_at_month_specified_min")]
        public string? MaxSalaryAtMonthSpecifiedMin { get; init; }

        [Range(0, 10000000)]
        [FromQuery(Name = "max_salary_at_month_specified_max")]
        public string? MaxSalaryAtMonthSpecifiedMax { get; init; }

        [MaxLength(10)]
        [FromQuery(Name = "work_place")]
        public string? WorkPlace { get; init; }

        [MaxLength(20)]
        [FromQuery(Name = "source_site")]
        public string? SourceSite { get; init; }

        [FromQuery(Name = "feature_name")]
        public IEnumerable<string>? FeatureName { get; init; }

        [FromQuery(Name = "options")]
        public IEnumerable<string>? Options { get; init; }

        [Range(1, 100)]
        [FromQuery(Name = "page")]
        public int Page { get; init; } = 1;

        [Range(10, 100)]
        [FromQuery(Name = "pageSize")]
        public int PageSize { get; init; } = 50; // 大量取得防止
    }
}