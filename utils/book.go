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
	Title      string
	Author     string
	Created_At time.Time
	Updated_At time.Time
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
	book.Created_At = time.Now()
	book.Updated_At = time.Now()
	stmt, err := db.Exec(query, book.Title, book.Author, book.Created_At, book.Updated_At)
	fmt.Println(stmt)
	if err != nil {
		panic(err.Error())
	}
	return "Book created successfully"
}

func getBooks(db *sql.DB) []Book {

	query := "SELECT * from books;"
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var books []Book
	for result.Next() {
		var book Book
		err := result.Scan(&book.Id, &book.Title, &book.Author, &book.Created_At, &book.Updated_At)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}

	return books
}

func Book_routes(r *mux.Router, connection *sql.DB) {
	book_route := r.PathPrefix("/books").Subrouter()

	book_route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		books := getBooks(connection)
		fmt.Fprintf(w, "Books: %v", books)
	}).Methods("GET")

	book_route.HandleFunc("/create_table", func(w http.ResponseWriter, r *http.Request) {
		createBookTable(connection)
		fmt.Fprintf(w, "Table created successfully")
	}).Methods("GET")

	book_route.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		author := r.FormValue("author")
		book := Book{Title: title, Author: author}
		fmt.Fprintf(w, book.createBook(connection))
	}).Methods("POST")
}
