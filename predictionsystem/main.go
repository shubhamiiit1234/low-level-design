package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

/*
	- Users predict outcomes (e.g., Team A wins)
	- Store predictions before match starts
	- After result, system updates user status (win/loss)
	- Maintain leaderboards based on correct predictions
*/

func main() {
	fmt.Println("Prediction System LLD!!!")

	users := map[string]*User{
		"U1": {ID: "U1", Name: "Tyler"},
		"U2": {ID: "U2", Name: "Gamma"},
		"U3": {ID: "U3", Name: "Bablu"},
	}

	// Initialize services
	ps := &PredictionService{}
	ls := &LeaderboardService{}

	// Create sample matches
	match1 := &Match{
		ID:      uuid.NewString(),
		TeamA:   "India",
		TeamB:   "Australia",
		StartAt: time.Now().Add(5 * time.Second), // Starts in future
	}

	match2 := &Match{
		ID:      uuid.NewString(),
		TeamA:   "England",
		TeamB:   "South Africa",
		StartAt: time.Now().Add(5 * time.Second),
	}

	// Users make predictions
	ps.MakePrediction(users["U1"].ID, match1.ID, "India", match1)
	ps.MakePrediction(users["U2"].ID, match1.ID, "Australia", match1)
	ps.MakePrediction(users["U3"].ID, match1.ID, "India", match1)

	ps.MakePrediction(users["U1"].ID, match2.ID, "South Africa", match2)
	ps.MakePrediction(users["U2"].ID, match2.ID, "England", match2)
	ps.MakePrediction(users["U3"].ID, match2.ID, "England", match2)

	// Simulate matches ending
	match1.Result = "India"
	match1.IsClosed = true

	match2.Result = "England"
	match2.IsClosed = true

	// Evaluate all predictions
	ps.EvaluatePredictions(match1, users)
	ps.EvaluatePredictions(match2, users)

	// Show leaderboard
	ls.ShowLeaderboard(users)
}
