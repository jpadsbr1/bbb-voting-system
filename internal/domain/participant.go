package domain

type ParticipantRequest struct {
	Name string `json:"name" binding:"required"`
}

type Participant struct {
	ParticipantID string `json:"participant_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	IsEliminated  bool   `json:"is_eliminated" binding:"required"`
}

type ParticipantRepository interface {
	AddParticipant(ParticipantID string, name string) (*Participant, error)
	GetAllParticipants() ([]*Participant, error)
}
