/* score.go
*
* When a match is started, scores are created for each player. Based on the match configuration
* is how many holes are added to the score. Also it is worth keeping net and gross since a
* players handicap may change over time if recalculated at viewing.
*
*
*
*
 */
package services

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	PlayerID  uint
	MatchID   uint
	Completed bool
	Holes     []HoleScore
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

//
// func generateStrokes(hs []HoleScore) []Stroke {
// 	s := make([]Stroke, len(hs))
//
// 	for h := range hs {
// 		var st Stroke
// 		st.HoleNumber = uint(hs[h].Number)
// 		st.Hole = uint(h)
// 		st.Strokes = 0
//
// 		s = append(s, st)
// 	}
// 	return s
// }

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

// func (sc *ServicesScore) Create(players []Player, match Match) error {
// holeScores := generateHoles(match.StartingHole, match.Length)
// holeStrokes := generateStrokes(holeScores)
//
// // sh := match.StartingHole + match.Length
//
// var scores []Score
//
// for _, p := range players {
// 	var s Score
// 	s.Player = p
// 	s.PlayerID = p.ID
// 	s.Match = match
// 	s.MatchID = match.ID
// 	s.Holes = holeScores
// 	s.Strokes = holeStrokes
//
// 	scores = append(scores, s)
// }
//
// sc.DB.Create(&scores)
//
// 	return nil
// }

// func (sc *ServicesScore) Get(scoreID int) (Score, error) {
// 	var scores []*Score
//
// 	if res := sc.DB.Model(Score{}).Preload("Players").Preload("Match").Find(&scores, scoreID); res.Error != nil {
// 		return Score{}, res.Error
// 	}
//
// 	return *scores[0], nil
// }

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

// func (sc *ServicesScore) GetScoresByPlayerID() {}
// func (sc *ServicesScore) GetScoresByMatchID()  {}
// func (sc *ServicesScore) GetScoresByStatus()   {}
func (ss *ServicesScore) GetAll() ([]*Score, error) {
	var scores []*Score

	// if res := ss.DB.Model(Score{}).Preload("Player").Preload("Match").Find(&scores); res.Error != nil {
	if res := ss.DB.Find(&scores); res.Error != nil {
		return nil, res.Error
	}

	return scores, nil
}

func (sc *ServicesScore) UpdateHole() {}
