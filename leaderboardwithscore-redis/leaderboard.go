package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Leaderboard struct {
	client *redis.Client
	key    string
}

func NewLeaderboard() *Leaderboard {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &Leaderboard{
		client: rdb,
		key:    "leaderboard",
	}
}

func (lb *Leaderboard) AddUserScore(userID string, score float64) error {
	return lb.client.ZIncrBy(ctx, lb.key, score, userID).Err()
}

func (lb *Leaderboard) GetTopUsers(k int64) ([]redis.Z, error) {
	return lb.client.ZRevRangeWithScores(ctx, lb.key, 0, k-1).Result()
}

func (lb *Leaderboard) GetRank(userID string) (int64, error) {
	return lb.client.ZRevRank(ctx, lb.key, userID).Result()
}

func (lb *Leaderboard) ResetUser(userID string) error {
	return lb.client.ZAdd(ctx, lb.key, redis.Z{Score: 0, Member: userID}).Err()
}
