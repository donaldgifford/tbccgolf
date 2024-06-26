package services

import (
	"fmt"

	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	PlayerID  uint
	MatchID   uint
	Completed bool
	// ScoreID   uint
	HoleScores []HoleScore
	Total      uint `gorm:"default:0"`
}

func NewServicesScore(s Score, db *gorm.DB) *ServicesScore {
	return &ServicesScore{
		Score: s,
		DB:    db,
	}
}

type ServicesScore struct {
	Score Score
	DB    *gorm.DB
}

func (ss *ServicesScore) Create(matchID uint) error {
	// get player ids from match id
	var match *Match

	if err := ss.DB.Preload("Players").Where("id = ?", matchID).Find(&match).Error; err != nil {
		return err
	}

	fmt.Println(match)

	var pids []uint

	for _, p := range match.Players {
		pids = append(pids, p.ID)
	}

	var scores []Score

	for _, s := range pids {
		scores = append(scores, Score{
			PlayerID:  s,
			MatchID:   matchID,
			Completed: false,
		})
	}

	//
	// score := Score{
	// 	PlayerID:  playerID,
	// 	MatchID:   matchID,
	// 	Completed: false,
	// }

	if err := ss.DB.Create(&scores).Error; err != nil {
		return err
	}

	for _, hs := range scores {
		if err := ss.GenerateHoleScores(hs.ID); err != nil {
			return err
		}
	}

	return nil
}

func (ss *ServicesScore) Get(scoreID int) (Score, error) {
	var scores []*Score

	if res := ss.DB.Preload("HoleScores").Find(&scores, scoreID); res.Error != nil {
		return Score{}, res.Error
	}

	return *scores[0], nil
}

func (ss *ServicesScore) GetScores(playerID uint, matchID uint) (Score, error) {
	var scores []Score

	if res := ss.DB.Where("player_id = ? AND match_id = ?", playerID, matchID).Preload("HoleScores").Find(&scores); res.Error != nil {
		return Score{}, res.Error
	}

	return scores[0], nil
}

func (ss *ServicesScore) GenerateHoleScores(scoreID uint) error {
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

	if err := ss.DB.Create(&holescores).Error; err != nil {
		return err
	}

	return nil
}

func (s *Score) AfterUpdate(tx *gorm.DB) (err error) {
	var tt uint
	for _, x := range s.HoleScores {
		tt += x.Strokes
	}
	s.Total = tt
	return nil
}
