package main

import (
	"bbb-voting-system/internal/delivery/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var total_votes int = 0

func main() {

	server := http.NewServer()
	go func() {
		if err := server.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	log.Println("Shutting down...")
	server.Shutdown()
}
