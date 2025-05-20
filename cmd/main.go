package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ihgazi/shiritori/internal/matchmaker"
)

func main() {
	serveMux := http.NewServeMux()

	gameMatcher := new(matchmaker.MatchMaker)

	serveMux.HandleFunc("/queue", gameMatcher.ServeHTTP)

	server := http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)

	exitSignal := <-sigChan
	log.Println("Received terminate, graceful shutdown", exitSignal)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
