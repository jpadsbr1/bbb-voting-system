package repository

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/infrastructure/storage"
	"context"
)

type VotePostgresRepository struct {
	postgres *storage.Postgres
}

func NewVotePostgresRepository(postgres *storage.Postgres) *VotePostgresRepository {
	return &VotePostgresRepository{postgres: postgres}
}

func (v *VotePostgresRepository) Vote(BigWallID string, ParticipantID string) (*domain.Vote, error) {
	vote := &domain.Vote{}

	addVoteQuery := `INSERT INTO votes(bigwall_id, participant_id) VALUES ($1, $2)
		RETURNING vote_id, bigwall_id, participant_id, created_at`

	if err := v.postgres.GetPool().QueryRow(context.Background(),
		addVoteQuery, BigWallID, ParticipantID).Scan(
		&vote.VoteID,
		&vote.BigWallID,
		&vote.ParticipantID,
		&vote.CreatedAt,
	); err != nil {
		return nil, err
	}

	return vote, nil
}

func (v *VotePostgresRepository) GetTotalVoteCountByBigWallID(BigWallID string) (int, error) {
	var vote_count int

	getVoteCountQuery := `SELECT COUNT(*) FROM votes WHERE bigwall_id = $1`

	if err := v.postgres.GetPool().QueryRow(context.Background(),
		getVoteCountQuery, BigWallID).Scan(
		&vote_count,
	); err != nil {
		return 0, err
	}

	return vote_count, nil
}

func (v *VotePostgresRepository) GetVoteCountByParticipantID(ParticipantID string, BigWallID string) (int, error) {
	var vote_count int

	getVoteCountByParticipantQuery := `SELECT COUNT(*) FROM votes WHERE participant_id = $1 AND bigwall_id = $2`

	if err := v.postgres.GetPool().QueryRow(context.Background(),
		getVoteCountByParticipantQuery, ParticipantID, BigWallID).Scan(
		&vote_count,
	); err != nil {
		return 0, err
	}

	return vote_count, nil
}
