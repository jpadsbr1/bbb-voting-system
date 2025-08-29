package domain

type Participant struct {
	ParticipantID string `json:"participant_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
}
