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

func (r *ParticipantPostgresRepository) AddParticipant(ParticipantID string, name string) (*domain.Participant, error) {
	participant := &domain.Participant{}

	addParticipantQuery := `INSERT INTO participants(participant_id, name) VALUES ($1, $2)
		RETURNING participant_id, name, is_eliminated`

	if err := r.postgres.GetPool().QueryRow(context.Background(),
		addParticipantQuery, ParticipantID, name).Scan(
		&participant.ParticipantID,
		&participant.Name,
		&participant.IsEliminated,
	); err != nil {
		return nil, err
	}

	return participant, nil
}

func (r *ParticipantPostgresRepository) GetAllParticipants() ([]*domain.Participant, error) {
	rows, err := r.postgres.GetPool().Query(context.Background(),
		"SELECT participant_id, name, is_eliminated FROM participants")

	participants := make([]*domain.Participant, 0, 24) // Aloca previamente um array com a capacidade total de participantes do programa para economizar memória
	for rows.Next() {
		participant := &domain.Participant{}
		if err = rows.Scan(&participant.ParticipantID, &participant.Name, &participant.IsEliminated); err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}

	return participants, err
}
