package repository

import (
	"bbb-voting-system/internal/infrastructure/storage"
	"context"
	"time"
)

type VotePostgresRepository struct {
	postgres *storage.Postgres
}

func NewVotePostgresRepository(postgres *storage.Postgres) *VotePostgresRepository {
	return &VotePostgresRepository{postgres: postgres}
}

func (v *VotePostgresRepository) IncrementVotes(BigWallID string, count int) error {
	incrementVotesQuery := `
        INSERT INTO bigwall (bigwall_id, total_votes)
        VALUES ($1, $2)
        ON CONFLICT (bigwall_id)
        DO UPDATE SET total_votes = bigwall.total_votes + EXCLUDED.total_votes;
    `
	_, err := v.postgres.GetPool().Exec(context.Background(),
		incrementVotesQuery, BigWallID, count)

	if err != nil {
		return err
	}

	return nil
}

func (v *VotePostgresRepository) IncrementVotesPerParticipant(BigWallID string, ParticipantID string, count int) error {
	incrementVotesPerParticipantQuery := `INSERT INTO participants_bigwall (participant_id, bigwall_id, votes)
		VALUES ($1, $2, $3)
		ON CONFLICT (participant_id, bigwall_id)
		DO UPDATE SET votes = participants_bigwall.votes + EXCLUDED.votes;
	`
	_, err := v.postgres.GetPool().Exec(context.Background(),
		incrementVotesPerParticipantQuery, ParticipantID, BigWallID, count)

	if err != nil {
		return err
	}

	return nil
}

func (v *VotePostgresRepository) IncrementHourlyVotes(BigWallID string, ParticipantID string, hour time.Time, count int) error {
	incrementHourlyVotesQuery := `INSERT INTO votes_hourly (bigwall_id, participant_id, hour, total_votes)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (bigwall_id, participant_id, hour)
		DO UPDATE SET total_votes = votes_hourly.total_votes + EXCLUDED.total_votes;
	`

	_, err := v.postgres.GetPool().Exec(context.Background(),
		incrementHourlyVotesQuery, BigWallID, ParticipantID, hour, count)
	if err != nil {
		return err
	}

	return nil
}

func (v *VotePostgresRepository) GetTotalVoteCountByBigWallID(BigWallID string) (int, error) {
	var vote_count int

	getVoteCountQuery := `SELECT total_votes FROM bigwall WHERE bigwall_id = $1`

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

	getVoteCountByParticipantQuery := `SELECT votes FROM participants_bigwall WHERE participant_id = $1 AND bigwall_id = $2`

	if err := v.postgres.GetPool().QueryRow(context.Background(),
		getVoteCountByParticipantQuery, ParticipantID, BigWallID).Scan(
		&vote_count,
	); err != nil {
		return 0, err
	}

	return vote_count, nil
}
