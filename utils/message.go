package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type Message struct {
	Id         int
	title      string
	body       string
	created_at time.Time
	updated_at time.Time
}

func (message Message) createMessage(db *sql.DB) string {
	query := "INSERT INTO messages (title, body, created_at, updated_at) VALUES (?, ? , ? , ?)"
	message.created_at = time.Now()
	message.updated_at = time.Now()
	stmt, err := db.Exec(query, message.title, message.body, message.created_at, message.updated_at)
	fmt.Println(stmt)
	if err != nil {
		panic(err.Error())
	}
	return "Message created successfully"
}

func createMessageTable(db *sql.DB) {
	query := "CREATE TABLE IF NOT EXISTS messages (id INT NOT NULL AUTO_INCREMENT, title VARCHAR(255), body VARCHAR(255), created_at DATETIME, updated_at DATETIME, PRIMARY KEY (id))"
	stmt, err := db.Exec(query)

	fmt.Println(stmt, "success")
	if err != nil {
		panic(err.Error())
	}
}

func Message_routes(r *mux.Router, connection *sql.DB) {

	message_route := r.PathPrefix("/messages").Subrouter()

	message_route.HandleFunc("/create_table", Logging(func(w http.ResponseWriter, r *http.Request) {
		createMessageTable(connection)
		fmt.Fprintf(w, "Table created successfully")
	})).Methods("GET")

	message_route.HandleFunc("/create", Logging(func(w http.ResponseWriter, r *http.Request) {
		message := Message{
			title: r.FormValue("title"),
			body:  r.FormValue("body"),
		}
		fmt.Println(message.createMessage(connection))
		fmt.Fprintf(w, "added a new message")
	})).Methods("POST")

	message_route.HandleFunc("/form", Logging(func(w http.ResponseWriter, r *http.Request) {
		formTemplate(w, r, connection)
	}))
}

func formTemplate(w http.ResponseWriter, r *http.Request, connection *sql.DB) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	message := Message{
		title: r.FormValue("title"),
		body:  r.FormValue("body"),
	}
	_ = message

	tmpl.Execute(w, struct{ Success bool }{true})
}
