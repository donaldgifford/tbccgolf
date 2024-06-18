package services

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// TODO: Add scoring, scores, etc.

type Match struct {
	gorm.Model
	NetScore    bool `form:"NetScore"`
	ScoringType string
	Length      int       `form:"Length"`
	Players     []*Player `gorm:"many2many:player_matches" form:"Players"`
	Completed   bool
	Title       string
}

type PlayerList struct {
	PlayerID int
	Name     string
	Email    string
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

// Return all matches paginated or error
func (ms *ServicesMatch) GetMatches() ([]*Match, error) {
	var matches []*Match

	if res := ms.DB.Model(Match{}).Preload("Players").Find(&matches); res.Error != nil {
		return nil, res.Error
	}

	return matches, nil
}

// Return match by match id or error
func (ms *ServicesMatch) GetMatch(matchID int) (Match, error) {
	var matches []*Match

	if res := ms.DB.Find(&matches, matchID); res.Error != nil {
		return Match{}, res.Error
	}

	return *matches[0], nil
}

// func (ms *ServicesMatch) GetMatchesByPlayer() {}

// Create new match
func (ms *ServicesMatch) CreateMatch(m Match) error {
	newMatch := Match{
		ScoringType: m.ScoringType,
		NetScore:    m.NetScore,
		Players:     m.Players,
		Completed:   false,
		Length:      m.Length,
	}

	if err := ms.DB.Create(&newMatch).Error; err != nil {
		return err
	}

	return nil
}

// Completes match
func (ms *ServicesMatch) CompleteMatch(matchID int) error {
	m, err := ms.GetMatch(matchID)
	if err != nil {
		return err
	}
	ms.DB.Model(&m).Update("Completed", true)
	return nil
}

// List players for matches
// Only returns the username, email, and ID
func (ms *ServicesMatch) ListPlayers() ([]PlayerList, error) {
	var players []*Player
	var playerList []PlayerList

	if res := ms.DB.Find(&players); res.Error != nil {
		return nil, res.Error
	}

	for _, player := range players {
		fmt.Printf("Player Name: %s", player.Username)
		pList := PlayerList{
			PlayerID: int(player.ID),
			Name:     player.Username,
			Email:    player.Email,
		}
		playerList = append(playerList, pList)
	}

	return playerList, nil
	// return players, nil
}

func (ms *ServicesMatch) GetPlayer(name string) (*Player, error) {
	var player *Player

	res := ms.DB.Where("username = ?", name).First(&player)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return player, res.Error
	}
	fmt.Println(player.Username)
	return player, nil
}
