package constants

// ギターメーカー
const (
	ESP int = iota + 1
	Fender
	Gibson
	Strandberg
	SCHECTER
	ESPsignature
)

const (
	InvalidNumber int = -1
	DecoLabel = "◆◇◆◇ %v ◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇◆◇\n"
)

type wood struct {
	Name string
	Code int
}

// ギターで使われる木材セット
func GetWoods() []wood {
	woods := []wood{
		{"Unknown", 0},
		{"HardMaple", 1},
		{"FlameMaple", 2},
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
	}
	return woods
}
