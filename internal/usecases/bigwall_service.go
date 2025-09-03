package usecases

import "bbb-voting-system/internal/domain"

type BigWallService struct {
	bigWallRepository domain.BigWallRepository
}

func NewBigWallService(bigWallRepository domain.BigWallRepository) *BigWallService {
	return &BigWallService{bigWallRepository: bigWallRepository}
}

func (b *BigWallService) CreateBigWall(ParticipantIDs []string) (*domain.BigWall, error) {
	return b.bigWallRepository.CreateBigWall(ParticipantIDs)
}
