package http

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/usecases"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ParticipantHandler struct {
	participantService *usecases.ParticipantService
}

func NewParticipantHandler(participantService *usecases.ParticipantService) *ParticipantHandler {
	return &ParticipantHandler{participantService: participantService}
}

func (h *ParticipantHandler) handleAddParticipant(c *gin.Context) {
	var participantRequest *domain.ParticipantRequest
	if err := c.ShouldBindJSON(&participantRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Participant, err := h.participantService.AddParticipant(participantRequest.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Participant": Participant,
	})
}

func (h *ParticipantHandler) handleGetAllParticipants(c *gin.Context) {
	participants, err := h.participantService.GetAllParticipants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"participants": participants,
	})
}
