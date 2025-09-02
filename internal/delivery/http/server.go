package http

import (
	"bbb-voting-system/internal/infrastructure/storage"
	"bbb-voting-system/internal/repository"
	"bbb-voting-system/internal/usecases"

	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r        *gin.Engine
	http     *http.Server
	postgres *storage.Postgres
}

func NewServer(postgres *storage.Postgres) *Server {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	s := &Server{r: r, postgres: postgres}

	voteRepository := repository.NewVoteLocalRepository()
	voteService := usecases.NewVoteService(voteRepository)
	voteHandler := NewVoteHandler(voteService)

	r.POST("/vote", voteHandler.handleVote)
	r.GET("/total_votes", voteHandler.handleTotalVotes)

	return s
}

func (s *Server) Run(api_port string) error {
	s.http = &http.Server{Addr: api_port, Handler: s.r}
	log.Printf("Listening on %s", api_port)
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if s.http != nil {
		log.Println("Shutting down...")
		_ = s.http.Shutdown(ctx)
	}
}
