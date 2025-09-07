package domain

import "time"

type VoteRequest struct {
	BigWallID     string `json:"bigwall_id" binding:"required"`
	ParticipantID string `json:"participant_id" binding:"required"`
}

type VoteRepository interface {
	IncrementVotes(BigWallID string, count int) error
	IncrementVotesPerParticipant(BigWallID string, ParticipantID string, count int) error
	IncrementHourlyVotes(BigWallID string, ParticipantID string, hour time.Time, count int) error
	GetTotalVoteCountByBigWallID(BigWallID string) (int, error)
	GetVoteCountByParticipantID(ParticipantID string, BigWallID string) (int, error)
}
