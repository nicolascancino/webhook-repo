package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = godotenv.Load()
	mux := http.NewServeMux()

	errs := make(chan error)
	go func() {
		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-stopChan)
	}()

	go func() {
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), mux); err != nil {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
	<-errs
}
