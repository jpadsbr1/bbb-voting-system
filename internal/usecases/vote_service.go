package usecases

import "bbb-voting-system/internal/domain"

type VoteService struct {
	voteRepository domain.VoteRepository
}

func NewVoteService(voteRepository domain.VoteRepository) *VoteService {
	return &VoteService{voteRepository: voteRepository}
}

func (s *VoteService) SaveVote(participantID string) error {
	return s.voteRepository.SaveVote(participantID)
}

func (s *VoteService) GetTotalVotes() (int, error) {
	return s.voteRepository.GetTotalVotes()
}
