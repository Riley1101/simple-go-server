package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

type Message struct {
	Id    int
	title string
	body  string
}

func (message Message) createMessage(db *sql.DB) string {
	query := "INSERT INTO messages (title, body) VALUES (?, ?)"
	stmt, err := db.Exec(query, message.title, message.body)
	fmt.Println(stmt)
	if err != nil {
		panic(err.Error())
	}
	return "Message created successfully"
}

func createMessageTable(db *sql.DB) {
	query := "CREATE TABLE IF NOT EXISTS messages (id INT NOT NULL AUTO_INCREMENT, title VARCHAR(255), body VARCHAR(255), created_at DATETIME, updated_at DATETIME, PRIMARY KEY (id))"
	stmt, err := db.Exec(query)
	fmt.Println(stmt)
	if err != nil {
		panic(err.Error())
	}
}

func Message_routes(r *mux.Router, connection *sql.DB) {

	message_route := r.PathPrefix("/messages").Subrouter()

	message_route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := Message{
			title: r.FormValue("title"),
			body:  r.FormValue("body"),
		}
		fmt.Println(message.createMessage(connection))
	}).Methods("POST")

}
