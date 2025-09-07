package usecases

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/infrastructure/worker"
	"context"
	"fmt"
	"log"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type BigWallService struct {
	bigWallRepository   domain.BigWallRepository
	voteCacheRepository domain.VoteCacheRepository
	voteRepository      domain.VoteRepository
	voteFlusher         *worker.VoteFlusher
	flusherContext      context.Context
	flusherCancel       context.CancelFunc
}

func NewBigWallService(bigWallRepository domain.BigWallRepository, voteCacheRepository domain.VoteCacheRepository, voteRepository domain.VoteRepository) *BigWallService {
	return &BigWallService{bigWallRepository: bigWallRepository, voteCacheRepository: voteCacheRepository, voteRepository: voteRepository}
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

	b.voteFlusher = worker.NewVoteFlusher(b.voteCacheRepository, b.voteRepository, 5*time.Second)
	b.flusherContext, b.flusherCancel = context.WithCancel(context.Background())

	go b.voteFlusher.Start(b.flusherContext, bigWall.BigWallID)

	return bigWall, nil

}

func (b *BigWallService) GetBigWallInfo() (*domain.BigWall, error) {
	return b.bigWallRepository.GetBigWallInfo()
}

func (b *BigWallService) EndBigWall(BigWallID string, p *ParticipantService) (*domain.BigWall, error) {
	if b.voteFlusher != nil {
		b.voteFlusher.FlushOnce(BigWallID)
	}

	if b.flusherCancel != nil {
		b.flusherCancel()
		log.Printf("Flusher stopped for Big Wall %s", BigWallID)
	}

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

func (b *BigWallService) GetBigWallParticipants(BigWallID string) ([]*domain.BigWallParticipant, error) {
	return b.bigWallRepository.GetBigWallParticipants(BigWallID)
}
