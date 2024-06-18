package services

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Username string   `json:"username"`
	Matches  []*Match `gorm:"many2many:player_matches"`
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

func (sp *ServicesPlayer) GetAllPlayers() ([]*Player, error) {
	var players []*Player

	if res := sp.DB.Find(&players); res.Error != nil {
		return nil, res.Error
	}

	return players, nil
}

func ConverDateTime(tz string, dt time.Time) string {
	loc, _ := time.LoadLocation(tz)

	return dt.In(loc).Format(time.RFC822Z)
}

func (sp *ServicesPlayer) GetPlayerById(id int) (Player, error) {
	var players []*Player

	if res := sp.DB.Find(&players, id); res.Error != nil {
		return Player{}, res.Error
	}

	return *players[0], nil
}

func (sp *ServicesPlayer) CreatePlayer(p Player) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), 8)
	if err != nil {
		return err
	}

	newPlayer := Player{
		Email:    p.Email,
		Username: p.Username,
		Name:     p.Name,
		// Profile:  Profile{Link: p.Profile.Link},
		Password: string(hashedPassword),
	}

	if err := sp.DB.Create(&newPlayer).Error; err != nil {
		return err
	}

	return err
}

func (sp *ServicesPlayer) CheckEmail(email string) (Player, error) {
	var p Player

	if err := sp.DB.Where("email = ?", email).First(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func ConvertDateTime(tz string, dt time.Time) string {
	loc, _ := time.LoadLocation(tz)

	return dt.In(loc).Format(time.RFC822Z)
}

func (sp *ServicesPlayer) UpdatePlayer(player Player) error {
	existingPlayer := new(Player)

	if res := sp.DB.Find(&existingPlayer, "id = ?", player.ID); res.Error != nil {
		return res.Error
	}

	if player.Name != "" {
		existingPlayer.Name = player.Name
	}

	if player.Username != "" {
		existingPlayer.Username = player.Username
	}
	if player.Email != "" {
		existingPlayer.Email = player.Email
	}

	if err := sp.DB.Save(&existingPlayer).Error; err != nil {
		return err
	}

	return nil
}
