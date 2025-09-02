package usecases

import "bbb-voting-system/internal/domain"

type ParticipantService struct {
	participantRepository domain.ParticipantRepository
}

func NewParticipantService(participantRepository domain.ParticipantRepository) *ParticipantService {
	return &ParticipantService{participantRepository: participantRepository}
}

func (p *ParticipantService) AddParticipant(name string) error {
	return p.participantRepository.AddParticipant(name)
}

func (p *ParticipantService) GetAllParticipants() ([]domain.Participant, error) {
	return p.participantRepository.GetAllParticipants()
}
