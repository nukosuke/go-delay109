package delay109

import (
	"time"
)

// LineStatus 路線ステータス
type LineStatus struct {
	lineID        LineID
	direction     int
	date          time.Time
	delayDuration time.Duration
}

// 読み取り専用構造体のためコンストラクタ非公開
func createLineStatus(
	lineID LineID,
	direction int,
	date time.Time,
	delayDuration time.Duration) *LineStatus {
	return &LineStatus{
		lineID:        lineID,
		direction:     direction,
		date:          date,
		delayDuration: delayDuration,
	}
}

// LineID 路線の識別番号を返す
func (ls *LineStatus) LineID() LineID {
	return ls.lineID
}

// LineDirection 路線の方向を返す(上り/下り)
func (ls *LineStatus) LineDirection() int {
	return ls.direction
}

// LineNameJa 路線の日本語名を返す
func (ls *LineStatus) LineNameJa() string {
	return lineNameJa[ls.lineID]
}

// LineNameEn 路線の英語名を返す
func (ls *LineStatus) LineNameEn() string {
	return lineNameEn[ls.lineID]
}

// DirectionText 方向を返す
func (ls *LineStatus) DirectionText() string {
	return directionText[ls.direction]
}

// DelayDuration 遅延時間を返す
func (ls *LineStatus) DelayDuration() time.Duration {
	return ls.delayDuration
}
