package http

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/usecases"

	"net/http"

	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	voteService    *usecases.VoteService
	bigWallService *usecases.BigWallService
}

func NewVoteHandler(voteService *usecases.VoteService, bigWallService *usecases.BigWallService) *VoteHandler {
	return &VoteHandler{voteService: voteService, bigWallService: bigWallService}
}

func (h *VoteHandler) handleVote(c *gin.Context) {
	var voteRequest *domain.VoteRequest
	if err := c.ShouldBindJSON(&voteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.voteService.Vote(voteRequest.BigWallID, voteRequest.ParticipantID, h.bigWallService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vote": "success",
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

func (h *VoteHandler) handleGetVoteCountByParticipantID(c *gin.Context) {
	voteCountRequest := domain.VoteRequest{}

	if err := c.ShouldBindJSON(&voteCountRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	voteCountByParticipantID, err := h.voteService.GetVoteCountByParticipantID(voteCountRequest.ParticipantID, voteCountRequest.BigWallID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ParticipantID": voteCountRequest.ParticipantID,
		"BigWallID":     voteCountRequest.BigWallID,
		"voteCount":     voteCountByParticipantID,
	})
}
