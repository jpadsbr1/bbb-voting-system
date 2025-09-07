package worker

import (
	"bbb-voting-system/internal/domain"
	"context"
	"log"
	"time"
)

type VoteFlusher struct {
	VoteCacheRepository domain.VoteCacheRepository
	VoteRepository      domain.VoteRepository
	Interval            time.Duration
}

func NewVoteFlusher(voteCacheRepository domain.VoteCacheRepository, voteRepository domain.VoteRepository, interval time.Duration) *VoteFlusher {
	return &VoteFlusher{VoteCacheRepository: voteCacheRepository, VoteRepository: voteRepository, Interval: interval}
}

func (v *VoteFlusher) Start(ctx context.Context, BigWallID string) {
	ticker := time.NewTicker(v.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Vote flusher stopped")
			return
		case <-ticker.C:
			v.flush(BigWallID)
		}
	}
}

func (v *VoteFlusher) flush(BigWallID string) {
	log.Println("Running unified flusher...")

	totals, err := v.VoteCacheRepository.GetTotalVotes(BigWallID)
	if err != nil {
		log.Printf("Failed to get total votes: %v", err)
	} else {
		for ParticipantID, count := range totals {
			if err := v.VoteRepository.IncrementVotes(BigWallID, count); err != nil {
				log.Printf("Failed to flush total votes: %v", err)
			}
			if err := v.VoteRepository.IncrementVotesPerParticipant(BigWallID, ParticipantID, count); err != nil {
				log.Printf("Failed to flush total votes per participant: %v", err)
			}
		}
		if err := v.VoteCacheRepository.ResetVotes(BigWallID); err != nil {
			log.Printf("Failed to reset total votes in redis: %v", err)
		}
	}

	currentHour := time.Now().Truncate(time.Hour)
	hourly, err := v.VoteCacheRepository.GetHourlyVotes(BigWallID, currentHour)
	if err != nil {
		log.Printf("Failed to get hourly votes: %v", err)
	} else {
		for participantID, count := range hourly {
			if err := v.VoteRepository.IncrementHourlyVotes(BigWallID, participantID, currentHour, count); err != nil {
				log.Printf("Failed to flush hourly votes: %v", err)
			}
		}
		if err := v.VoteCacheRepository.ResetHourlyVotes(BigWallID, currentHour); err != nil {
			log.Printf("Failed to reset hourly votes in redis: %v", err)
		}
	}
}
