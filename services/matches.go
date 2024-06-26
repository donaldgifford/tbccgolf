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

func (ms *ServicesMatch) Create(m Match) error {
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

func (ms *ServicesMatch) DbQuery() *gorm.DB {
	return ms.DB
}

// Return match by match id or error
func (ms *ServicesMatch) GetMatch(matchID int) (Match, error) {
	var matches []*Match

	if res := ms.DB.Model(Match{}).Preload("Players").Preload("Scores").Find(&matches, matchID); res.Error != nil {
		return Match{}, res.Error
	}

	return *matches[0], nil
}
