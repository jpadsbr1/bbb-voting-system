package http

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/usecases"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	voteService *usecases.VoteService
}

func NewVoteHandler(voteService *usecases.VoteService) *VoteHandler {
	return &VoteHandler{voteService: voteService}
}

func (h *VoteHandler) handleVote(c *gin.Context) {
	var voteRequest *domain.VoteRequest
	if err := c.ShouldBindJSON(&voteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.voteService.SaveVote(voteRequest.ParticipantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"participant_id": voteRequest.ParticipantID,
		"timestamp":      time.Now().Format(time.RFC3339),
	})
}

func (h *VoteHandler) handleTotalVotes(c *gin.Context) {
	total_votes, err := h.voteService.GetTotalVotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_votes": total_votes,
	})
}
