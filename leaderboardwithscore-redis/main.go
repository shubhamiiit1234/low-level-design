package main

import "fmt"

func main() {
	lb := NewLeaderboard()
	lb.client.Del(ctx, lb.key) // reset redis scores
	_ = lb.AddUserScore("user1", 100)
	_ = lb.AddUserScore("user2", 120)
	_ = lb.AddUserScore("user3", 80)
	_ = lb.AddUserScore("user2", 90)

	top, _ := lb.GetTopUsers(3)
	for _, user := range top {
		fmt.Printf("User: %v, Score: %.0f\n", user.Member, user.Score)
	}

	rank, _ := lb.GetRank("user1")
	fmt.Println("Rank of user1:", rank+1)

}
