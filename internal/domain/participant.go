package domain

type ParticipantRequest struct {
	Name string `json:"name" binding:"required"`
}

type Participant struct {
	ParticipantID string `json:"participant_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
}

type ParticipantRepository interface {
	AddParticipant(name string) error
	GetAllParticipants() ([]Participant, error)
}
