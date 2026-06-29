using Microsoft.AspNetCore.Mvc;
using System.ComponentModel.DataAnnotations;

namespace PublicApi.RequestDtos
{
    public record SeriesRequest
    {
        [Required]
        [Range(1, 50)]
        [FromQuery(Name = "makerCd")]
        public int MakerCd { get; init; }
    }
}
