package usecases

import (
	"bbb-voting-system/internal/domain"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockParticipantRepository struct {
	mock.Mock
}

func (m *MockParticipantRepository) AddParticipant(id string, name string) (*domain.Participant, error) {
	args := m.Called(id, name)
	if p, ok := args.Get(0).(*domain.Participant); ok {
		return p, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockParticipantRepository) GetAllParticipants() ([]*domain.Participant, error) {
	args := m.Called()
	if p, ok := args.Get(0).([]*domain.Participant); ok {
		return p, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockParticipantRepository) EliminateParticipant(id string) (*domain.Participant, error) {
	args := m.Called(id)
	if p, ok := args.Get(0).(*domain.Participant); ok {
		return p, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestAddParticipant(t *testing.T) {
	t.Run("Success Add Participant", func(t *testing.T) {
		mockRepo := new(MockParticipantRepository)
		expected := &domain.Participant{ParticipantID: "123", Name: "João"}

		mockRepo.On("AddParticipant", mock.Anything, "João").Return(expected, nil)

		service := NewParticipantService(mockRepo)
		p, err := service.AddParticipant("João")

		assert.NoError(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, "João", p.Name)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Error adding participant", func(t *testing.T) {
		mockRepo := new(MockParticipantRepository)

		mockRepo.On("AddParticipant", mock.Anything, "Maria").Return(nil, errors.New("Database error"))

		service := NewParticipantService(mockRepo)
		p, err := service.AddParticipant("Maria")

		assert.Error(t, err)
		assert.Nil(t, p)

		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllParticipants(t *testing.T) {
	t.Run("Success Return All Participants", func(t *testing.T) {
		mockRepo := new(MockParticipantRepository)
		expected := []*domain.Participant{
			{ParticipantID: "1", Name: "João"},
			{ParticipantID: "2", Name: "Maria"},
		}

		mockRepo.On("GetAllParticipants").Return(expected, nil)

		service := NewParticipantService(mockRepo)
		result, err := service.GetAllParticipants()

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Maria", result[1].Name)

		mockRepo.AssertExpectations(t)
	})
}

func TestEliminateParticipant(t *testing.T) {
	t.Run("Success Eliminate Participant", func(t *testing.T) {
		mockRepo := new(MockParticipantRepository)
		expected := &domain.Participant{ParticipantID: "1", Name: "João"}

		mockRepo.On("EliminateParticipant", "1").Return(expected, nil)

		service := NewParticipantService(mockRepo)
		p, err := service.EliminateParticipant("1")

		assert.NoError(t, err)
		assert.Equal(t, "João", p.Name)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Eliminating Participant", func(t *testing.T) {
		mockRepo := new(MockParticipantRepository)

		mockRepo.On("EliminateParticipant", "999").Return(nil, errors.New("Not found"))

		service := NewParticipantService(mockRepo)
		p, err := service.EliminateParticipant("999")

		assert.Error(t, err)
		assert.Nil(t, p)

		mockRepo.AssertExpectations(t)
	})
}
