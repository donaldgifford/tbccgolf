package services

import (
	"fmt"

	"gorm.io/gorm"
)

type (
	HoleScore struct {
		gorm.Model
		Par      uint
		Handicap int
		Number   int
		ScoreID  uint
	}
)

var scoreCard = map[int]HoleScore{
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

func generateHoles(start int, end int) []HoleScore {
	holeCount := end - start + 1
	fmt.Printf("Hole Count: %d\n", holeCount)

	var hs []HoleScore

	for i := 1; i <= holeCount; i++ {
		fmt.Println(i)
		hs = append(hs, scoreCard[i])
	}

	return hs
}
