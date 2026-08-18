namespace PrivateApi.Domain.DTO
{
    public record JobsResponse
    {
        public int TotalCount { get; init; }

        public int Page { get; init; }

        public int PageSize { get; init; }

        public int TotalPages { get; init; }

        public bool HasPrev { get; init; }

        public bool HasNext { get; init; }

        public IEnumerable<JobDTO> Jobs { get; init; } = [];
    }
}
