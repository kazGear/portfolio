package constants

// guitar makers
const (
    Esp int = iota + 1
    Fender
    Gibson
    Strandberg
    SCHECTER
    EspSignature
    Ibanez
    PRS
    Suhr
    MusicMan
    ZEMAITIS
    Momose
)

// format of date
const (
	DateOnly     = "2006-01-02"
	DateTime     = "2006-01-02_T15-04-05"
	DateTimeMin  = "2006-01-02_T15-04"
	FileDateTime = "20060102_150405"
)

// color code
const (
    Red int = iota + 1
    Pink
    Orange
    Yellow
    Green
    SkyBlue
    Blue
    Purple
    Gray
    Black
    White
    Natural
    Brown
    Gold
    Silver
    OthersColor int = 99
)

// Guitar Spec Keys
const (
    Maker            = "Maker"
    Name             = "Name"
    BodyFinish       = "BodyFinish"
    BodyMaterialTop  = "BodyMaterialTop"
    BodyMaterialBack = "BodyMaterialBack"
    Bridge           = "Bridge"
    Color            = "Color"
    Controls         = "Controls"
    Comment          = "Comment"
    Fingerboard      = "Fingerboard"
    FretCount        = "FretCount"
    Inlays           = "Inlays"
    Joint            = "Joint"
    NeckMaterial     = "NeckMaterial"
    Pickups          = "Pickups"
	NeckPickup		 = "NeckPickup"
	CenterPickup	 = "CenterPickup"
	BridgePickup	 = "BridgePickup"
    Price            = "Price"
    ScaleLengthMM    = "ScaleLengthMM"
    Series           = "Series"
    Src              = "Src"
    Weight           = "Weight"
)

// price
const (
	ParseErrorPrice int = -1
	OpenPrice 		int = -2
	UndefinedPrice  int = -3
)

// others
const (
	InvalidNumber int = -1
	DecoLabel = "◆◇◆◇ %v ◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇\n"
)