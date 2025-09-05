package usecases

import (
	"bbb-voting-system/internal/domain"
	"fmt"
)

type VoteService struct {
	voteRepository domain.VoteRepository
}

func NewVoteService(voteRepository domain.VoteRepository) *VoteService {
	return &VoteService{voteRepository: voteRepository}
}

func (s *VoteService) Vote(BigWallID string, ParticipantID string, b *BigWallService) (*domain.Vote, error) {
	currentBigWall, err := b.GetBigWallInfo()
	if err != nil {
		return nil, err
	}

	if currentBigWall.BigWallID != BigWallID {
		return nil, fmt.Errorf("error: This Big Wall is not active")
	}

	return s.voteRepository.Vote(BigWallID, ParticipantID)
}

func (s *VoteService) GetTotalVoteCountByBigWallID(BigWallID string) (int, error) {
	return s.voteRepository.GetTotalVoteCountByBigWallID(BigWallID)
}

func (s *VoteService) GetVoteCountByParticipantID(ParticipantID string, BigWallID string) (int, error) {
	return s.voteRepository.GetVoteCountByParticipantID(ParticipantID, BigWallID)
}
