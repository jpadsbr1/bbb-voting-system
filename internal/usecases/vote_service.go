package usecases

import "bbb-voting-system/internal/domain"

type VoteService struct {
	voteRepository domain.VoteRepository
}

func NewVoteService(voteRepository domain.VoteRepository) *VoteService {
	return &VoteService{voteRepository: voteRepository}
}

func (s *VoteService) Vote(BigWallID string, ParticipantID string) error {
	return s.voteRepository.Vote(BigWallID, ParticipantID)
}

func (s *VoteService) GetTotalVoteCountByBigWallID(BigWallID string) (int, error) {
	return s.voteRepository.GetTotalVoteCountByBigWallID(BigWallID)
}

func (s *VoteService) GetVoteCountByParticipantID(ParticipantID string, BigWallID string) (int, error) {
	return s.voteRepository.GetVoteCountByParticipantID(ParticipantID, BigWallID)
}
