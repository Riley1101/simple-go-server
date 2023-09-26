package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

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

func main() {
	const PORT = ":5173"
	r := mux.NewRouter()
	serve_static()
	book_routes(r)
	author_routes(r)

	err := http.ListenAndServe(PORT, r)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Serving static files\n")
	}
}
