package main

import (
	"fmt"
	"sync"

	pq "github.com/jupp0r/go-priority-queue"
)

type Leaderboard struct {
	Users    map[string]*User // userID -> User
	TopUsers *pq.PriorityQueue
	mu       sync.RWMutex
}

func NewLeaderBoard() *Leaderboard {
	priQueue := pq.New()
	return &Leaderboard{
		Users:    map[string]*User{},
		TopUsers: &priQueue,
	}
}

func (l *Leaderboard) AddUser(user *User) {
	l.mu.Lock()
	l.Users[user.ID] = user
	fmt.Println("added")
	l.TopUsers.Insert(user.ID, -float64(user.Score))
	l.mu.Unlock()
}

func (l *Leaderboard) AddScore(userID string, score int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	user, exist := l.Users[userID]
	if !exist {
		user = NewUser(userID)
		l.Users[user.ID] = user
		fmt.Println("added")
		l.TopUsers.Insert(user.ID, -float64(user.Score))
	}

	user.AddScore(score)
	l.TopUsers.UpdatePriority(user.ID, -float64(user.Score))
	user.ViewScore()
}

func (l *Leaderboard) ResetScore(userID string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	user, exist := l.Users[userID]
	if !exist {
		fmt.Println("User doesn't exist!!")
		return
	}

	user.ResetScore()
	l.TopUsers.UpdatePriority(user.ID, -float64(user.Score))
	user.ViewScore()

}

func (l *Leaderboard) ViewAllScores() {
	l.mu.RLock()
	defer l.mu.RUnlock()
	for userID, user := range l.Users {
		fmt.Println(userID, "->", user.Score)
	}
}

func (l *Leaderboard) GetTopKUsers(k int) []User {
	l.mu.Lock()
	defer l.mu.Unlock()
	topKUsers := []User{}
	count := 0
	for l.TopUsers.Len() != 0 && count < k {
		user, _ := l.TopUsers.Pop()
		userID := user.(string)
		topKUsers = append(topKUsers, *l.Users[userID])
		count++
	}

	for _, user := range topKUsers {
		l.TopUsers.Insert(user.ID, -float64(user.Score))
	}

	return topKUsers
}

func (l *Leaderboard) GetRank(userID string) int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	rank := 1
	score := l.Users[userID].Score
	for _, user := range l.Users {
		if user.Score > score {
			rank++
		}
	}
	return rank
}
