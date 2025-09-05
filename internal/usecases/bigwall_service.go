package usecases

import (
	"bbb-voting-system/internal/domain"
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

type BigWallService struct {
	bigWallRepository domain.BigWallRepository
}

func NewBigWallService(bigWallRepository domain.BigWallRepository) *BigWallService {
	return &BigWallService{bigWallRepository: bigWallRepository}
}

func (b *BigWallService) CreateBigWall(ParticipantIDs []string) (*domain.BigWall, error) {
	if len(ParticipantIDs) < 2 {
		return nil, fmt.Errorf("error: A Big Wall must contain at least 2 participants")
	}

	if _, err := b.bigWallRepository.GetBigWallInfo(); err == nil {
		return nil, fmt.Errorf("error: A Big Wall is already active")
	}

	BigWallID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	bigWall, err := b.bigWallRepository.CreateBigWallUnit(BigWallID.String(), ParticipantIDs)
	if err != nil {
		return nil, err
	}

	if err := b.bigWallRepository.InsertCrossParticipantBigWall(BigWallID.String(), ParticipantIDs); err != nil {
		return nil, err
	}

	return bigWall, nil

}

func (b *BigWallService) GetBigWallInfo() (*domain.BigWall, error) {
	return b.bigWallRepository.GetBigWallInfo()
}

func (b *BigWallService) EndBigWall(BigWallID string, p *ParticipantService) (*domain.BigWall, error) {
	currentBigWall, err := b.bigWallRepository.GetBigWallInfo()
	if err != nil {
		return nil, err
	}

	if currentBigWall.BigWallID != BigWallID {
		return nil, fmt.Errorf("error: This Big Wall is not active")
	}

	finishedBigWall, err := b.bigWallRepository.EndBigWall(BigWallID)
	if err != nil {
		return nil, err
	}

	participantID, votes, err := b.bigWallRepository.GetMostVotedParticipants(BigWallID)
	if err != nil {
		return nil, err
	}

	if votes > 0 {
		_, err = p.EliminateParticipant(participantID)
		if err != nil {
			return nil, err
		}
	}

	return finishedBigWall, nil
}
