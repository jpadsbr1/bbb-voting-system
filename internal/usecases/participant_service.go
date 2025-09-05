package usecases

import (
	"bbb-voting-system/internal/domain"

	uuid "github.com/nu7hatch/gouuid"
)

type ParticipantService struct {
	participantRepository domain.ParticipantRepository
}

func NewParticipantService(participantRepository domain.ParticipantRepository) *ParticipantService {
	return &ParticipantService{participantRepository: participantRepository}
}

func (p *ParticipantService) AddParticipant(name string) (*domain.Participant, error) {
	ParticipantID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	Participant, err := p.participantRepository.AddParticipant(ParticipantID.String(), name)
	if err != nil {
		return nil, err
	}

	return Participant, nil
}

func (p *ParticipantService) GetAllParticipants() ([]*domain.Participant, error) {
	return p.participantRepository.GetAllParticipants()
}

func (p *ParticipantService) EliminateParticipant(ParticipantID string) (*domain.Participant, error) {
	return p.participantRepository.EliminateParticipant(ParticipantID)
}
