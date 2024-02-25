package main

import (
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	_ = godotenv.Load()
	mux := http.NewServeMux()

	_ = http.ListenAndServe(":8080", mux)
}
