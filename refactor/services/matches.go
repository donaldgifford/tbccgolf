package services

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	Players  []Player `gorm:"many2many:player_matches"`
	Title    string
	GameType string
	Holes    string
	Scoring  string
	Scores   []Score
}

func NewServicesMatch(m Match, db *gorm.DB) *ServicesMatch {
	return &ServicesMatch{
		Match: m,
		DB:    db,
	}
}

type ServicesMatch struct {
	Match Match
	DB    *gorm.DB
}

func (ms *ServicesMatch) CreateMatch(m Match) error {
	newMatch := Match{
		Players:  m.Players,
		Title:    m.Title,
		GameType: m.GameType,
		Holes:    m.Holes,
		Scoring:  m.Scoring,
	}

	if err := ms.DB.Create(&newMatch).Error; err != nil {
		return err
	}

	return nil
}

func (ms *ServicesMatch) GetMatches() ([]*Match, error) {
	var matches []*Match

	if res := ms.DB.Model(Match{}).Preload("Players").Preload("Scores").Find(&matches); res.Error != nil {
		return nil, res.Error
	}

	return matches, nil
}
