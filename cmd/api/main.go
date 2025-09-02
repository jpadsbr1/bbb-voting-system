package main

import (
	"bbb-voting-system/internal/config"
	"bbb-voting-system/internal/delivery/http"
	"bbb-voting-system/internal/infrastructure/storage"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.LoadEnvironmentVariables()

	postgres_url := config.GetPostgresURL()
	postgres := storage.NewPostgres(postgres_url)
	defer postgres.Close()

	server := http.NewServer(postgres)
	go func() {
		if err := server.Run(":" + os.Getenv("API_PORT")); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	server.Shutdown()
}
