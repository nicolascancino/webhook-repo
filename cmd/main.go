package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	_ = godotenv.Load()
	mux := http.NewServeMux()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), mux); err != nil {
		log.Print(err)
	}
}
