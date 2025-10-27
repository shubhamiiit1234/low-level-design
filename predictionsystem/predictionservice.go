package main

import (
	"fmt"
	"time"
)

type PredictionService struct {
	Predictions []Prediction
}

func (ps *PredictionService) MakePrediction(userID, matchID, winner string, match *Match) error {
	if time.Now().After(match.StartAt) {
		return fmt.Errorf("cannot predict after match start")
	}
	ps.Predictions = append(ps.Predictions, Prediction{UserID: userID, MatchID: matchID, PredictedWinner: winner})
	fmt.Printf("[PREDICTION] User %s predicted %s to win %s vs %s\n",
		userID, winner, match.TeamA, match.TeamB)
	return nil
}

func (ps *PredictionService) EvaluatePredictions(match *Match, users map[string]*User) {
	fmt.Printf("\n[EVALUATION] Evaluating predictions for match %s (%s vs %s)...\n",
		match.ID, match.TeamA, match.TeamB)

	for i := range ps.Predictions {
		p := &ps.Predictions[i]
		if p.MatchID == match.ID {
			p.IsCorrect = (p.PredictedWinner == match.Result)
			if p.IsCorrect {
				users[p.UserID].Score += 10 // reward points
			}
		}
	}
}
