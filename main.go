package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"godev/utils"
	"net/http"
)


func main() {
	const PORT = ":5173"
	r := mux.NewRouter()
	connection := db()
	serve_static()
	author_routes(r)
	utils.Message_routes(r, connection)
	utils.Book_routes(r, connection)
	er := http.ListenAndServe(PORT, r)
	if er != nil {
		panic(er)
	} else {
		fmt.Printf("Serving static files\n")
	}
}

func db() *sql.DB {
	db, _ := sql.Open("mysql", "root:admin@(127.0.0.1:3306)/godev?parseTime=true")
	db_err := db.Ping()
	if db_err != nil {
		panic(db_err.Error())
	}
	fmt.Printf("Connected to database\n")
	return db
}

func serve_static() {
	fmt.Printf("Serving static files\n")
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func book_routes(r *mux.Router) {
	r.HandleFunc("/books/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the book: %s on page \n", title)
	}).Methods("GET")

	r.HandleFunc("/books/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the book: %s on page \n", title)
	}).Methods("POST")

	r.HandleFunc("/books/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the book: %s on page \n", title)
	}).Methods("PUT")

	r.HandleFunc("/books/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the book: %s on page \n", title)
	}).Methods("DELETE")
}

func author_routes(r *mux.Router) {
	author_router := r.PathPrefix("/authors").Subrouter()
	author_router.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "You've requested the author: %s\n", name)
	}).Methods("GET")

	author_router.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "You've requested the author: %s\n", name)
	}).Methods("POST")
}
