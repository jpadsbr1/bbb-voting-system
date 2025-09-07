package repository

import (
	"bbb-voting-system/internal/infrastructure/cache"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

type VoteRedisRepository struct {
	redis *cache.RedisClient
}

func NewVoteRedisRepository(redis *cache.RedisClient) *VoteRedisRepository {
	return &VoteRedisRepository{redis: redis}
}

func (r *VoteRedisRepository) IncrementVote(BigWallID string, ParticipantID string) error {
	ctx := context.Background()

	keyTotal := fmt.Sprintf("votes:%s:total", BigWallID)
	if err := r.redis.GetRedisClient().HIncrBy(ctx, keyTotal, ParticipantID, 1).Err(); err != nil {
		return err
	}

	hour := time.Now().Truncate(time.Hour).Format("2006-01-02-15")
	keyHourly := fmt.Sprintf("votes:%s:%s", BigWallID, hour)
	if err := r.redis.GetRedisClient().HIncrBy(ctx, keyHourly, ParticipantID, 1).Err(); err != nil {
		return err
	}

	log.Printf("Incremented vote for participant %s in Big Wall %s", ParticipantID, BigWallID)

	return nil
}

func (r *VoteRedisRepository) GetTotalVotes(BigWallID string) (map[string]int, error) {
	ctx := context.Background()

	key := fmt.Sprintf("votes:%s:total", BigWallID)
	result, err := r.redis.GetRedisClient().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	votes := make(map[string]int, len(result))
	for ParticipantID, countStr := range result {
		count, _ := strconv.Atoi(countStr)
		votes[ParticipantID] = count
	}

	return votes, nil
}

func (r *VoteRedisRepository) GetHourlyVotes(BigWallID string, hour time.Time) (map[string]int, error) {
	ctx := context.Background()

	key := fmt.Sprintf("votes:%s:%s", BigWallID, hour.Format("2006-01-02-15"))
	result, err := r.redis.GetRedisClient().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	votes := make(map[string]int, len(result))
	for participantID, countStr := range result {
		count, _ := strconv.Atoi(countStr)
		votes[participantID] = count
	}

	return votes, nil
}

func (r *VoteRedisRepository) ResetVotes(BigWallID string) error {
	ctx := context.Background()

	key := fmt.Sprintf("votes:%s:total", BigWallID)

	return r.redis.GetRedisClient().Del(ctx, key).Err()
}

func (r *VoteRedisRepository) ResetHourlyVotes(BigWallID string, hour time.Time) error {
	ctx := context.Background()

	key := fmt.Sprintf("votes:%s:%s", BigWallID, hour.Format("2006-01-02-15"))

	return r.redis.GetRedisClient().Del(ctx, key).Err()
}
