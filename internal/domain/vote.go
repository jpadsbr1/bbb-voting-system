package domain

type VoteRequest struct {
	ParticipantID string `json:"participant_id" binding:"required"`
}

type VoteRepository interface {
	SaveVote(ParticipantID string) error
	GetTotalVotes() (int, error)
}
