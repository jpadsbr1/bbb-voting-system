package repository

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/infrastructure/storage"

	"context"
)

type ParticipantPostgresRepository struct {
	postgres *storage.Postgres
}

func NewParticipantPostgresRepository(postgres *storage.Postgres) *ParticipantPostgresRepository {
	return &ParticipantPostgresRepository{postgres: postgres}
}

func (r *ParticipantPostgresRepository) AddParticipant(name string) error {
	_, err := r.postgres.GetPool().Exec(context.Background(),
		"INSERT INTO participants(name) VALUES ($1)", name)

	return err
}

func (r *ParticipantPostgresRepository) GetAllParticipants() ([]domain.Participant, error) {
	return nil, nil
}
