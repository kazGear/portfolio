namespace PublicApi.ResponseDtos
{
    public record GuitarsResponse
    {
        public int TotalCount { get; init; }

        public int Page { get; init; }

        public int PageSize { get; init; }

        public int TotalPages { get; init; }

        public bool HasPrev { get; init; }

        public bool HasNext { get; init; }

        public IEnumerable<GuitarResponse> Guitars { get; init; } = [];
    }
}
