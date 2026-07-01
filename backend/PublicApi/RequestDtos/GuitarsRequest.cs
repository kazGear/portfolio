using Microsoft.AspNetCore.Mvc;
using System.ComponentModel.DataAnnotations;

namespace PublicApi.RequestDtos
{
    public record GuitarsRequest
    {
        [Range(0, 50)]
        [FromQuery(Name = "makerCd")]
        public int? MakerCd { get; init; }

        [MaxLength(100)]
        [FromQuery(Name = "name")]
        public string? Name { get; init; }

        [MaxLength(100)]
        [FromQuery(Name = "series")]
        public string? Series { get; init; }

        [Range(0, 20)]
        [FromQuery(Name = "colorCd")]
        public int? ColorCd { get; init; }

        [Range(0, 50)]
        [FromQuery(Name = "bodyMaterialTopCd")]
        public int? BodyMaterialTopCd { get; init; }

        [Range(0, 50)]
        [FromQuery(Name = "bodyMaterialBackCd")]
        public int? BodyMaterialBackCd { get; init; }

        [Range(-3, int.MaxValue)] // 1: parse error, 2: open price, 3: undefined
        [FromQuery(Name = "minPrice")]
        public int? MinPrice { get; init; }

        [Range(0, int.MaxValue)]
        [FromQuery(Name = "maxPrice")]
        public int? MaxPrice { get; init; }

        [FromQuery(Name = "order")]
        public string Order { get; init; } = "ASC";

        [FromQuery(Name = "sort")]
        public string? Sort { get; init; }

        [Required]
        [Range(1, 50)]
        [FromQuery(Name = "page")]
        public int Page { get; init; } = 1; // 大量取得防止

        [Required]
        [Range(10, 100)]
        [FromQuery(Name = "pageSize")]
        public int PageSize { get; init; } = 25; // 大量取得防止
    }
}
