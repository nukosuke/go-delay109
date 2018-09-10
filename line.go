package delay109

// LineID 路線識別子
type LineID int

const (
	// TY 東横線
	TY LineID = 1 + iota
	// MG 目黒線
	MG
	// DT 田園都市線
	DT
	// OM 大井町線
	OM
	// IK 池上線
	IK
	// TM 東急多摩川線
	TM
	// SG 世田谷線
	SG
	// KD こどもの国線
	KD
)

const (
	// DirectionUp 上り
	DirectionUp = 0
	// DirectionDown 下り
	DirectionDown = 10
)

const (
	// Before10am 始発〜10時
	Before10am = 1 + iota
	// After10am 10時以降
	After10am
)

var lineNameJa = map[LineID]string{
	TY: "東横線",
	MG: "目黒線",
	DT: "田園都市線",
	OM: "大井町線",
	IK: "池上線",
	TM: "東急多摩川線",
	SG: "世田谷線",
	KD: "こどもの国線",
}

var lineNameEn = map[LineID]string{
	TY: "Toyoko Line",
	MG: "Meguro Line",
	DT: "Den-en-toshi Line",
	OM: "Oimachi Line",
	IK: "Ikegami Line",
	TM: "Tokyu Tamagawa Line",
	SG: "Setagaya Line",
	KD: "Kodomonokuni Line",
}

var directionText = map[int]string{
	DirectionUp:   "上り",
	DirectionDown: "下り",
}

// LineNameJa 路線の日本語名を取得する
func LineNameJa(id LineID) string {
	return lineNameJa[id]
}

// LineNameEn 路線の英語名を取得する
func LineNameEn(id LineID) string {
	return lineNameEn[id]
}

// DirectionText 方面のテキスト(上り/下り)を取得する
func DirectionText(direction int) string {
	return directionText[direction]
}
