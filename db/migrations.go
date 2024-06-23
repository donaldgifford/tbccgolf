package db

import (
	"github.com/donaldgifford/tbccgolf/services"
)

func AutoMigrate() {
	// database.AutoMigrate(&models.Profile{})
	database.AutoMigrate(&services.Player{})
	database.AutoMigrate(&services.Match{})
	database.AutoMigrate(&services.Score{})
	database.AutoMigrate(&services.Stroke{})
	database.AutoMigrate(&services.HoleScore{})
	// database.AutoMigrate(&services.Course{})
	// database.AutoMigrate(&services.ScoreCard{})
	// database.AutoMigrate(&services.Hole{})
	// database.AutoMigrate(&models.Hole{})
	// database.AutoMigrate(&models.Round{})
	// database.AutoMigrate(&models.ScoreCard{})
	// database.AutoMigrate(&models.HoleScore{})
	// database.AutoMigrate(&models.Course{})
}
