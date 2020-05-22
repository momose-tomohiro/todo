package main

import (
	"database/sql"
	"encoding/json"
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

var dbConn = connect()

func connect() *sql.DB {
	dbConn, err := sql.Open("mysql", "root:0111@/todo")
	if err != nil {
		log.Println(err)
	}
	return dbConn
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("views")))
	http.HandleFunc("/register", register)
	http.HandleFunc("/display", display)
	http.HandleFunc("/remove", remove)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func register(w http.ResponseWriter, r *http.Request) {
	form := todo{}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	schedule := form.Schedule
	priority := form.Priority
	timeLimit := form.TimeLimit

	stmt, err := dbConn.Prepare("INSERT INTO trn_todo(schedule, priority, time_limit) VALUES(?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(schedule, priority, timeLimit)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
func display(w http.ResponseWriter, r *http.Request) {
	todoList := todoList{}
	rows, err := dbConn.Query("SELECT id, schedule, priority, time_limit FROM trn_todo")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for rows.Next() {
		todo := todo{}
		err := rows.Scan(&todo.ID, &todo.Schedule, &todo.Priority, &todo.TimeLimit)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		todoList = append(todoList, todo)
	}
	defer rows.Close()

	JSONTodoList, err := json.Marshal(todoList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(JSONTodoList)
}
func remove(w http.ResponseWriter, r *http.Request) {
	form := todo{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.Unmarshal(b, &form)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	targetID := form.ID

	stmt, err := dbConn.Prepare("DELETE FROM trn_todo WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(targetID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
