package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"go_exam/internal/handlers"
	"go_exam/internal/repository"
)

func main() {
	db, err := sql.Open("sqlite", "./data/exam.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	categoryRepo := repository.NewCategoryRepository(db)
	homeHandler := handlers.NewHomeHandler(categoryRepo)

	http.Handle("/", homeHandler)

	log.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
