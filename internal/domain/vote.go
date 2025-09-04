package domain

import "time"

type VoteRequest struct {
	BigWallID     string `json:"bigwall_id" binding:"required"`
	ParticipantID string `json:"participant_id" binding:"required"`
}

type Vote struct {
	VoteID        int       `json:"vote_id" binding:"required"`
	BigWallID     string    `json:"bigwall_id" binding:"required"`
	ParticipantID string    `json:"participant_id" binding:"required"`
	CreatedAt     time.Time `json:"created_at" binding:"required"`
}

type VoteRepository interface {
	Vote(BigWallID string, ParticipantID string) (*Vote, error)
	GetTotalVoteCountByBigWallID(BigWallID string) (int, error)
	GetVoteCountByParticipantID(ParticipantID string, BigWallID string) (int, error)
}
