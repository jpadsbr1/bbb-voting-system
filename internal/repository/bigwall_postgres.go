package repository

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/infrastructure/storage"

	"context"
)

type BigWallPostgresRepository struct {
	postgres *storage.Postgres
}

func NewBigWallPostgresRepository(postgres *storage.Postgres) *BigWallPostgresRepository {
	return &BigWallPostgresRepository{postgres: postgres}
}

func (r *BigWallPostgresRepository) CreateBigWallUnit(BigWallID string, ParticipantIDs []string) (*domain.BigWall, error) {
	bigWall := &domain.BigWall{}

	createBigWallQuery := `INSERT INTO bigwall(bigwall_id)
		VALUES ($1) RETURNING bigwall_id, start_time, is_active, total_votes`

	if err := r.postgres.GetPool().QueryRow(context.Background(),
		createBigWallQuery, BigWallID).Scan(
		&bigWall.BigWallID,
		&bigWall.StartTime,
		&bigWall.IsActive,
		&bigWall.TotalVotes,
	); err != nil {
		return nil, err
	}

	return bigWall, nil
}

func (r *BigWallPostgresRepository) InsertCrossParticipantBigWall(BigWallID string, ParticipantIDs []string) error {
	crossParticipantBigWallQuery := `INSERT INTO participants_bigwall(bigwall_id, participant_id)
		VALUES ($1, $2)`

	for _, ParticipantID := range ParticipantIDs {
		_, err := r.postgres.GetPool().Exec(context.Background(),
			crossParticipantBigWallQuery, BigWallID, ParticipantID)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *BigWallPostgresRepository) GetBigWallInfo() (*domain.BigWall, error) {
	bigWall := &domain.BigWall{}

	getBigWallInfoQuery := `SELECT bigwall_id, start_time, end_time, is_active, total_votes
		FROM bigwall WHERE is_active = true`

	if err := r.postgres.GetPool().QueryRow(context.Background(),
		getBigWallInfoQuery).Scan(
		&bigWall.BigWallID,
		&bigWall.StartTime,
		&bigWall.EndTime,
		&bigWall.IsActive,
		&bigWall.TotalVotes,
	); err != nil {
		return nil, err
	}

	return bigWall, nil
}

func (r *BigWallPostgresRepository) EndBigWall(bigWallID string) (*domain.BigWall, error) {
	bigWall := &domain.BigWall{}

	endBigWallQuery := `UPDATE bigwall
		SET is_active = false,
		end_time = NOW()
		WHERE bigwall_id = $1
		AND is_active = true
		RETURNING bigwall_id, start_time, end_time, is_active, total_votes`

	if err := r.postgres.GetPool().QueryRow(context.Background(),
		endBigWallQuery, bigWallID).Scan(
		&bigWall.BigWallID,
		&bigWall.StartTime,
		&bigWall.EndTime,
		&bigWall.IsActive,
		&bigWall.TotalVotes,
	); err != nil {
		return nil, err
	}

	return bigWall, nil
}

func (r *BigWallPostgresRepository) GetBigWallParticipants(bigWallID string) ([]*domain.BigWallParticipant, error) {
	participants := []*domain.BigWallParticipant{}

	getBigWallParticipantsQuery := `SELECT participant_id, votes
		FROM participants_bigwall
		WHERE bigwall_id = $1`

	rows, err := r.postgres.GetPool().Query(context.Background(),
		getBigWallParticipantsQuery, bigWallID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var participantID string
		var votes int

		if err = rows.Scan(&participantID, &votes); err != nil {
			return nil, err
		}

		participant := &domain.BigWallParticipant{
			ParticipantID: participantID,
			Votes:         votes,
		}

		participants = append(participants, participant)
	}

	return participants, nil
}

func (r *BigWallPostgresRepository) GetMostVotedParticipants(bigWallID string) (string, int, error) {
	var participant_id string
	var votes int

	getMostVotedParticipantsQuery := `SELECT participant_id, votes
		FROM participants_bigwall
		WHERE bigwall_id = $1
		ORDER BY votes DESC
		LIMIT 1`

	if err := r.postgres.GetPool().QueryRow(context.Background(),
		getMostVotedParticipantsQuery, bigWallID).Scan(
		&participant_id,
		&votes,
	); err != nil {
		return "", 0, err
	}

	return participant_id, votes, nil
}
