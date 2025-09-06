package http

import (
	"bbb-voting-system/internal/infrastructure/cache"
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
	redis    *cache.RedisClient
}

func NewServer(postgres *storage.Postgres, redis *cache.RedisClient) *Server {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	s := &Server{r: r, postgres: postgres, redis: redis}

	participantRepository := repository.NewParticipantPostgresRepository(s.postgres)
	participantService := usecases.NewParticipantService(participantRepository)
	participantHandler := NewParticipantHandler(participantService)

	bigWallRepository := repository.NewBigWallPostgresRepository(s.postgres)
	bigWallService := usecases.NewBigWallService(bigWallRepository)
	bigWallHandler := NewBigWallHandler(bigWallService, participantService)

	voteRepository := repository.NewVotePostgresRepository(s.postgres)
	voteService := usecases.NewVoteService(voteRepository)
	voteHandler := NewVoteHandler(voteService, bigWallService)

	r.POST("/participant", participantHandler.handleAddParticipant)
	r.GET("/participants", participantHandler.handleGetAllParticipants)

	r.POST("/bigwall/create", bigWallHandler.handleCreateBigWall)
	r.GET("/bigwall", bigWallHandler.handleGetBigWallInfo)
	r.PATCH("/bigwall/end/:bigWallID", bigWallHandler.handleEndBigWall)

	r.POST("/vote", voteHandler.handleVote)
	r.GET("/votes/total/:bigWallID", voteHandler.handleTotalVoteCountByBigWallID)
	r.GET("/votes/participant", voteHandler.handleGetVoteCountByParticipantID)

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
