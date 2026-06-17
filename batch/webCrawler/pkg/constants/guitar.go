package constants

// guitar makers
const (
    Esp int = iota + 1
    Fender
    Gibson
    Strandberg
    Schecter
    EspSignature
    Ibanez
    PRS
    Suhr
    MusicMan
    Zemaitis
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

type wood struct {
	Name string
	Code int
}

// wood materials
func GetWoods() []wood {
	woods := []wood{
		{"Unknown", 0},
		{"HardMaple", 1},
		{"FlameMaple", 2},
		{"FlamedMaple", 2},
		{"QuiltedMaple", 3},
		{"BirdseyeMaple", 4},
		{"RoastedMaple", 5},
		{"Maple", 6},
		{"HonduranMahogany", 7},
		{"Mahogany", 8},
		{"Sapele", 9},
		{"Korina", 10},
		{"WhiteKorina", 11},
		{"Alder", 12},
		{"Ash", 13},
		{"Basswood", 14},
		{"Linden", 14},
		{"Poplar", 15},
		{"Spruce", 16},
		{"Cedar", 17},
		{"IndianRosewood", 18},
		{"BrazilianRosewood", 19},
		{"Rosewood", 20},
		{"PauFerro", 21},
		{"Ovangkol", 22},
		{"Ebony", 23},
		{"Walnut", 24},
		{"Padauk", 25},
		{"Koa", 26},
		{"Nato", 27},
		{"Agathis", 28},
		{"Bubinga", 29},
		{"Wenge", 30},
		{"Purpleheart", 31},
		{"Zebrawood", 32},
		{"Okoume", 33},
		{"Meranti", 34},
	}
	return woods
}

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
    Price            = "Price"
    ScaleLengthMM    = "ScaleLengthMM"
    Series           = "Series"
    Src              = "Src"
    Weight           = "Weight"
)

// others
const (
	InvalidNumber int = -1
	DecoLabel = "◆◇◆◇ %v ◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇\n"
)