package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()

	log.Println("Server running!")
	http.ListenAndServe(":8080", mux)

}
