package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type Message struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func (message Message) createMessage(db *sql.DB) string {
	query := "INSERT INTO messages (title, body, created_at, updated_at) VALUES (?, ? , ? , ?)"
	message.Created_At = time.Now()
	message.Updated_At = time.Now()
	stmt, err := db.Exec(query, message.Title, message.Body, message.Created_At, message.Updated_At)
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
			Title: r.FormValue("title"),
			Body:  r.FormValue("body"),
		}
		fmt.Println(message.createMessage(connection))
		fmt.Fprintf(w, "added a new message")
	})).Methods("POST")

	message_route.HandleFunc("/form", Logging(func(w http.ResponseWriter, r *http.Request) {
		formTemplate(w, r, connection)
	}))

	message_route.HandleFunc("/json", Logging(func(w http.ResponseWriter, r *http.Request) {
		messages := []Message{}
		query := "SELECT * FROM messages"
		rows, err := connection.Query(query)
		if err != nil {
			panic(err.Error())
		}
		for rows.Next() {
			var message Message
			rows.Scan(&message.Id, &message.Title, &message.Body, &message.Created_At, &message.Updated_At)
			messages = append(messages, message)
		}
		defer rows.Close()
        for _, message := range messages {
            fmt.Println(message.Title,message.Body)
        }
		// send messages as json
        json.NewEncoder(w).Encode(messages)
        
	}))
}

func formTemplate(w http.ResponseWriter, r *http.Request, connection *sql.DB) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	message := Message{
		Title: r.FormValue("title"),
		Body:  r.FormValue("body"),
	}
	fmt.Println(message)
	message.createMessage(connection)

	tmpl.Execute(w, struct{ Success bool }{true})
}
