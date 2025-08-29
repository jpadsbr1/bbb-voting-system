package domain

type Participant struct {
	ParticipantID string `json:"participant_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
}

type ParticipantRepository interface {
	AddParticipant(name string) (Participant, error)
	GetAllParticipants() ([]Participant, error)
}
