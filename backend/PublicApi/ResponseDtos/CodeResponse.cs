namespace PublicApi.ResponseDtos
{
    public record CodeResponse
    {
        public string Category { get; set; } = string.Empty;

        public int Code { get; init; }

        public string Name { get; init; } = string.Empty;
    }
}
