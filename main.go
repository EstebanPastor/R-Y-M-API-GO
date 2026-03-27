package main

import (
	"database/sql"
	"log"
	"net/http"

	"rymapi/internal/service"
	"rymapi/internal/store"
	"rymapi/internal/transport"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./books.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := `
CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    author TEXT NOT NULL
)`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("books/", bookHandler.HandleBookById)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
