package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Requirements:
		1. Add to user's existing score
		2. Should return the top K users with the highest scores
		3. Reset the score of user to 0
		4. Should return the rank of a user (TBD)

	Core Entities:
		1. User
			- id, score
			- AddScore(score int), ResetScore()
		2. Leaderboard
			- map[string]*User (id -> User)
			- GetTopKUsers(k int) []User, ResetScore(userID string), GetRank(userID string)


*/

// func main() {
// 	fmt.Println("leaderboard with score LLD!!!")
// 	leaderBoard := NewLeaderBoard()

// 	user1 := &User{
// 		ID:    "u1",
// 		Score: 0,
// 	}
// 	user2 := &User{
// 		ID:    "u2",
// 		Score: 0,
// 	}
// 	user3 := &User{
// 		ID:    "u3",
// 		Score: 0,
// 	}
// 	leaderBoard.AddUser(user1)
// 	leaderBoard.AddUser(user2)
// 	leaderBoard.AddUser(user3)

// 	leaderBoard.ViewAllScores()

// 	leaderBoard.AddScore("u1", 20)
// 	leaderBoard.AddScore("u2", 10)
// 	leaderBoard.AddScore("u3", 30)

// 	leaderBoard.ViewAllScores()

// 	topKUsers := leaderBoard.GetTopKUsers(3)
// 	for _, user := range topKUsers {
// 		fmt.Println(user.ID, " -> ", user.Score)
// 	}

// 	leaderBoard.AddScore("u2", 30)

// 	topKUsers = leaderBoard.GetTopKUsers(3)
// 	for _, user := range topKUsers {
// 		fmt.Println(user.ID, " -> ", user.Score)
// 	}

// 	// leaderBoard.ResetScore(user3)

// 	// topKUsers = leaderBoard.GetTopKUsers()
// 	// for _, user := range topKUsers {
// 	// 	fmt.Println(user.ID, " -> ", user.Score)
// 	// }

// 	fmt.Println(leaderBoard.GetRank("u3"))
// }

func main() {
	fmt.Println("Concurrent Leaderboard Test!")
	leaderBoard := NewLeaderBoard()

	users := []*User{
		{ID: "u1", Score: 0},
		{ID: "u2", Score: 0},
		{ID: "u3", Score: 0},
		{ID: "u4", Score: 0},
		{ID: "u5", Score: 0},
	}

	// Add users to leaderboard
	for _, user := range users {
		leaderBoard.AddUser(user)
	}

	var wg sync.WaitGroup

	// Simulate concurrent AddScore
	for _, user := range users {
		wg.Add(1)
		go func(uid string) {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				leaderBoard.AddScore(uid, 10)
				time.Sleep(10 * time.Millisecond)
			}
		}(user.ID)
	}

	// Simulate concurrent ResetScore in the middle
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(25 * time.Millisecond) // Let some AddScores happen
		leaderBoard.ResetScore("u3")
	}()

	wg.Wait()

	fmt.Println("\nFinal Scores:")
	leaderBoard.ViewAllScores()

	fmt.Println("\nTop 3 Users:")
	topK := leaderBoard.GetTopKUsers(3)
	for _, user := range topK {
		fmt.Println(user.ID, "->", user.Score)
	}

	fmt.Println("\nRanks:")
	for _, user := range users {
		fmt.Println("Rank of", user.ID, ":", leaderBoard.GetRank(user.ID))
	}
}
