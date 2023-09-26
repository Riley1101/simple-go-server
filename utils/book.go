package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Book struct {
	Id         int
	title      string
	author     string
	created_at time.Time
	updated_at time.Time
}

func createBookTable(db *sql.DB) {
	query := "CREATE TABLE IF NOT EXISTS books (id INT NOT NULL AUTO_INCREMENT, title VARCHAR(255), author VARCHAR(255), created_at DATETIME, updated_at DATETIME, PRIMARY KEY (id))"
	stmt, err := db.Exec(query)
	fmt.Println(stmt)
	if err != nil {
		panic(err.Error())
	}
}


func (book Book) createBook(db *sql.DB) string {
	query := "INSERT INTO books (title, author,created_at,updated_at) VALUES (?, ?, ?, ?)"
	book.created_at = time.Now()
	book.updated_at = time.Now()
	stmt, err := db.Exec(query, book.title, book.author, book.created_at, book.updated_at)
	fmt.Println(stmt)
	if err != nil {
		panic(err.Error())
	}
	return "Book created successfully"
}

func Book_routes(r *mux.Router, connection *sql.DB) {
	book_route := r.PathPrefix("/books").Subrouter()

	book_route.HandleFunc("/create_table", func(w http.ResponseWriter, r *http.Request) {
		createBookTable(connection)
		fmt.Fprintf(w, "Table created successfully")
	}).Methods("GET")

	book_route.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		author := r.FormValue("author")
		book := Book{title: title, author: author}
		fmt.Fprintf(w, book.createBook(connection))
	}).Methods("POST")
}
