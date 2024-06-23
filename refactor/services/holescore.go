package services

import (
	"fmt"

	"gorm.io/gorm"
)

type HoleScore struct {
	gorm.Model
	ScoreID         uint
	Strokes         uint `gorm:"default:0"`
	HoleNumber      uint
	ScoreCardNumber uint
	Par             uint
	Handicap        uint
}

type ScoreCardScore struct {
	Par      uint
	Handicap uint
	Number   uint
}

var scoreCard = map[int]ScoreCardScore{
	1:  {Par: 4, Handicap: 5, Number: 1},
	2:  {Par: 4, Handicap: 3, Number: 2},
	3:  {Par: 4, Handicap: 9, Number: 3},
	4:  {Par: 3, Handicap: 17, Number: 4},
	5:  {Par: 5, Handicap: 7, Number: 5},
	6:  {Par: 3, Handicap: 11, Number: 6},
	7:  {Par: 4, Handicap: 13, Number: 7},
	8:  {Par: 4, Handicap: 1, Number: 8},
	9:  {Par: 4, Handicap: 15, Number: 9},
	10: {Par: 4, Handicap: 10, Number: 10},
	11: {Par: 4, Handicap: 6, Number: 11},
	12: {Par: 4, Handicap: 8, Number: 12},
	13: {Par: 3, Handicap: 16, Number: 13},
	14: {Par: 4, Handicap: 2, Number: 14},
	15: {Par: 4, Handicap: 14, Number: 15},
	16: {Par: 4, Handicap: 18, Number: 17},
	17: {Par: 4, Handicap: 14, Number: 17},
	18: {Par: 4, Handicap: 12, Number: 18},
}

func generateHoles(start int, end int) []ScoreCardScore {
	holeCount := end - start + 1
	// fmt.Printf("Hole Count: %d\n", holeCount)

	var hs []ScoreCardScore

	for i := 1; i <= holeCount; i++ {
		fmt.Println(i)
		hs = append(hs, scoreCard[i])
	}

	return hs
}

func NewServicesHoleScore(hs HoleScore, db *gorm.DB) *ServicesHoleScore {
	return &ServicesHoleScore{
		HoleScore: hs,
		DB:        db,
	}
}

type ServicesHoleScore struct {
	HoleScore HoleScore
	DB        *gorm.DB
}

func (shs *ServicesHoleScore) GenerateHoleScores(scoreID uint) error {
	scoreCardScores := generateHoles(1, 9)

	var holescores []HoleScore

	for h, scs := range scoreCardScores {
		hscore := HoleScore{
			ScoreID:         scoreID,
			Strokes:         0,
			HoleNumber:      uint(h + 1),
			ScoreCardNumber: uint(scs.Number),
			Par:             uint(scs.Par),
			Handicap:        uint(scs.Handicap),
		}

		holescores = append(holescores, hscore)
	}

	if err := shs.DB.Create(&holescores).Error; err != nil {
		return err
	}

	return nil
}
