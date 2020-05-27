package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type todo struct {
	ID        int    `json:"id"`
	Schedule  string `json:"schedule"`
	Priority  string `json:"priority"`
	TimeLimit string `json:"timeLimit"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:tomoaki7@/todo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	db.SingularTable(true)

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
	/*
		if _, err := db.Exec("INSERT INTO trn_todo (schedule, priority, time_limit) values(?, ?, ?)", body.Schedule, body.Priority, body.TimeLimit); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	*/
}
func display(w http.ResponseWriter, r *http.Request) {
	var todoList []todo
	db.Find(&todoList)

	if err := json.NewEncoder(w).Encode(&todoList); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
func remove(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")
	/*
		if _, err := db.Exec("DELETE FROM trn_todo WHERE id = ?", id); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	*/
}
