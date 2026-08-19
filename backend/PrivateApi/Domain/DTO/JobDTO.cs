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

    public record ProjectUsageByLanguageDTO
    {
        public string FeatureName { get; init; } = string.Empty;

        public int JobCount { get; init; }

        public double Ratio { get; init; }
    }
}
