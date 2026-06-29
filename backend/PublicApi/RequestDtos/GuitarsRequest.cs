namespace PublicApi.RequestDtos
{
    public record GuitarsRequest
    {
        public string? MakerCd { get; init; }
 
        public string? Name { get; init; }

        public string? Series { get; init; }

        public int? ColorCd { get; init; }

        public int? BodyMaterialTopCd { get; init; }

        public int? BodyMaterialBackCd { get; init; }

        public int? MinPrice { get; init; }

        public int? MaxPrice { get; init; }

        public string Order { get; init; } = "ASC";

        public string? Sort { get; init; }

        public int Page { get; init; } = 1; // 大量取得防止

        public int PageSize { get; init; } = 50; // 大量取得防止
    }
}
