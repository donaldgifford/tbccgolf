package services

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	PlayerID  uint
	MatchID   uint
	Completed bool
	// ScoreID   uint
	HoleScores []HoleScore
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

func (ss *ServicesScore) Create(playerID uint, matchID uint) error {
	score := Score{
		PlayerID:  playerID,
		MatchID:   matchID,
		Completed: false,
	}

	if err := ss.DB.Create(&score).Error; err != nil {
		return err
	}

	return nil
}

func (ss *ServicesScore) Get(scoreID int) (Score, error) {
	var scores []*Score

	if res := ss.DB.Find(&scores, scoreID); res.Error != nil {
		// if res := ss.DB.Model(Score{}).Preload("Players").Preload("Match").Find(&scores, scoreID); res.Error != nil {
		return Score{}, res.Error
	}

	return *scores[0], nil
}

func (ss *ServicesScore) GetScoresByPlayerID(playerID uint) ([]Score, error) {
	var scores []Score

	if res := ss.DB.Where("player_id = ?", playerID).Find(&scores); res.Error != nil {
		return nil, res.Error
	}

	return scores, nil
}
