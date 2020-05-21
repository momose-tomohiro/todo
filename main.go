package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type todo struct {
	ID        string `db:"id" json:"id"`
	Schedule  string `db:"schedule" json:"schedule"`
	Priority  string `db:"priority" json:"priority"`
	TimeLimit string `db:"time_limit" json:"timeLimit"`
}

type todoList []todo

func main() {
	db, err := sql.Open("mysql", "root:0111@/todo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("views")))
	http.HandleFunc("/register", register)
	http.HandleFunc("/display", display)
	http.HandleFunc("/remove", remove)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func register(w http.ResponseWriter, r *http.Request) {
	form := todo{}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.Unmarshal(b, &form)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	schedule := form.Schedule
	priority := form.Priority
	timeLimit := form.TimeLimit

	db, err := sql.Open("mysql", "root:0111@/todo")
	if err != nil {
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO trn_todo(schedule, priority, time_limit) VALUES(?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	defer stmt.Close()

	_, err = stmt.Exec(schedule, priority, timeLimit)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}
func display(w http.ResponseWriter, r *http.Request) {
	todoList := todoList{}

	db, err := sql.Open("mysql", "root:0111@/todo")
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, schedule, priority, time_limit FROM trn_todo")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	for rows.Next() {
		todo := todo{}
		err := rows.Scan(&todo.ID, &todo.Schedule, &todo.Priority, &todo.TimeLimit)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		todoList = append(todoList, todo)
	}
	defer rows.Close()

	JSONTodoList, err := json.Marshal(todoList)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	//json.NewEncoder(w).Encode(JSONTodoList)
	w.Header().Set("Content-Type", "application/json")
	w.Write(JSONTodoList)
}
func remove(w http.ResponseWriter, r *http.Request) {
	form := todo{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.Unmarshal(b, &form)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	targetID := form.ID
	log.Println(targetID)
	db, err := sql.Open("mysql", "root:0111@/todo")
	if err != nil {
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM trn_todo WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	defer stmt.Close()
	_, err = stmt.Exec(targetID)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}
