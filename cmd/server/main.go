package main

import (
	"log"
	"net/http"

	"github.com/Taukom/Book-App/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Library!"))
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
