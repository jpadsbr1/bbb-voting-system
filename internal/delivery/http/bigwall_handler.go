package http

import (
	"bbb-voting-system/internal/domain"
	"bbb-voting-system/internal/usecases"

	"net/http"

	"github.com/gin-gonic/gin"
)

type BigWallHandler struct {
	bigWallService *usecases.BigWallService
}

func NewBigWallHandler(bigWallService *usecases.BigWallService) *BigWallHandler {
	return &BigWallHandler{bigWallService: bigWallService}
}

func (h *BigWallHandler) handleCreateBigWall(c *gin.Context) {
	var bigWallRequest *domain.BigWallRequest
	if err := c.ShouldBindJSON(&bigWallRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bigWall, err := h.bigWallService.CreateBigWall(bigWallRequest.ParticipantIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bigWall": bigWall,
	})
}
