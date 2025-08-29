package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type VoteRequest struct {
	ParticipantID string `json:"participant_id" binding:"required"`
}

type Vote struct {
	ParticipantID string `json:"participant_id" binding:"required"`
	Timestamp     string `json:"timestamp" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/vote", func(c *gin.Context) {
		var voteRequest VoteRequest
		if err := c.ShouldBindJSON(&voteRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		vote := Vote{
			ParticipantID: voteRequest.ParticipantID,
			Timestamp:     time.Now().UTC().Format(time.RFC3339),
		}

		c.JSON(http.StatusAccepted, vote)
	})

	r.Run(":8080")
}
