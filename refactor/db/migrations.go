package db

import (
	"github.com/donaldgifford/tbcctest/services"
)

func AutoMigrate() {
	database.AutoMigrate(&services.Player{})
	database.AutoMigrate(&services.Match{})
	database.AutoMigrate(&services.HoleScore{})
	database.AutoMigrate(&services.Score{})
}