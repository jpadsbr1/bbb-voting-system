package domain

import "time"

type VoteCacheRepository interface {
	IncrementVote(BigWallID string, ParticipantID string) error
	GetTotalVotes(BigWallID string) (map[string]int, error)
	GetHourlyVotes(BigWallID string, hour time.Time) (map[string]int, error)
	ResetVotes(BigWallID string) error
	ResetHourlyVotes(BigWallID string, hour time.Time) error
}
