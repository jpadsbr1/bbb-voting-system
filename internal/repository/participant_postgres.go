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

func (r *ParticipantPostgresRepository) GetAllParticipants() ([]*domain.Participant, error) {
	rows, err := r.postgres.GetPool().Query(context.Background(),
		"SELECT participant_id, name FROM participants")

	participants := make([]*domain.Participant, 0, 24) // Aloca previamente um array com a capacidade total de participantes do programa para economizar memória
	for rows.Next() {
		participant := &domain.Participant{}
		if err = rows.Scan(&participant.ParticipantID, &participant.Name); err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}

	return participants, err
}
