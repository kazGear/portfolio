namespace PublicApi.ResponseDtos
{
    public record GuitarResponse
    {
        public int Maker { get; init; }
        public string Name { get; init; } = string.Empty;
        public string Color { get; init; } = string.Empty;
        public int ColorCd { get; init; }
        public string BodyFinish { get; init; } = string.Empty;
        public string BodyMaterial { get; init; } = string.Empty;
        public int BodyMaterialTop { get; init; }
        public int BodyMaterialBack { get; init; }
        public string Bridge { get; init; } = string.Empty;
        public string Controls { get; init; } = string.Empty;
        public string Comment { get; init; } = string.Empty;
        public int Fingerboard { get; init; }
        public int FretCount { get; init; }
        public string Inlays { get; init; } = string.Empty;
        public string Joint { get; init; } = string.Empty;
        public int NeckMaterial { get; init; }
        public string Pickups { get; init; } = string.Empty;
        public int Price { get; init; }
        public int ScaleLengthMm { get; init; }
        public string Series { get; init; } = string.Empty;
        public string Src { get; init; } = string.Empty;
        public double Weight { get; init; }
    }
}
