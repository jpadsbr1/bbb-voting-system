package repository

import "bbb-voting-system/internal/infrastructure/storage"

type VotePostgresRepository struct {
	postgres *storage.Postgres
}

func NewVotePostgresRepository(postgres *storage.Postgres) *VotePostgresRepository {
	return &VotePostgresRepository{postgres: postgres}
}

func (v *VotePostgresRepository) Vote(BigWallID string, ParticipantID string) error {
	return nil
}

func (v *VotePostgresRepository) GetTotalVoteCountByBigWallID(BigWallID string) (int, error) {
	return 0, nil
}

func (v *VotePostgresRepository) GetVoteCountByParticipantID(ParticipantID string, BigWallID string) (int, error) {
	return 0, nil
}
