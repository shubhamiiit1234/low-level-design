package main

import (
	"fmt"
	"sort"
)

type LeaderboardService struct{}

// ShowLeaderboard prints users in descending order of their scores.
func (ls *LeaderboardService) ShowLeaderboard(users map[string]*User) {
	fmt.Println("Leaderboard:")

	// Convert map → slice for sorting
	userList := make([]*User, 0, len(users))
	for _, u := range users {
		userList = append(userList, u)
	}

	// Sort by score (descending)
	sort.Slice(userList, func(i, j int) bool {
		return userList[i].Score > userList[j].Score
	})

	// Display leaderboard
	for rank, u := range userList {
		fmt.Printf("%d. %s - %d points\n", rank+1, u.Name, u.Score)
	}
}
