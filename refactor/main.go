package main

import (
	"fmt"

	"github.com/donaldgifford/tbcctest/db"
	"github.com/donaldgifford/tbcctest/services"
)

func main() {
	db.Init()
	gorm := db.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate()
	dbGorm.Ping()

	// ps := services.NewServicesPlayer(services.Player{}, gorm)
	//
	// ms := services.NewServicesMatch(services.Match{}, gorm)

	ss := services.NewServicesScore(services.Score{}, gorm)
	// hs := services.NewServicesHoleScore(services.HoleScore{}, gorm)

	// err = ps.CreatePlayer(services.Player{
	// 	Name:     "Donald",
	// 	Email:    "donald@poop.com",
	// 	Handicap: 9,
	// })
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	//
	// err = ps.CreatePlayer(services.Player{
	// 	Name:     "Bob",
	// 	Email:    "bob@poop.com",
	// 	Handicap: 13,
	// })
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	//
	// var pp []services.Player
	//
	// p, err := ps.GetPlayerById(1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// pp = append(pp, p)
	//
	// p2, err := ps.GetPlayerById(2)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	//
	// pp = append(pp, p2)
	//
	// err = ms.CreateMatch(services.Match{
	// 	Players:  pp,
	// 	Title:    "new match",
	// 	GameType: "Stroke",
	// 	Holes:    "9",
	// 	Scoring:  "Gross",
	// })
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// initialize scores for matches for each player based on Holes

	// err = ss.Create(1, 1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	//
	// err = ss.Create(2, 1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// gm, err := ms.GetMatches()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	//
	// for _, m := range gm {
	// 	fmt.Println(m.ID)
	// 	fmt.Println(m.Title)
	// 	fmt.Println(m.GameType)
	// 	fmt.Println(m.Scoring)
	// 	for _, pl := range m.Players {
	// 		fmt.Println(pl.Name)
	// 	}
	// 	for _, sc := range m.Scores {
	// 		fmt.Println(sc.ID)
	//
	// 		err := hs.GenerateHoleScores(sc.ID)
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}
	// 	}
	// }

	// get scores by score id
	score, err := ss.Get(1)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(score.PlayerID)

	p1, err := ss.GetScoresByPlayerID(1)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(p1)

	for _, pp := range p1 {
		fmt.Println(pp.MatchID)
	}
}
