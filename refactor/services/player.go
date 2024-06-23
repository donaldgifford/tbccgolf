package services

import (
	"fmt"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name     string
	Email    string
	Handicap int
	Matches  []*Match `gorm:"many2many:player_matches"`
	Scores   []Score
}

func NewServicesPlayer(p Player, db *gorm.DB) *ServicesPlayer {
	return &ServicesPlayer{
		Player: p,
		DB:     db,
	}
}

type ServicesPlayer struct {
	Player Player
	DB     *gorm.DB
}

func (sp *ServicesPlayer) CreatePlayer(p Player) error {
	newPlayer := Player{
		Email:    p.Email,
		Name:     p.Name,
		Handicap: p.Handicap,
	}

	// if err := sp.DB.Create(&newPlayer).Error; err != nil {
	// 	return err
	// }

	res := sp.DB.FirstOrCreate(&newPlayer, Player{Name: newPlayer.Name})
	if res.RowsAffected > 0 {
		fmt.Println("New player created")
	} else {
		fmt.Println("Player already exists")
	}

	return nil
}

func (sp *ServicesPlayer) GetAllPlayers() ([]*Player, error) {
	var players []*Player

	if res := sp.DB.Find(&players); res.Error != nil {
		return nil, res.Error
	}

	return players, nil
}

func (sp *ServicesPlayer) GetPlayerById(id int) (Player, error) {
	var players []*Player

	if res := sp.DB.Find(&players, id); res.Error != nil {
		return Player{}, res.Error
	}

	return *players[0], nil
}
