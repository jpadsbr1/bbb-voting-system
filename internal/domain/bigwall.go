package domain

import "time"

type BigWallRequest struct {
	ParticipantIDs []string `json:"participant_ids" binding:"required"`
}

type BigWallParticipant struct {
	ParticipantID string `json:"participant_id" binding:"required"`
	Votes         int    `json:"votes" binding:"required"`
}

type BigWall struct {
	BigWallID  string     `json:"bigwall_id" binding:"required"`
	StartTime  time.Time  `json:"start_time" binding:"required"`
	EndTime    *time.Time `json:"end_time"`
	IsActive   bool       `json:"is_active" binding:"required"`
	TotalVotes int        `json:"total_votes" binding:"required"`
}

type BigWallRepository interface {
	CreateBigWallUnit(BigWallID string, ParticipantIDs []string) (*BigWall, error)
	InsertCrossParticipantBigWall(BigWallID string, ParticipantIDs []string) error
	GetBigWallInfo() (*BigWall, error)
	EndBigWall(BigWallID string) (*BigWall, error)
	GetBigWallParticipants(BigWallID string) ([]*BigWallParticipant, error)
	GetMostVotedParticipants(BigWallID string) (string, int, error)
}
