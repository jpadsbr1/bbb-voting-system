package usecases

import (
	"bbb-voting-system/internal/domain"
	"fmt"
	"log"
)

type VoteService struct {
	voteRepository      domain.VoteRepository
	voteCacheRepository domain.VoteCacheRepository
}

func NewVoteService(voteRepository domain.VoteRepository, voteCacheRepository domain.VoteCacheRepository) *VoteService {
	return &VoteService{voteRepository: voteRepository, voteCacheRepository: voteCacheRepository}
}

func (s *VoteService) Vote(BigWallID string, ParticipantID string, b *BigWallService) error {
	currentBigWall, err := b.GetBigWallInfo()
	if err != nil {
		return err
	}

	if currentBigWall.BigWallID != BigWallID {
		return fmt.Errorf("error: This Big Wall is not active")
	}

	bigWallParticipants, err := b.GetBigWallParticipants(BigWallID)
	if err != nil {
		return err
	}

	for _, participant := range bigWallParticipants {
		if participant.ParticipantID == ParticipantID {
			break
		} else {
			return fmt.Errorf("error: Participant %s is not in this Big Wall", ParticipantID)
		}
	}

	if err := s.voteCacheRepository.IncrementVote(BigWallID, ParticipantID); err != nil {
		log.Printf("Error incrementing vote: %v", err)
	}

	return nil
}

func (s *VoteService) GetTotalVoteCountByBigWallID(BigWallID string) (int, error) {
	return s.voteRepository.GetTotalVoteCountByBigWallID(BigWallID)
}

func (s *VoteService) GetVoteCountByParticipantID(ParticipantID string, BigWallID string) (int, error) {
	return s.voteRepository.GetVoteCountByParticipantID(ParticipantID, BigWallID)
}

func (s *VoteService) GetVoteHourlyCountByBigWallID(BigWallID string) ([]*domain.VoteHourlyCount, error) {
	return s.voteRepository.GetVoteHourlyCountByBigWallID(BigWallID)
}
