package main

import (
	"log"
	"net/http"

	"github.com/Taukom/Book-App/pkg/db"
	"github.com/Taukom/Book-App/pkg/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()

	mux := http.NewServeMux()

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetBooks(w, r)
		case http.MethodPost:
			handlers.CreateBook(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetBookByID(w, r)
		case http.MethodPut:
			handlers.UpdateBook(w, r)
		case http.MethodDelete:
			handlers.DeleteBook(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
