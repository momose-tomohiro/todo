package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type todo struct {
	ID        int    `db:"id" json:"id"`
	Schedule  string `db:"schedule" json:"schedule"`
	Priority  string `db:"priority" json:"priority"`
	TimeLimit string `db:"time_limit" json:"timeLimit"`
}

type todoList []todo

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:tomoaki7@/todo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("views")))
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			display(w, r)
		case http.MethodPost:
			register(w, r)
		case http.MethodDelete:
			remove(w, r)
		}
	})
	log.Println("start http server :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func register(w http.ResponseWriter, r *http.Request) {
	var body todo

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if _, err := db.Exec("INSERT INTO trn_todo (schedule, priority, time_limit) values(?, ?, ?)", body.Schedule, body.Priority, body.TimeLimit); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
func display(w http.ResponseWriter, r *http.Request) {
	var todoList todoList
	rows, err := db.Query("SELECT id, schedule, priority, time_limit FROM trn_todo")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for rows.Next() {
		var (
			id        int
			schedule  string
			priority  string
			timeLimit string
		)
		if err := rows.Scan(&id, &schedule, &priority, &timeLimit); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		todoList = append(todoList, todo{id, schedule, priority, timeLimit})
	}
	if err := json.NewEncoder(w).Encode(&todoList); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
func remove(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if _, err := db.Exec("DELETE FROM trn_todo WHERE id = ?", id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
