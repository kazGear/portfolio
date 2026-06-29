namespace PublicApi.ResponseDtos
{
    public record CodeResponse
    {
        public string Category { get; init; } = string.Empty;
        public int Code { get; init; }
        public string Name { get; init; } = string.Empty;
    }
}
