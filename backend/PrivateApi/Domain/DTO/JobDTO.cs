using Microsoft.EntityFrameworkCore.Storage.ValueConversion.Internal;

namespace PrivateApi.Domain.DTO
{
    public record JobDTO
    {
        public long Id { get; init; }

        public string Url { get; init; } = string.Empty;

        public string Title { get; init; } = string.Empty;

        public string Location { get; init; } = string.Empty;

        public int MinSalaryAtMonth { get; init; }

        public int MaxSalaryAtMonth { get; init; }

        public string EmploymentType { get; init; } = string.Empty;

        public string WorkPlace { get; init; } = string.Empty;

        public string SourceSite { get; init; } = string.Empty;

        public DateTime CreatedAt { get; init; }

        public DateTime UpdatedAt { get; init; }

        public string FeatureNamesCSV {  get; init; } = string.Empty;

        public string OptionsCSV { get; init; } = string.Empty;

        public IEnumerable<string> FeatureNames { get; set; } = [];

        public IEnumerable<string> Options { get; set; } = [];
    }

    public record ProjectUsageByFeatureDTO
    {
        public string FeatureName { get; init; } = string.Empty;

        public int FeatureCount { get; init; }

        public double Ratio { get; init; }
    }

    public record WorkPlaceByPrefectureDTO
    {
        public string Location { get; set; } = string.Empty;

        public int FullRemote { get; init; }

        public int Hybrid { get; init; }

        public int OnSite { get; init; }
    }

    public record SalaryRangeByFeatureDTO
    {
        public string FeatureName { get; init; } = string.Empty;

        public double SalaryLower { get; set; }

        public double SalaryMedian { get; set; }

        public double SalaryHigher { get; set; }
    }
}
