package http

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/usecases"

	"net/http"

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

	vote, err := h.voteService.Vote(voteRequest.BigWallID, voteRequest.ParticipantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vote": vote,
	})
}

func (h *VoteHandler) handleTotalVoteCountByBigWallID(c *gin.Context) {
	BigWallID := c.Param("bigWallID")

	totalVoteCount, err := h.voteService.GetTotalVoteCountByBigWallID(BigWallID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalVoteCount": totalVoteCount,
	})
}
