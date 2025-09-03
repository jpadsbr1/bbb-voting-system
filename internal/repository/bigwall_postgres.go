package repository

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/infrastructure/storage"
	"log"

	"context"

	uuid "github.com/nu7hatch/gouuid"
)

type BigWallPostgresRepository struct {
	postgres *storage.Postgres
}

func NewBigWallPostgresRepository(postgres *storage.Postgres) *BigWallPostgresRepository {
	return &BigWallPostgresRepository{postgres: postgres}
}

func (r *BigWallPostgresRepository) CreateBigWall(ParticipantIDs []string) (*domain.BigWall, error) {
	bigWall := &domain.BigWall{}

	bigwall_id, err := uuid.NewV4()
	if err != nil {
		log.Fatal("error:", err)
	}

	createBigWallQuery := `INSERT INTO bigwall(bigwall_id)
		VALUES ($1) RETURNING bigwall_id, start_time, is_active`

	if err := r.postgres.GetPool().QueryRow(context.Background(),
		createBigWallQuery, bigwall_id).Scan(
		&bigWall.BigWallID,
		&bigWall.StartTime,
		&bigWall.IsActive,
	); err != nil {
		return nil, err
	}

	crossParticipantBigWallQuery := `INSERT INTO participants_bigwall(bigwall_id, participant_id)
		VALUES ($1, $2)`

	for _, participant_id := range ParticipantIDs {
		_, err := r.postgres.GetPool().Exec(context.Background(),
			crossParticipantBigWallQuery, bigwall_id, participant_id)

		if err != nil {
			return nil, err
		}
	}

	return bigWall, err
}

func (r *BigWallPostgresRepository) GetBigWallInfo() (*domain.BigWall, error) {
	bigWall := &domain.BigWall{}

	getBigWallInfoQuery := `SELECT bigwall_id, start_time, end_time, is_active
		FROM bigwall WHERE is_active = true`

	if err := r.postgres.GetPool().QueryRow(context.Background(),
		getBigWallInfoQuery).Scan(
		&bigWall.BigWallID,
		&bigWall.StartTime,
		&bigWall.EndTime,
		&bigWall.IsActive,
	); err != nil {
		return nil, err
	}

	return bigWall, nil
}

func (r *BigWallPostgresRepository) EndBigWall(bigWallID string) (*domain.BigWall, error) {
	return nil, nil
}
