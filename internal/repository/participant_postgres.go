package repository

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/infrastructure/storage"
	"log"

	"context"

	uuid "github.com/nu7hatch/gouuid"
)

type ParticipantPostgresRepository struct {
	postgres *storage.Postgres
}

func NewParticipantPostgresRepository(postgres *storage.Postgres) *ParticipantPostgresRepository {
	return &ParticipantPostgresRepository{postgres: postgres}
}

func (r *ParticipantPostgresRepository) AddParticipant(name string) error {
	participant_id, err := uuid.NewV4()
	if err != nil {
		log.Fatal("error:", err)
	}

	if _, err := r.postgres.GetPool().Exec(context.Background(),
		"INSERT INTO participants(participant_id, name) VALUES ($1, $2)", participant_id, name); err != nil {
		return err
	}

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
